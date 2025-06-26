package arm

import (
	"fmt"
	parser "compiler/parser"
	"github.com/antlr4-go/antlr/v4"
    
)

// ArmVisitor recorre el árbol de análisis y genera código ARM64.
type ArmVisitor struct {
	parser.BaseVlangVisitor
	Generator   *ArmGenerator
	LastResult  string // ← agrega esto
}

// NewArmVisitor construye un nuevo visitor para generación ARM.
func NewArmVisitor() *ArmVisitor {
	return &ArmVisitor{
        Generator: &ArmGenerator{
            Instructions: []string{},
            Variables:    make(map[string]string),
            Output:       "",
            StdLib:       NewStandardLibrary(),
            VarOffset:    0,
            tempRegs:     []string{X1, X2, X3, X4}, // evita X0 (lo usa syscall)
            tempIndex:    0,
        },

	}
}

func (g *ArmGenerator) NextTempReg() string {
	reg := g.tempRegs[g.tempIndex]
	g.tempIndex = (g.tempIndex + 1) % len(g.tempRegs)
	return reg
}


// Visit permite que ArmVisitor recorra cualquier árbol ANTLR.
func (v *ArmVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

// VisitBloque ejecuta las sentencias dentro de un bloque { ... }
func (v *ArmVisitor) VisitBloque(ctx *parser.BloqueContext) interface{} {
    fmt.Println("📦 Visitando arm bloque")
    for _, decl := range ctx.AllDeclaraciones() {
        v.Visit(decl)
    }
    return nil
}


func (v *ArmVisitor) VisitDeclaraciones(ctx *parser.DeclaracionesContext) interface{} {
	fmt.Println("🧩 Revisando tipo de declaración")

	if ctx.Funcion() != nil {
		fmt.Println("✅ ctx.Funcion() no es nil")
		fn := ctx.Funcion()
		fmt.Println("➡️ Nombre función:", fn.ID().GetText())
		return v.Visit(fn)
	}

	if ctx.FuncionStruct() != nil {
		fmt.Println("🧱 ctx.FuncionStruct() detectado")
		return v.Visit(ctx.FuncionStruct())
	}

	if ctx.VarDcl() != nil {
		fmt.Println("📦 ctx.VarDcl detectado")
		return v.Visit(ctx.VarDcl())
	}

	if ctx.StructDecl() != nil {
		fmt.Println("🏗 ctx.StructDecl detectado")
		return v.Visit(ctx.StructDecl())
	}

	if ctx.Stmt() != nil {
		fmt.Println("📜 ctx.Stmt detectado")
		return v.Visit(ctx.Stmt())
	}

	fmt.Println("⚠️ No se detectó ningún tipo válido en Declaraciones")
	return nil
}








// VisitPrograma recorre todas las declaraciones del programa.
func (v *ArmVisitor) VisitPrograma(ctx *parser.ProgramaContext) interface{} {
    fmt.Println("🌲 Visitando Programa (ARM)", ctx.GetText())
    for _, decl := range ctx.AllDeclaraciones() {
        fmt.Println("Declaración:", decl.GetText())
        v.Visit(decl)
    }
    return nil
}







// VisitValorEntero genera código para una literal entera.
func (v *ArmVisitor) VisitValorEntero(ctx *parser.ValorEnteroContext) interface{} {
    fmt.Println("🔢 Visitando valor arm entero:", ctx.GetText())

    value := ctx.GetText()
    reg := v.Generator.NextTempReg()
    v.Generator.Comment("Literal entero: " + value)
    v.Generator.Add(fmt.Sprintf("MOV %s, #%s", reg, value))
    v.LastResult = reg
    return reg
}


// VisitPrintStatement traduce una instrucción println(...).
func (v *ArmVisitor) VisitPrintStatement(ctx *parser.PrintStatementContext) interface{} {
    fmt.Println("🖨️ VisitPrintStatement ARM")

    v.Generator.StdLib.Use("print_integer")

    if ctx.Parametros() == nil || len(ctx.Parametros().AllExpresion()) == 0 {
        return nil
    }

    expr := ctx.Parametros().AllExpresion()[0]
    fmt.Println("Nodo expresión:", expr.GetText(), "Tipo:", fmt.Sprintf("%T", expr))
    val := v.Visit(expr)
    fmt.Println("Valor devuelto por expresión:", val, "Tipo:", fmt.Sprintf("%T", val))

    reg, ok := val.(string)
    if !ok {
        fmt.Println("⚠️ No se pudo obtener el registro temporal para impresión. Tipo:", fmt.Sprintf("%T", val))
        return nil
    }

    v.Generator.Comment("Print statement")
	v.Generator.Add(fmt.Sprintf("MOV X0, %s", reg))
    return nil
}



// VisitStmt permite visitar una sentencia (stmt).
func (v *ArmVisitor) VisitStmt(ctx *parser.StmtContext) interface{} {
	child := ctx.GetChild(0)
	if parseTree, ok := child.(antlr.ParseTree); ok {
		return v.Visit(parseTree)
	}

	fmt.Println("⚠️ No se pudo castear el hijo de StmtContext")
	return nil
}



// VisitExpresionStatement permite ejecutar expresiones dentro de sentencias.
func (v *ArmVisitor) VisitExpresionStatement(ctx *parser.ExpresionStatementContext) interface{} {
	fmt.Println("📏 Visitando arm ExpresionStatement")

	// Simplemente visita la expresión contenida
	return v.Visit(ctx.Expresion())
}



func (v *ArmVisitor) VisitFuncion(ctx *parser.FuncionContext) interface{} {
    funcName := ctx.ID().GetText()
    fmt.Println("🔧 Generando función:", funcName)

    if funcName == "main" {
        v.Generator.Add("main:")
    }

    bloque := ctx.Bloque()
    if bloque != nil {
        fmt.Println("Procesando bloque de", funcName, "con texto:", bloque.GetText())
        v.Visit(bloque)
    } else {
        fmt.Println("⚠️ Bloque no encontrado para", funcName)
    }
    return nil
}




func (v *ArmVisitor) VisitReturnStatement(ctx *parser.ReturnStatementContext) interface{} {
	v.Generator.Comment("Return statement")

	// Si retorna una expresión, podrías evaluarla
	if ctx.Expresion() != nil {
		v.Visit(ctx.Expresion())
	}

	// Terminar ejecución (equivalente a exit syscall)
	v.Generator.Add("mov x8, #93") // syscall: exit
	v.Generator.Add("svc #0")
	return nil
}


func (v *ArmVisitor) VisitVariableDeclaration(ctx *parser.VariableDeclarationContext) interface{} {
	varName := ctx.ID().GetText()
	fmt.Println("🔧 Declarando variable:", varName)

	// Reservamos espacio en la pila
	offset := v.Generator.VarOffset
	memLocation := fmt.Sprintf("[SP, #%d]", offset)

	if ctx.Expresion() != nil {
		val := v.Visit(ctx.Expresion())
		reg, ok := val.(string)
		if !ok {
			fmt.Printf("⚠️ Error: expresión no retornó un registro válido para variable '%s'\n", varName)
			return nil
		}

		v.Generator.Add(fmt.Sprintf("STR %s, %s", reg, memLocation))
		v.Generator.Comment(fmt.Sprintf("Variable %s inicializada con valor en %s", varName, reg))
	} else {
		v.Generator.Add(fmt.Sprintf("MOV %s, #0", X0))
		v.Generator.Add(fmt.Sprintf("STR %s, %s", X0, memLocation))
		v.Generator.Comment(fmt.Sprintf("Variable %s declarada sin valor (inicializada en 0)", varName))
	}

	v.Generator.Variables[varName] = memLocation
	v.Generator.VarOffset += 8

	return nil
}




func (v *ArmVisitor) VisitId(ctx *parser.IdContext) interface{} {
	varName := ctx.GetText()
	memLoc, ok := v.Generator.Variables[varName]
	if !ok {
		fmt.Printf("⚠️ Variable no encontrada: %s\n", varName)
		return nil
	}

	v.Generator.Comment(fmt.Sprintf("Accediendo variable %s", varName))
	v.Generator.Add(fmt.Sprintf("LDR X0, %s", memLoc)) // cargar a X0
	return nil
}


func (v *ArmVisitor) VisitValorexpr(ctx *parser.ValorexprContext) interface{} {
	if ctx.GetChildCount() == 1 {
		child := ctx.GetChild(0)
		if parseTree, ok := child.(antlr.ParseTree); ok {
			return v.Visit(parseTree)
		}
	}

	fmt.Println("⚠️ ValorExpr no reconocido:", ctx.GetText())
	return nil
}



func (v *ArmVisitor) VisitSumres(ctx *parser.SumresContext) interface{} {
    fmt.Println("➕➖ VisitSumres")

    left := v.Visit(ctx.GetChild(0).(antlr.ParseTree)).(string)
    right := v.Visit(ctx.GetChild(2).(antlr.ParseTree)).(string)
    result := v.Generator.NextTempReg()

    switch ctx.GetOp().GetText() {
    case "+":
        v.Generator.Comment(fmt.Sprintf("Suma: %s + %s", left, right))
        Add(v.Generator, result, left, right)
    case "-":
        v.Generator.Comment(fmt.Sprintf("Resta: %s - %s", left, right))
        Sub(v.Generator, result, left, right)
    default:
        fmt.Println("⚠️ Operador desconocido en sumres:", ctx.GetOp().GetText())
    }

    return result
}



