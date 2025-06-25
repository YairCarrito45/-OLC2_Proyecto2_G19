package arm

import (
	"fmt"
	parser "compiler/parser"
	"github.com/antlr4-go/antlr/v4"
)

// ArmVisitor recorre el árbol de análisis y genera código ARM64.
type ArmVisitor struct {
	parser.BaseVlangVisitor
	Generator *ArmGenerator // Generador de instrucciones ARM
}

// NewArmVisitor construye un nuevo visitor para generación ARM.
func NewArmVisitor() *ArmVisitor {
	return &ArmVisitor{
		Generator: &ArmGenerator{
			Instructions: []string{},
			Variables:    make(map[string]string),
			Output:       "",
			StdLib:       NewStandardLibrary(),
		},
	}
}

// Visit permite que ArmVisitor recorra cualquier árbol ANTLR.
func (v *ArmVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

// VisitBloque ejecuta las sentencias dentro de un bloque { ... }
func (v *ArmVisitor) VisitBloque(ctx *parser.BloqueContext) interface{} {
	fmt.Println("📦 Visitando bloque")

	for _, decl := range ctx.AllDeclaraciones() {
		v.Visit(decl)
	}
	return nil
}

func (v *ArmVisitor) VisitDeclaraciones(ctx *parser.DeclaracionesContext) interface{} {
	switch {
	case ctx.Stmt() != nil:
		return v.Visit(ctx.Stmt())
	case ctx.Funcion() != nil:
		return v.Visit(ctx.Funcion())
	case ctx.FuncionStruct() != nil:
		return v.Visit(ctx.FuncionStruct())
	case ctx.VarDcl() != nil:
		return v.Visit(ctx.VarDcl())
	case ctx.StructDecl() != nil:
		return v.Visit(ctx.StructDecl())
	default:
		return nil
	}
}


// VisitPrograma recorre todas las declaraciones del programa.
func (v *ArmVisitor) VisitPrograma(ctx *parser.ProgramaContext) interface{} {
	for _, decl := range ctx.AllDeclaraciones() {
        if fn := decl.Funcion(); fn != nil && fn.ID().GetText() == "main" {
            fmt.Println("🔁 Generando código para función main")
            v.Generator.Add("main:") // Etiqueta ARM
            v.Visit(fn)
        }
	}
	return nil
}







// VisitValorEntero genera código para una literal entera.
func (v *ArmVisitor) VisitValorEntero(ctx *parser.ValorEnteroContext) interface{} {
	value := ctx.GetText()
	fmt.Println("Visiting IntLiteral:", value)
	v.Generator.Comment("Literal agregada: " + value)
	v.Generator.Add("MOV X0, #" + value) // Mueve el valor a X0
	v.Generator.Comment("Valor movido a X0: " + value)
	return nil
}

// VisitPrintStatement traduce una instrucción println(...).
func (v *ArmVisitor) VisitPrintStatement(ctx *parser.PrintStatementContext) interface{} {
	if ctx.Parametros() != nil && ctx.Parametros().GetChildCount() > 0 {
		// Cast necesario: GetChild devuelve Tree, pero necesitamos ParseTree
		expr, ok := ctx.Parametros().GetChild(0).(antlr.ParseTree)
		if !ok {
			fmt.Println("⚠️ No se pudo castear el parámetro a ParseTree")
			return nil
		}

		v.Visit(expr)

		v.Generator.Comment("Print statement")
		v.Generator.StdLib.Use("print_integer")
		v.Generator.Add("BL print_integer")
	}
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
	return v.Visit(ctx.Expresion())
}


func (v *ArmVisitor) VisitFuncion(ctx *parser.FuncionContext) interface{} {
	funcName := ctx.ID().GetText()
	fmt.Println("🔧 Generando función:", funcName)

	// Marcar como etiqueta (opcional si solo tienes main)
	v.Generator.Add(fmt.Sprintf("%s:", funcName))

	// Visitar cuerpo
	return v.Visit(ctx.Bloque())
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
		v.Visit(ctx.Expresion()) // Resultado se espera en X0
		v.Generator.Add(fmt.Sprintf("STR X0, %s", memLocation)) // Guardar en memoria
		v.Generator.Comment(fmt.Sprintf("Variable %s inicializada con valor en X0", varName))
	} else {
		v.Generator.Add("MOV X0, #0")
		v.Generator.Add(fmt.Sprintf("STR X0, %s", memLocation))
		v.Generator.Comment(fmt.Sprintf("Variable %s declarada sin valor (inicializada en 0)", varName))
	}

	v.Generator.Variables[varName] = memLocation
	v.Generator.VarOffset += 8 // siguiente variable irá 8 bytes más adelante

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
