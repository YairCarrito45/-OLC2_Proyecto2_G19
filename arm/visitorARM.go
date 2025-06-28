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

type BoolReg struct {
	Reg string
}


type FloatReg struct {
    Reg string // e.g., "v1"
}

type TypedValue struct {
	Reg  string // puede ser xN o dN
	Type string // "int", "float64", "string", etc.
}

type Symbol struct {
	Name string // nombre del registro o label asociado
	Type string // "int", "float64", "string", etc.
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

func (g *ArmGenerator) NextTempFloatReg() string {
	reg := fmt.Sprintf("d%d", g.tempFloatIndex)
	g.tempFloatIndex = (g.tempFloatIndex + 1) % 8 // usa 8 registros por ejemplo
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

	reg := v.Generator.NextTempFloatReg() // ✅ Usa registro flotante distinto

	v.Generator.Comment(fmt.Sprintf("Carga decimal %f desde memoria", val))
	v.Generator.Add(fmt.Sprintf("ADR x1, %s", label)) // Dirección a x1
	v.Generator.Add(fmt.Sprintf("LDR %s, [x1]", reg)) // Carga a reg

	return FloatReg{Reg: reg}
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


/// visit para controlar parentesis
func (v *ArmVisitor) VisitParentesisexpre(ctx *parser.ParentesisexpreContext) interface{} {
	// Simplemente visita la expresión interna, conserva su tipo y resultado
	return v.Visit(ctx.Expresion())
}






func (v *ArmVisitor) VisitValorBooleano(ctx *parser.ValorBooleanoContext) interface{} {
	text := ctx.BOOLEANO().GetText() // "true" o "false"

	reg := v.Generator.NextTempReg()
	if text == "true" {
		v.Generator.Comment("Literal booleano: true")
		v.Generator.Add(fmt.Sprintf("MOV %s, #1", reg))
	} else {
		v.Generator.Comment("Literal booleano: false")
		v.Generator.Add(fmt.Sprintf("MOV %s, #0", reg))
	}

	return BoolReg{Reg: reg}
}


func (v *ArmVisitor) VisitUnario(ctx *parser.UnarioContext) interface{} {
	fmt.Println("➖ VisitUnario")

	op := ctx.GetOp().GetText()
	val := v.Visit(ctx.Expresion())

	switch op {
	case "-":
		switch t := val.(type) {
		case string: // entero
			result := v.Generator.NextTempReg()
			v.Generator.Comment("Negación unaria de entero")
			v.Generator.Add(fmt.Sprintf("NEG %s, %s", result, t))
			return result

		case FloatReg: // float
			result := v.Generator.NextTempFloatReg()
			v.Generator.Comment("Negación unaria de float64")
			v.Generator.Add(fmt.Sprintf("FNEG %s, %s", result, t.Reg))
			return FloatReg{Reg: result}
		}

	case "!":
		switch b := val.(type) {
		case BoolReg:
			result := v.Generator.NextTempReg()
			v.Generator.Comment("Negación lógica de booleano (!)")
			v.Generator.Add(fmt.Sprintf("EOR %s, %s, #1", result, b.Reg))
			return BoolReg{Reg: result}

		default:
			v.Generator.Comment("❌ Operando no booleano para negación lógica")
		}
	}

	fmt.Println("⚠️ Tipo o operador no compatible en negación unaria:", op)
	return nil
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
		if val == nil {
			v.Generator.Comment("⚠️ Valor nulo, omitiendo impresión")
			continue
		}

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
			v.Generator.StdLib.Use("float1000")            // Constante para 3 decimales
			v.Generator.StdLib.Use("print_float")           // Función principal
			v.Generator.StdLib.Use("print_3digit_integer")  // Necesaria para los decimales
			v.Generator.Add(fmt.Sprintf("FMOV D0, %s", value.Reg))
			v.Generator.Add("BL print_float")

		case BoolReg:
			v.Generator.Comment("Print booleano")

			// Generar etiquetas únicas
			labelPrefix := v.Generator.GenerateLabel("bool")
			trueJump := labelPrefix + "_true_jump"
			falseJump := labelPrefix + "_false_jump"
			endLabel := labelPrefix + "_end"

			trueStrLabel := labelPrefix + "_true_str"
			falseStrLabel := labelPrefix + "_false_str"

			// Agregar datos en la sección .data
			v.Generator.AddData(trueStrLabel, "true")
			v.Generator.AddData(falseStrLabel, "false")

			// Comparar el valor del registro con 1
			v.Generator.Add(fmt.Sprintf("CMP %s, #1", value.Reg))
			v.Generator.Add(fmt.Sprintf("BEQ %s", trueJump))
			v.Generator.Add(fmt.Sprintf("B %s", falseJump))

			// true
			v.Generator.AddLabel(trueJump)
			v.Generator.Add(fmt.Sprintf("ADR X0, %s", trueStrLabel))
			v.Generator.Add("BL print_string")
			v.Generator.Add(fmt.Sprintf("B %s", endLabel))

			// false
			v.Generator.AddLabel(falseJump)
			v.Generator.Add(fmt.Sprintf("ADR X0, %s", falseStrLabel))
			v.Generator.Add("BL print_string")

			v.Generator.AddLabel(endLabel)




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

var floatLabelCounter int = 0

func (g *ArmGenerator) GenerateConstLabel() string {
	label := fmt.Sprintf("float_const_%d", floatLabelCounter)
	floatLabelCounter++
	return label
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
			tipo = "int"
		case ArmString:
			reg = value.Reg
			tipo = "string"
		case FloatReg:
			reg = value.Reg
			tipo = "float64"
		default:
			fmt.Printf("⚠️ Error: expresión no retornó un registro válido para variable '%s'\n", varName)
			return nil
		}

		v.Generator.Add(fmt.Sprintf("STR %s, %s", reg, memLocation))
		v.Generator.Comment(fmt.Sprintf("Variable %s inicializada con valor en %s", varName, reg))

		v.Generator.Variables[varName] = SymbolInfo{
			Location: memLocation,
			Type:     tipo,
		}
	} else {
		var tipo string
		if ctx.Tipo() != nil {
			tipoTexto := ctx.Tipo().GetText()
			switch tipoTexto {
			case "float64":
				tipo = "float64"
				label := v.Generator.GenerateConstLabel()
				v.Generator.AddDataDouble(label, 0.0)
				v.Generator.Add(fmt.Sprintf("ADR x1, %s", label))
				v.Generator.Add("LDR d0, [x1]")
				v.Generator.Add(fmt.Sprintf("STR d0, %s", memLocation))

			case "string":
				tipo = "string"
				label := v.Generator.GenerateStringLabel()
				v.Generator.AddData(label, "")
				v.Generator.Add(fmt.Sprintf("ADR x1, %s", label))
				v.Generator.Add(fmt.Sprintf("STR x1, %s", memLocation))
			default:
				tipo = "int"
				v.Generator.Add(fmt.Sprintf("MOV %s, #0", X0))
				v.Generator.Add(fmt.Sprintf("STR %s, %s", X0, memLocation))
			}
		} else {
			tipo = "int"
			v.Generator.Add(fmt.Sprintf("MOV %s, #0", X0))
			v.Generator.Add(fmt.Sprintf("STR %s, %s", X0, memLocation))
		}

		v.Generator.Comment(fmt.Sprintf("Variable %s declarada sin valor (tipo: %s)", varName, tipo))
		v.Generator.Variables[varName] = SymbolInfo{
			Location: memLocation,
			Type:     tipo,
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

	v.Generator.Comment(fmt.Sprintf("Accediendo variable %s", varName))

	switch info.Type {
	case "float64":
		reg := v.Generator.NextTempFloatReg()
		v.Generator.Add(fmt.Sprintf("LDR %s, %s", reg, info.Location))
		return FloatReg{Reg: reg}

	case "string":
		reg := v.Generator.NextTempReg()
		v.Generator.Add(fmt.Sprintf("LDR %s, %s", reg, info.Location))
		return ArmString{Reg: reg, Label: ""}

	case "bool":
		reg := v.Generator.NextTempReg()
		v.Generator.Add(fmt.Sprintf("LDR %s, %s", reg, info.Location))
		return BoolReg{Reg: reg}

	default:
		// Asumimos int, char, etc.
		reg := v.Generator.NextTempReg()
		v.Generator.Add(fmt.Sprintf("LDR %s, %s", reg, info.Location))
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
			fmt.Printf("🧮 %s %s %s (int)\n", l, op, r)
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
			fmt.Printf("🧮 %s %s %s (int y float)\n", l, op, r.Reg)
			tmpFloat := v.Generator.NextTempFloatReg()
			v.Generator.Comment("Convierte int a float64")
			v.Generator.Add(fmt.Sprintf("SCVTF %s, %s", tmpFloat, l))

			result := v.Generator.NextTempFloatReg()
			v.Generator.Comment("Suma/Resta int y float")
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
			fmt.Printf("🧮 %s %s %s (float)\n", l.Reg, op, r.Reg)
			result := v.Generator.NextTempFloatReg()
			v.Generator.Comment("Suma/Resta de float64")
			if op == "+" {
				v.Generator.Add(fmt.Sprintf("FADD %s, %s, %s", result, l.Reg, r.Reg))
			} else {
				v.Generator.Add(fmt.Sprintf("FSUB %s, %s, %s", result, l.Reg, r.Reg))
			}
			return FloatReg{Reg: result}

		case string:
			fmt.Printf("🧮 %s %s %s (float y int)\n", l.Reg, op, r)
			tmpFloat := v.Generator.NextTempFloatReg()
			v.Generator.Comment("Convierte int a float64")
			v.Generator.Add(fmt.Sprintf("SCVTF %s, %s", tmpFloat, r))

			result := v.Generator.NextTempFloatReg()
			v.Generator.Comment("Suma/Resta float y int")
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
			fmt.Printf("🧮 \"%s\" + \"%s\" (concatenación)\n", l.Label, r.Label)
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
	fmt.Printf("⚠️ Tipos no soportados para suma o resta: %T %s %T\n", left, op, right)
	return nil
}


func (v *ArmVisitor) VisitMultdivmod(ctx *parser.MultdivmodContext) interface{} {
	fmt.Println("✖️➗ VisitMultdivmod")

	left := v.Visit(ctx.GetChild(0).(antlr.ParseTree))
	right := v.Visit(ctx.GetChild(2).(antlr.ParseTree))
	op := ctx.GetOp().GetText()

	switch l := left.(type) {

	// 👉 int * int, int / int, int % int, int * float, int / float
	case string:
		switch r := right.(type) {

		case string:
			result := v.Generator.NextTempReg()
			switch op {
			case "*":
				v.Generator.Comment("Multiplicación de enteros")
				v.Generator.Add(fmt.Sprintf("MUL %s, %s, %s", result, l, r))
			case "/":
				v.Generator.Comment("División de enteros")
				v.Generator.Add(fmt.Sprintf("UDIV %s, %s, %s", result, l, r))
			case "%":
				v.Generator.Comment("Módulo de enteros")
				tmp := v.Generator.NextTempReg()
				v.Generator.Add(fmt.Sprintf("UDIV %s, %s, %s", tmp, l, r))
				v.Generator.Add(fmt.Sprintf("MSUB %s, %s, %s, %s", result, tmp, r, l))
			}
			return result

		case FloatReg:
			tmpFloat := v.Generator.NextTempFloatReg()
			v.Generator.Comment("Convierte int a float64")
			v.Generator.Add(fmt.Sprintf("SCVTF %s, %s", tmpFloat, l))

			result := v.Generator.NextTempFloatReg()
			switch op {
			case "*":
				v.Generator.Comment("Multiplicación int * float")
				v.Generator.Add(fmt.Sprintf("FMUL %s, %s, %s", result, tmpFloat, r.Reg))
			case "/":
				v.Generator.Comment("División int / float")
				v.Generator.Add(fmt.Sprintf("FDIV %s, %s, %s", result, tmpFloat, r.Reg))
			default:
				v.Generator.Comment("⚠️ Módulo no válido con float")
			}
			return FloatReg{Reg: result}
		}

	// 👉 float * float, float / float, float * int, float / int
	case FloatReg:
		switch r := right.(type) {

		case FloatReg:
			result := v.Generator.NextTempFloatReg()
			switch op {
			case "*":
				v.Generator.Comment("Multiplicación de flotantes")
				v.Generator.Add(fmt.Sprintf("FMUL %s, %s, %s", result, l.Reg, r.Reg))
			case "/":
				v.Generator.Comment("División de flotantes")
				v.Generator.Add(fmt.Sprintf("FDIV %s, %s, %s", result, l.Reg, r.Reg))
			default:
				v.Generator.Comment("⚠️ Módulo no válido con float")
			}
			return FloatReg{Reg: result}

		case string:
			tmpFloat := v.Generator.NextTempFloatReg()
			v.Generator.Comment("Convierte int a float64")
			v.Generator.Add(fmt.Sprintf("SCVTF %s, %s", tmpFloat, r))

			result := v.Generator.NextTempFloatReg()
			switch op {
			case "*":
				v.Generator.Comment("Multiplicación float * int")
				v.Generator.Add(fmt.Sprintf("FMUL %s, %s, %s", result, l.Reg, tmpFloat))
			case "/":
				v.Generator.Comment("División float / int")
				v.Generator.Add(fmt.Sprintf("FDIV %s, %s, %s", result, l.Reg, tmpFloat))
			default:
				v.Generator.Comment("⚠️ Módulo no válido con float")
			}
			return FloatReg{Reg: result}
		}
	}

	fmt.Println("⚠️ Tipos no compatibles en multiplicación/división")
	return nil
}



// asignacion compuesta suma (incremento)
func (v *ArmVisitor) VisitAsignacionCompuestaSuma(ctx *parser.AsignacionCompuestaSumaContext) interface{} {
	id := ctx.ID().GetText()
	right := v.Visit(ctx.Expresion())

	// Obtener tipo y dirección
	symInfo, ok := v.Generator.Variables[id]
	if !ok {
		v.Generator.Comment(fmt.Sprintf("❌ Variable '%s' no declarada", id))
		return nil
	}

	switch symInfo.Type {
	case "int":
		rightReg, ok := right.(string)
		if !ok {
			v.Generator.Comment("❌ Incompatibilidad de tipo en += int")
			return nil
		}
		current := v.Generator.NextTempReg()
		v.Generator.Add(fmt.Sprintf("LDR %s, %s", current, symInfo.Location))
		result := v.Generator.NextTempReg()
		v.Generator.Add(fmt.Sprintf("ADD %s, %s, %s", result, current, rightReg))
		v.Generator.Add(fmt.Sprintf("STR %s, %s", result, symInfo.Location))

	case "float64":
		rightF, ok := right.(FloatReg)
		if !ok {
			v.Generator.Comment("❌ Incompatibilidad de tipo en += float64")
			return nil
		}
		tmp := v.Generator.NextTempReg()
		v.Generator.Add(fmt.Sprintf("LDR %s, %s", tmp, symInfo.Location))
		current := v.Generator.NextTempFloatReg()
		v.Generator.Add(fmt.Sprintf("LDR %s, [%s]", current, tmp))
		result := v.Generator.NextTempFloatReg()
		v.Generator.Add(fmt.Sprintf("FADD %s, %s, %s", result, current, rightF.Reg))
		v.Generator.Add(fmt.Sprintf("STR %s, [%s]", result, tmp))

	case "string":
		rightStr, ok := right.(ArmString)
		if !ok {
			v.Generator.Comment("❌ Incompatibilidad de tipo en += string")
			return nil
		}
		v.Generator.StdLib.Use("concat_strings")
		v.Generator.Add(fmt.Sprintf("MOV X0, %s", symInfo.Location)) // ❓ Si Location es dirección de string
		v.Generator.Add(fmt.Sprintf("MOV X1, %s", rightStr.Reg))
		v.Generator.Add("BL concat_strings")
		// X0 tiene el resultado, si quieres puedes guardarlo en una nueva dirección o actualizar la existente

	default:
		v.Generator.Comment("❌ Tipo no soportado en '+='")
	}

	return nil
}



// asignacion compuesta resta (decremento)

func (v *ArmVisitor) VisitAsignacionCompuestaResta(ctx *parser.AsignacionCompuestaRestaContext) interface{} {
	id := ctx.ID().GetText()
	right := v.Visit(ctx.Expresion())

	sym, ok := v.Generator.Variables[id]
	if !ok {
		v.Generator.Comment(fmt.Sprintf("❌ Variable '%s' no declarada", id))
		return nil
	}

	switch sym.Type {
	case "int":
		rightReg, ok := right.(string)
		if !ok {
			v.Generator.Comment("❌ Incompatibilidad de tipo en -= int")
			return nil
		}
		result := v.Generator.NextTempReg()
		v.Generator.Add(fmt.Sprintf("LDR %s, %s", result, sym.Location))         // carga valor actual
		v.Generator.Add(fmt.Sprintf("SUB %s, %s, %s", result, result, rightReg)) // resta
		v.Generator.Add(fmt.Sprintf("STR %s, %s", result, sym.Location))         // guarda de nuevo

	case "float64":
		rightF, ok := right.(FloatReg)
		if !ok {
			v.Generator.Comment("❌ Incompatibilidad de tipo en -= float64")
			return nil
		}
		left := v.Generator.NextTempFloatReg()
		v.Generator.Add(fmt.Sprintf("LDR %s, %s", left, sym.Location)) // carga valor actual
		res := v.Generator.NextTempFloatReg()
		v.Generator.Add(fmt.Sprintf("FSUB %s, %s, %s", res, left, rightF.Reg)) // resta float
		v.Generator.Add(fmt.Sprintf("STR %s, %s", res, sym.Location))          // guarda de nuevo

	default:
		v.Generator.Comment("❌ Tipo no soportado en '-='")
	}

	return nil
}


// visit de operaciones comparativas
func (v *ArmVisitor) VisitIgualdad(ctx *parser.IgualdadContext) interface{} {
	left := v.Visit(ctx.Expresion(0))
	right := v.Visit(ctx.Expresion(1))
	op := ctx.GetOp().GetText() // "==" o "!="

	resultReg := v.Generator.NextTempReg()
	v.Generator.Comment(fmt.Sprintf("Comparación %s", op))

	switch l := left.(type) {

	// Comparación de enteros
	case string:
		switch r := right.(type) {
		case string:
			v.Generator.Add(fmt.Sprintf("CMP %s, %s", l, r))
			if op == "==" {
				v.Generator.Add(fmt.Sprintf("CSET %s, EQ", resultReg))
			} else {
				v.Generator.Add(fmt.Sprintf("CSET %s, NE", resultReg))
			}
			return resultReg

		case FloatReg:
			tmpf := v.Generator.NextTempFloatReg()
			v.Generator.Comment("Convirtiendo entero a float64 para comparación")
			v.Generator.Add(fmt.Sprintf("SCVTF %s, %s", tmpf, l))
			v.Generator.Add(fmt.Sprintf("FCMP %s, %s", tmpf, r.Reg))
			if op == "==" {
				v.Generator.Add(fmt.Sprintf("CSET %s, EQ", resultReg))
			} else {
				v.Generator.Add(fmt.Sprintf("CSET %s, NE", resultReg))
			}
			return resultReg
		}

	// Comparación de floats
	case FloatReg:
		switch r := right.(type) {
		case FloatReg:
			v.Generator.Add(fmt.Sprintf("FCMP %s, %s", l.Reg, r.Reg))
			if op == "==" {
				v.Generator.Add(fmt.Sprintf("CSET %s, EQ", resultReg))
			} else {
				v.Generator.Add(fmt.Sprintf("CSET %s, NE", resultReg))
			}
			return resultReg

		case string:
			tmpf := v.Generator.NextTempFloatReg()
			v.Generator.Comment("Convirtiendo entero a float64 para comparación")
			v.Generator.Add(fmt.Sprintf("SCVTF %s, %s", tmpf, r))
			v.Generator.Add(fmt.Sprintf("FCMP %s, %s", l.Reg, tmpf))
			if op == "==" {
				v.Generator.Add(fmt.Sprintf("CSET %s, EQ", resultReg))
			} else {
				v.Generator.Add(fmt.Sprintf("CSET %s, NE", resultReg))
			}
			return resultReg
		}

	// Comparación de strings (por dirección de memoria)
	case ArmString:
		if r, ok := right.(ArmString); ok {
			v.Generator.Add(fmt.Sprintf("CMP %s, %s", l.Reg, r.Reg))
			if op == "==" {
				v.Generator.Add(fmt.Sprintf("CSET %s, EQ", resultReg))
			} else {
				v.Generator.Add(fmt.Sprintf("CSET %s, NE", resultReg))
			}
			return resultReg
		}

	// Comparación de booleanos
	case BoolReg:
		if r, ok := right.(BoolReg); ok {
			v.Generator.Add(fmt.Sprintf("CMP %s, %s", l.Reg, r.Reg))
			if op == "==" {
				v.Generator.Add(fmt.Sprintf("CSET %s, EQ", resultReg))
			} else {
				v.Generator.Add(fmt.Sprintf("CSET %s, NE", resultReg))
			}
			return resultReg
		}
	}

	v.Generator.Comment("❌ Tipos incompatibles en comparación")
	return nil
}





func (v *ArmVisitor) VisitRelacionales(ctx *parser.RelacionalesContext) interface{} {
	left := v.Visit(ctx.Expresion(0))
	right := v.Visit(ctx.Expresion(1))
	op := ctx.GetOp().GetText()

	// Comparación entre enteros
	if lreg, ok := left.(string); ok {
		if rreg, ok2 := right.(string); ok2 {
			result := v.Generator.NextTempReg()
			trueLabel := v.Generator.GenerateLabel("rel_true")
			endLabel := v.Generator.GenerateLabel("rel_end")

			v.Generator.Comment("Comparación relacional entre enteros")
			v.Generator.Add(fmt.Sprintf("CMP %s, %s", lreg, rreg))

			switch op {
			case ">":
				v.Generator.Add(fmt.Sprintf("BGT %s", trueLabel))
			case "<":
				v.Generator.Add(fmt.Sprintf("BLT %s", trueLabel))
			case ">=":
				v.Generator.Add(fmt.Sprintf("BGE %s", trueLabel))
			case "<=":
				v.Generator.Add(fmt.Sprintf("BLE %s", trueLabel))
			}

			v.Generator.Add(fmt.Sprintf("MOV %s, #0", result))
			v.Generator.Add(fmt.Sprintf("B %s", endLabel))

			v.Generator.AddLabel(trueLabel)
			v.Generator.Add(fmt.Sprintf("MOV %s, #1", result))

			v.Generator.AddLabel(endLabel)
			return result
		}
	}

	// Comparación entre float64
	if lf, ok := left.(FloatReg); ok {
		if rf, ok2 := right.(FloatReg); ok2 {
			return v.compareFloatRel(op, lf, rf)
		}
	}

	// int vs float64
	if lreg, ok := left.(string); ok {
		if rf, ok2 := right.(FloatReg); ok2 {
			tmpf := v.Generator.NextTempFloatReg()
			v.Generator.Comment("Convirtiendo entero a float64 para comparar")
			v.Generator.Add(fmt.Sprintf("SCVTF %s, %s", tmpf, lreg))
			return v.compareFloatRel(op, FloatReg{Reg: tmpf}, rf)
		}
	}

	// float64 vs int
	if lf, ok := left.(FloatReg); ok {
		if rreg, ok2 := right.(string); ok2 {
			tmpf := v.Generator.NextTempFloatReg()
			v.Generator.Comment("Convirtiendo entero a float64 para comparar")
			v.Generator.Add(fmt.Sprintf("SCVTF %s, %s", tmpf, rreg))
			return v.compareFloatRel(op, lf, FloatReg{Reg: tmpf})
		}
	}

	v.Generator.Comment("❌ Tipos incompatibles en operación relacional")
	return nil
}


func (v *ArmVisitor) compareFloatRel(op string, left FloatReg, right FloatReg) interface{} {
	result := v.Generator.NextTempReg()
	trueLabel := v.Generator.GenerateLabel("relf_true")
	endLabel := v.Generator.GenerateLabel("relf_end")

	v.Generator.Comment("Comparación relacional entre float64")
	v.Generator.Add(fmt.Sprintf("FCMP %s, %s", left.Reg, right.Reg))

	switch op {
	case ">":
		v.Generator.Add(fmt.Sprintf("BGT %s", trueLabel))
	case "<":
		v.Generator.Add(fmt.Sprintf("BLT %s", trueLabel))
	case ">=":
		v.Generator.Add(fmt.Sprintf("BGE %s", trueLabel))
	case "<=":
		v.Generator.Add(fmt.Sprintf("BLE %s", trueLabel))
	}

	v.Generator.Add(fmt.Sprintf("MOV %s, #0", result))
	v.Generator.Add(fmt.Sprintf("B %s", endLabel))

	v.Generator.AddLabel(trueLabel)
	v.Generator.Add(fmt.Sprintf("MOV %s, #1", result))

	v.Generator.AddLabel(endLabel)
	return result
}



// visitor operaciones logicas

func (v *ArmVisitor) VisitAndExpr(ctx *parser.AndExprContext) interface{} {
	left := v.Visit(ctx.Expresion(0))
	right := v.Visit(ctx.Expresion(1))

	lb, lok := left.(BoolReg)
	rb, rok := right.(BoolReg)
	if !lok || !rok {
		v.Generator.Comment("❌ Operadores no booleanos para &&")
		return nil
	}

	result := v.Generator.NextTempReg()
	v.Generator.Comment("Operador lógico AND (&&)")
	v.Generator.Add(fmt.Sprintf("AND %s, %s, %s", result, lb.Reg, rb.Reg))

	return BoolReg{Reg: result}
}



func (v *ArmVisitor) VisitOrExpr(ctx *parser.OrExprContext) interface{} {
	left := v.Visit(ctx.Expresion(0))
	right := v.Visit(ctx.Expresion(1))

	lb, lok := left.(BoolReg)
	rb, rok := right.(BoolReg)
	if !lok || !rok {
		v.Generator.Comment("❌ Operadores no booleanos para ||")
		return nil
	}

	result := v.Generator.NextTempReg()
	v.Generator.Comment("Operador lógico OR (||)")
	v.Generator.Add(fmt.Sprintf("ORR %s, %s, %s", result, lb.Reg, rb.Reg))

	return BoolReg{Reg: result}
}
