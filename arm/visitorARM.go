package arm

import (
	parser "compiler/parser"
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
)

type ArmString struct {
    Reg   string
    Label string
}

type FloatReg struct {
    Reg string // e.g., "v1"
}


// ArmVisitor recorre el árbol de análisis y genera código ARM64.
type ArmVisitor struct {
	parser.BaseVlangVisitor
	Generator  *ArmGenerator
	LastResult string
	StringData map[string]string
}


// NewArmVisitor construye un nuevo visitor para generación ARM.
func NewArmVisitor() *ArmVisitor {
	return &ArmVisitor{
		Generator: &ArmGenerator{
			Instructions: []string{},
			Variables:    make(map[string]SymbolInfo), // <-- ahora con tipo SymbolInfo
			Output:       "",
			StdLib:       NewStandardLibrary(),
			VarOffset:    0,
			tempRegs:     []string{X1, X2, X3, X4}, // evita X0 (lo usa syscall)
			tempIndex:    0,
			StringData:   make(map[string]string), // asegúrate de inicializar esto también
		},
	}
}


func (v *ArmVisitor) VisitValorNulo(ctx *parser.ValorNuloContext) interface{} {
	fmt.Println("🛑 Valor nulo detectado (nil)")
	return "nil"
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



func (v *ArmVisitor) VisitValorDecimal(ctx *parser.ValorDecimalContext) interface{} {
	text := ctx.DECIMAL().GetText()
	val, err := strconv.ParseFloat(text, 64)
	if err != nil {
		fmt.Println("⚠️ Error al convertir decimal:", text)
		return nil
	}

	label := v.Generator.GenerateStringLabel() // Etiqueta única
	v.Generator.AddDataDouble(label, val)      // Agrega .double a .data

	v.Generator.Comment(fmt.Sprintf("Carga decimal %f desde memoria", val))
	v.Generator.Add(fmt.Sprintf("ADR x1, %s", label)) // Dirección a x1
	v.Generator.Add("LDR d0, [x1]")                   // Carga double desde memoria

	return FloatReg{Reg: "d0"}
}





func (v *ArmVisitor) VisitValorCadena(ctx *parser.ValorCadenaContext) interface{} {
	text := ctx.CADENA().GetText()        // Incluye comillas: "hola"
	clean := text[1 : len(text)-1]        // Elimina las comillas

	label := v.Generator.GenerateStringLabel()
	v.Generator.AddData(label, clean)

	reg := v.Generator.NextTempReg()
	v.Generator.Comment(fmt.Sprintf("Cadena literal: %q", clean))
	v.Generator.Add(fmt.Sprintf("ADR %s, %s", reg, label)) // ✅ Cambio aquí

	v.Generator.StdLib.Use("print_string") // asegura inclusión en ensamblado

	return ArmString{Reg: reg, Label: label}
}








func (v *ArmVisitor) VisitValorBooleano(ctx *parser.ValorBooleanoContext) interface{} {
	text := ctx.BOOLEANO().GetText() // "true" o "false"

	// Crear etiqueta de string en .data
	label := v.Generator.GenerateStringLabel()
	v.Generator.AddData(label, text)

	// Registrar uso de print_string
	v.Generator.StdLib.Use("print_string")

	// Preparar registro con dirección del string
	reg := v.Generator.NextTempReg()
	v.Generator.Comment(fmt.Sprintf("Booleano literal: %s", text))
	v.Generator.Add(fmt.Sprintf("ADR %s, %s", reg, label))

	return ArmString{Reg: reg, Label: label}
}



func (v *ArmVisitor) VisitValorCaracter(ctx *parser.ValorCaracterContext) interface{} {
	text := ctx.CARACTER().GetText() // e.g. `'a'`
	if len(text) < 3 {
		fmt.Println("⚠️ Carácter inválido:", text)
		return nil
	}

	char := rune(text[1])
	reg := v.Generator.NextTempReg()
	v.Generator.Comment(fmt.Sprintf("Carácter '%c' ASCII=%d", char, char))
	v.Generator.Add(fmt.Sprintf("MOV %s, #%d", reg, char))
	return reg
}



// VisitPrintStatement traduce una instrucción println(...).
func (v *ArmVisitor) VisitPrintStatement(ctx *parser.PrintStatementContext) interface{} {
	fmt.Println("🖨️ VisitPrintStatement ARM")

	if ctx.Parametros() == nil {
		return nil
	}

	exprs := ctx.Parametros().AllExpresion()

	for i, expr := range exprs {
		val := v.Visit(expr)

		switch value := val.(type) {
		case string:
			v.Generator.Comment("Print entero")
			PrintInteger(v.Generator, value)

		case ArmString:
			v.Generator.Comment("Print cadena")
			v.Generator.Add(fmt.Sprintf("MOV X0, %s", value.Reg))
			v.Generator.Add("BL print_string")

		case FloatReg:
			v.Generator.Comment("Print float64")
			v.Generator.StdLib.Use("float1000")            // Para multiplicar decimales
			v.Generator.StdLib.Use("print_float")           // Lógica principal
			v.Generator.StdLib.Use("print_3digit_integer")  // NECESARIA para evitar error
			v.Generator.Add(fmt.Sprintf("FMOV D0, %s", value.Reg))
			v.Generator.Add("BL print_float")

		default:
			fmt.Println("⚠️ No se pudo imprimir, tipo no reconocido:", fmt.Sprintf("%T", val))
			v.Generator.Comment("Tipo no reconocido, omitiendo impresión")
		}

		// Agrega espacio si no es el último parámetro
		if i < len(exprs)-1 {
			spaceLabel := v.Generator.GenerateStringLabel()
			v.Generator.AddData(spaceLabel, " ")
			v.Generator.Add(fmt.Sprintf("ADR x1, %s", spaceLabel))
			v.Generator.Add("MOV X0, x1")
			v.Generator.Add("BL print_string")
		}
	}

	// Salto de línea después de println
	v.Generator.Comment("Salto de línea después de println")
	nlLabel := v.Generator.GenerateStringLabel()
	v.Generator.AddData(nlLabel, "\n")
	v.Generator.Add(fmt.Sprintf("ADR x1, %s", nlLabel))
	v.Generator.Add("MOV X0, x1")
	v.Generator.Add("BL print_string")

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

    // ✅ Agregar un RET al final de cada función
    v.Generator.Add("RET")
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

	offset := v.Generator.VarOffset
	memLocation := fmt.Sprintf("[SP, #%d]", offset)

	if ctx.Expresion() != nil {
		val := v.Visit(ctx.Expresion())

		var reg string
		var tipo string

		switch value := val.(type) {
		case string:
			reg = value
			tipo = "int" // por defecto, puedes mejorar esto más adelante con análisis semántico
		case ArmString:
			reg = value.Reg
			tipo = "string"
		default:
			fmt.Printf("⚠️ Error: expresión no retornó un registro válido para variable '%s'\n", varName)
			return nil
		}

		v.Generator.Add(fmt.Sprintf("STR %s, %s", reg, memLocation))
		v.Generator.Comment(fmt.Sprintf("Variable %s inicializada con valor en %s", varName, reg))

		// Guardar tipo y ubicación
		v.Generator.Variables[varName] = SymbolInfo{
			Location: memLocation,
			Type:     tipo,
		}
	} else {
		v.Generator.Add(fmt.Sprintf("MOV %s, #0", X0))
		v.Generator.Add(fmt.Sprintf("STR %s, %s", X0, memLocation))
		v.Generator.Comment(fmt.Sprintf("Variable %s declarada sin valor (inicializada en 0)", varName))

		// Por defecto, tipo int si no tiene valor inicial
		v.Generator.Variables[varName] = SymbolInfo{
			Location: memLocation,
			Type:     "int",
		}
	}

	v.Generator.VarOffset += 8
	return nil
}




func (v *ArmVisitor) VisitId(ctx *parser.IdContext) interface{} {
	varName := ctx.GetText()
	info, ok := v.Generator.Variables[varName]
	if !ok {
		fmt.Printf("⚠️ Variable no encontrada: %s\n", varName)
		return nil
	}

	reg := v.Generator.NextTempReg()
	v.Generator.Comment(fmt.Sprintf("Accediendo variable %s", varName))
	v.Generator.Add(fmt.Sprintf("LDR %s, %s", reg, info.Location))

	switch info.Type {
	case "string", "float64", "bool":
		// Si son valores que se imprimen como cadenas
		return ArmString{Reg: reg, Label: ""}
	default:
		// Enteros, etc.
		return reg
	}
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

	left := v.Visit(ctx.Expresion(0))
	right := v.Visit(ctx.Expresion(1))
	op := ctx.GetOp().GetText()

	switch l := left.(type) {

	// int + int o int - int
	case string:
		switch r := right.(type) {
		case string:
			result := v.Generator.NextTempReg()
			if op == "+" {
				v.Generator.Comment("Suma de enteros")
				Add(v.Generator, result, l, r)
			} else {
				v.Generator.Comment("Resta de enteros")
				Sub(v.Generator, result, l, r)
			}
			return result

		case FloatReg:
			tmpFloat := fmt.Sprintf("d%d", v.Generator.tempIndex)
			v.Generator.tempIndex++
			v.Generator.Comment("Convierte int a float64")
			v.Generator.Add(fmt.Sprintf("SCVTF %s, %s", tmpFloat, l))

			result := fmt.Sprintf("d%d", v.Generator.tempIndex)
			v.Generator.tempIndex++
			v.Generator.Comment("Suma int + float")
			if op == "+" {
				v.Generator.Add(fmt.Sprintf("FADD %s, %s, %s", result, tmpFloat, r.Reg))
			} else {
				v.Generator.Add(fmt.Sprintf("FSUB %s, %s, %s", result, tmpFloat, r.Reg))
			}
			return FloatReg{Reg: result}
		}

	// float + int, float + float, float - int, float - float
	case FloatReg:
		switch r := right.(type) {
		case FloatReg:
			result := fmt.Sprintf("d%d", v.Generator.tempIndex)
			v.Generator.tempIndex++
			v.Generator.Comment("Suma/Resta de float64")
			if op == "+" {
				v.Generator.Add(fmt.Sprintf("FADD %s, %s, %s", result, l.Reg, r.Reg))
			} else {
				v.Generator.Add(fmt.Sprintf("FSUB %s, %s, %s", result, l.Reg, r.Reg))
			}
			return FloatReg{Reg: result}

		case string:
			tmpFloat := fmt.Sprintf("d%d", v.Generator.tempIndex)
			v.Generator.tempIndex++
			v.Generator.Comment("Convierte int a float64")
			v.Generator.Add(fmt.Sprintf("SCVTF %s, %s", tmpFloat, r))

			result := fmt.Sprintf("d%d", v.Generator.tempIndex)
			v.Generator.tempIndex++
			v.Generator.Comment("Suma/Resta float + int")
			if op == "+" {
				v.Generator.Add(fmt.Sprintf("FADD %s, %s, %s", result, l.Reg, tmpFloat))
			} else {
				v.Generator.Add(fmt.Sprintf("FSUB %s, %s, %s", result, l.Reg, tmpFloat))
			}
			return FloatReg{Reg: result}
		}

	// string + string (concatenación)
	case ArmString:
		if r, ok := right.(ArmString); ok && op == "+" {
			leftContent := v.Generator.StringData[l.Label]
			rightContent := v.Generator.StringData[r.Label]
			content := leftContent + rightContent

			label := v.Generator.GenerateStringLabel()
			v.Generator.AddData(label, content)

			reg := v.Generator.NextTempReg()
			v.Generator.Add(fmt.Sprintf("ADR %s, %s", reg, label))

			return ArmString{Reg: reg, Label: label}
		}
	}

	// ⚠️ Caso no soportado
	fmt.Println("⚠️ Tipos no soportados para suma:", fmt.Sprintf("%T + %T", left, right))
	return nil
}





func (v *ArmVisitor) VisitMultdivmod(ctx *parser.MultdivmodContext) interface{} {
    fmt.Println("✖️➗ VisitMultdivmod")

    left := v.Visit(ctx.GetChild(0).(antlr.ParseTree)).(string)
    right := v.Visit(ctx.GetChild(2).(antlr.ParseTree)).(string)
    result := v.Generator.NextTempReg()

    switch ctx.GetOp().GetText() {
    case "*":
        v.Generator.Comment(fmt.Sprintf("Multiplicación: %s * %s", left, right))
        v.Generator.Add(fmt.Sprintf("MUL %s, %s, %s", result, left, right))
    case "/":
        v.Generator.Comment(fmt.Sprintf("División: %s / %s", left, right))
        v.Generator.Add(fmt.Sprintf("UDIV %s, %s, %s", result, left, right))
    case "%":
        v.Generator.Comment(fmt.Sprintf("Módulo: %s %% %s", left, right))
        // x26 = left - (left / right) * right
        v.Generator.Add(fmt.Sprintf("UDIV x25, %s, %s", left, right))    // x25 = left / right
        v.Generator.Add(fmt.Sprintf("MSUB %s, x25, %s, %s", result, right, left)) // result = left - (x25 * right)
    default:
        fmt.Println("⚠️ Operador desconocido en multdivmod:", ctx.GetOp().GetText())
    }

    return result
}
