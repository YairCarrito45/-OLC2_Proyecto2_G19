package repl

import (
    parser "compiler/parser"
    "compiler/value"
    "fmt"
    "log"
    "strconv"
    "strings"

    "github.com/antlr4-go/antlr/v4"
)


type SymbolInfo struct {
	ID      string
	Tipo    string
	Valor   string
	Ambito  string
	Linea   int
	Columna int
}

type ReplVisitor struct {
	parser.BaseVlangVisitor
	CurrentScope *BaseScope
	ErrorTable   *ErrorTable
	Console      *Console
	Symbols      []SymbolInfo 
}

// Constructor
func NewReplVisitor(console *Console) *ReplVisitor {
	return &ReplVisitor{
		CurrentScope: NewBaseScope("global", nil),
		ErrorTable:   NewErrorTable(),
		Console:      console,
	}
}

func (v *ReplVisitor) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *antlr.ErrorNodeImpl:
		log.Fatal(val.GetText())
		return nil
	default:
		return tree.Accept(v)
	}
}




func (v *ReplVisitor) VisitPrograma(ctx *parser.ProgramaContext) interface{} {
	// Limpiar funciones globales antes de registrar nuevamente
	value.GlobalFunctions = map[string]*value.Function{}

	// Procesar todas las declaraciones
	for _, decl := range ctx.AllDeclaraciones() {
		v.Visit(decl)
	}

	// Ejecutar main si existe
	if mainFunc, ok := value.GlobalFunctions["main"]; ok {
		fmt.Println("🟢 antes Ejecutando main()")

		scope := NewBaseScope("main", v.CurrentScope)
		oldScope := v.CurrentScope
		v.CurrentScope = scope

		fmt.Println("🟢 Ejecutando main()")
		v.Visit(mainFunc.Body)

		v.CurrentScope = oldScope
		v.Console.Show()
	} else {
		fmt.Println("❌ Función 'main' no encontrada")
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "No se encontró la función 'main'")
	}


	return nil
}








// Manejo de declaraciones
func (v *ReplVisitor) VisitDeclaraciones(ctx *parser.DeclaracionesContext) interface{} {
	if ctx.VarDcl() != nil {
		return v.Visit(ctx.VarDcl())
	}
	if ctx.StructDecl() != nil {
		return v.Visit(ctx.StructDecl())
	}
	if ctx.Funcion() != nil {
		return v.Visit(ctx.Funcion()) // ✅ REGISTRA FUNCIONES GLOBALES
	}
	if ctx.FuncionStruct() != nil {
		return v.Visit(ctx.FuncionStruct()) // ✅ REGISTRA FUNCIONES DE STRUCT
	}
	if ctx.Stmt() != nil {
		return v.Visit(ctx.Stmt())
	}
	return nil
}



// Declaración de variable
func (v *ReplVisitor) VisitVariableDeclaration(ctx *parser.VariableDeclarationContext) interface{} {
	id := ctx.ID().GetText()

	// Verificar si es una palabra reservada
	if isReservedWord(id) {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), fmt.Sprintf("'%s' es una palabra reservada", id))
		return nil
	}

	var tipo string
	var val value.IVOR

	// Obtener tipo explícito si lo hay
	if ctx.Tipo() != nil {
		rawType := ctx.Tipo().GetText()
		tipo = NormalizeTypeName(rawType)

		if tipo == "" {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "Tipo inválido: "+rawType)
			return nil
		}
	}

	// Evaluar expresión si existe
	if ctx.Expresion() != nil {
		val = v.Visit(ctx.Expresion()).(value.IVOR)

		if tipo == "" {
			tipo = val.Type()
		} else if tipo != val.Type() {
			// Manejo especial para slices/matrices: comparar tipo base
			if strings.HasPrefix(tipo, "slice_") && strings.HasPrefix(val.Type(), "slice_") {
				if !validateNestedSliceType(val, tipo) {
					v.ErrorTable.NewSemanticError(ctx.GetStart(),
						fmt.Sprintf("Los tipos de elementos no coinciden: se esperaba '%s' pero se recibió '%s'", tipo, val.Type()))
					return nil
				}
			} else {
				v.ErrorTable.NewSemanticError(ctx.GetStart(),
					fmt.Sprintf("No se puede asignar tipo '%s' a variable '%s' de tipo '%s'", val.Type(), id, tipo))
				return nil
			}
		}
	} else {
		// Valor por defecto si no hay asignación
		if tipo == "" {
			tipo = value.IVOR_INT
		}
		val = defaultValueForType(tipo)
	}

	v.Symbols = append(v.Symbols, SymbolInfo{
		ID:      id,
		Tipo:    tipo,
		Valor:   fmt.Sprintf("%v", val.Value()), // Convertir valor a string para mostrar
		Ambito:  v.CurrentScope.Name,
		Linea:   ctx.GetStart().GetLine(),
		Columna: ctx.GetStart().GetColumn(),
	})


	// Agregar la variable al entorno
	_, msg := v.CurrentScope.AddVariable(id, tipo, val, false, true, ctx.GetStart())
	if msg != "" {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
	}
	return nil
}




// visit print

func (v *ReplVisitor) VisitPrintStatement(ctx *parser.PrintStatementContext) interface{} {
	fmt.Println("🖨️ VisitPrintStatement") // ← depuración

	if ctx.Parametros() == nil {
		v.Console.Print("\n")
		return nil
	}

	var outputs []string
	for _, expr := range ctx.Parametros().AllExpresion() {
		val := v.Visit(expr)
		if ivor, ok := val.(value.IVOR); ok {
			switch v := ivor.(type) {
			case *value.SliceValue:
				outputs = append(outputs, formatSliceRecursively(v))
			case *value.StructInstance:
				outputs = append(outputs, v.String())
			default:
				// ✅ Esto imprime int, string, bool, etc.
				outputs = append(outputs, fmt.Sprintf("%v", ivor.Value()))
			}
		} else {
			outputs = append(outputs, "nil")
		}
	}

	v.Console.Print(strings.Join(outputs, " ") + "\n")
	return nil
}





func (v *ReplVisitor) VisitValorexpr(ctx *parser.ValorexprContext) interface{} {
	if ctx.Valor() != nil {
		return v.Visit(ctx.Valor())
	}
	return nil 
}

/// visitor Entero

func (v *ReplVisitor) VisitValorEntero(ctx *parser.ValorEnteroContext) interface{} {
	num := ctx.ENTERO().GetText()
	val, err := strconv.Atoi(num)
	if err != nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Entero inválido: "+num)
		return value.NewNil()
	}
	return value.NewInt(val)
}




func (v *ReplVisitor) visitAtoi(args []parser.IExpresionContext, ctx antlr.ParserRuleContext) value.IVOR {
	if len(args) != 1 {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Atoi() espera exactamente 1 argumento")
		return value.NewNil()
	}

	val := v.Visit(args[0])
	ivor, ok := val.(value.IVOR)
	if !ok || ivor.Type() != value.IVOR_STRING {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Atoi() solo acepta strings como argumento")
		return value.NewNil()
	}

	str, ok := ivor.Value().(string)
	if !ok {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Valor no es string válido")
		return value.NewNil()
	}

	num, err := strconv.Atoi(str)
	if err != nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Atoi() no pudo convertir '"+str+"' a entero")
		return value.NewNil()
	}

	return value.NewInt(num)
}



/// visitor decimal

func (v *ReplVisitor) VisitValorDecimal(ctx *parser.ValorDecimalContext) interface{} {
	text := ctx.DECIMAL().GetText()
	floatVal, err := strconv.ParseFloat(text, 64)
	if err != nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Error al parsear decimal")
		return value.NewNil()
	}
	return value.NewFloat(floatVal)
}

/// visitor caracter
func (v *ReplVisitor) VisitValorCaracter(ctx *parser.ValorCaracterContext) interface{} {
	text := ctx.CARACTER().GetText() // ejemplo: `'a'`
	if len(text) == 3 {
		return value.NewChar(rune(text[1])) // extrae el carácter entre comillas
	}
	v.ErrorTable.NewSemanticError(ctx.GetStart(), "Carácter inválido")
	return value.NewNil()
}

/// visitor Booleano

func (v *ReplVisitor) VisitValorBooleano(ctx *parser.ValorBooleanoContext) interface{} {
    text := ctx.BOOLEANO().GetText()
    if text == "true" {
        return value.NewBool(true)
    } else if text == "false" {
        return value.NewBool(false)
    }
    v.ErrorTable.NewSemanticError(ctx.GetStart(), "Valor booleano inválido")
    return value.NewNil()
}

/// visitor Cadena

func (v *ReplVisitor) VisitValorCadena(ctx *parser.ValorCadenaContext) interface{} {
	text := ctx.CADENA().GetText() // Incluye comillas
	return value.NewString(text)   // ← sin cortar las comillas
}




// VisitParentesisexpre evalúa expresiones dentro de paréntesis
func (v *ReplVisitor) VisitParentesisexpre(ctx *parser.ParentesisexpreContext) interface{} {
	val := v.Visit(ctx.Expresion())
	if val == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Expresión dentro de paréntesis inválida")
		return value.NewNil()
	}
	if ivorVal, ok := val.(value.IVOR); ok {
		return ivorVal
	}
	v.ErrorTable.NewSemanticError(ctx.GetStart(), "Tipo inválido dentro de paréntesis")
	return value.NewNil()
}



// Expresiones

func (v *ReplVisitor) VisitUnario(ctx *parser.UnarioContext) interface{} {
	result := v.Visit(ctx.Expresion())
	expr, ok := result.(value.IVOR)
	if !ok || expr == nil || expr.Type() == value.IVOR_NIL {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "No se puede aplicar operador unario sobre un valor nil")
		return value.NewNil()
	}

	op := ctx.GetOp().GetText()

	if op == "-" {
		switch expr.Type() {
		case value.IVOR_INT:
			return value.NewInt(-expr.Value().(int))
		case value.IVOR_FLOAT:
			return value.NewFloat(-expr.Value().(float64))
		default:
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "Negación no válida para este tipo")
			return value.NewNil()
		}
	} else if op == "!" {
		if expr.Type() == value.IVOR_BOOL {
			return value.NewBool(!expr.Value().(bool))
		}
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "El operador ! solo es válido para valores booleanos")
		return value.NewNil()
	}

	v.ErrorTable.NewSemanticError(ctx.GetStart(), "Operador unario no reconocido")
	return value.NewNil()
}

func (v *ReplVisitor) VisitMultdivmod(ctx *parser.MultdivmodContext) interface{} {
	leftResult := v.Visit(ctx.Expresion(0))
	rightResult := v.Visit(ctx.Expresion(1))

	left, ok1 := leftResult.(value.IVOR)
	right, ok2 := rightResult.(value.IVOR)

	if !ok1 || left == nil || left.Type() == value.IVOR_NIL ||
	   !ok2 || right == nil || right.Type() == value.IVOR_NIL {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "No se puede operar con valores nil")
		return value.NewNil()
	}

	op := ctx.GetOp().GetText()

	switch op {
	case "*":
		switch left.Type() {
		case value.IVOR_INT:
			switch right.Type() {
			case value.IVOR_INT:
				return value.NewInt(left.Value().(int) * right.Value().(int))
			case value.IVOR_FLOAT:
				return value.NewFloat(float64(left.Value().(int)) * right.Value().(float64))
			}
		case value.IVOR_FLOAT:
			switch right.Type() {
			case value.IVOR_FLOAT:
				return value.NewFloat(left.Value().(float64) * right.Value().(float64))
			case value.IVOR_INT:
				return value.NewFloat(left.Value().(float64) * float64(right.Value().(int)))
			}
		}
	case "/":
		if right.Type() == value.IVOR_INT && right.Value().(int) == 0 {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "División por cero")
			return value.NewNil()
		}
		if right.Type() == value.IVOR_FLOAT && right.Value().(float64) == 0.0 {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "División por cero")
			return value.NewNil()
		}
		switch left.Type() {
		case value.IVOR_INT:
			switch right.Type() {
			case value.IVOR_INT:
				return value.NewFloat(float64(left.Value().(int)) / float64(right.Value().(int)))
			case value.IVOR_FLOAT:
				return value.NewFloat(float64(left.Value().(int)) / right.Value().(float64))
			}
		case value.IVOR_FLOAT:
			switch right.Type() {
			case value.IVOR_FLOAT:
				return value.NewFloat(left.Value().(float64) / right.Value().(float64))
			case value.IVOR_INT:
				return value.NewFloat(left.Value().(float64) / float64(right.Value().(int)))
			}
		}
	case "%":
		if left.Type() == value.IVOR_INT && right.Type() == value.IVOR_INT {
			if right.Value().(int) == 0 {
				v.ErrorTable.NewSemanticError(ctx.GetStart(), "Módulo por cero")
				return value.NewNil()
			}
			return value.NewInt(left.Value().(int) % right.Value().(int))
		}
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Módulo solo válido para enteros")
		return value.NewNil()
	default:
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Operador no válido, solo se permiten * / %")
		return value.NewNil()
	}

	v.ErrorTable.NewSemanticError(ctx.GetStart(), "Operación no válida entre tipos incompatibles")
	return value.NewNil()
}


func (v *ReplVisitor) VisitSumres(ctx *parser.SumresContext) interface{} {
	// Evaluar las expresiones izquierda y derecha
	leftResult := v.Visit(ctx.Expresion(0))
	rightResult := v.Visit(ctx.Expresion(1))

	// Verificar que los resultados no sean nil
	left, ok1 := leftResult.(value.IVOR)
	right, ok2 := rightResult.(value.IVOR)

	if !ok1 || left == nil || left.Type() == value.IVOR_NIL {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Operación inválida: la expresión izquierda es nil")
		return value.NewNil()
	}
	if !ok2 || right == nil || right.Type() == value.IVOR_NIL {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Operación inválida: la expresión derecha es nil")
		return value.NewNil()
	}

	op := ctx.GetOp().GetText()

	switch op {
	case "+":
		switch left.Type() {
		case value.IVOR_INT:
			switch right.Type() {
			case value.IVOR_INT:
				return value.NewInt(left.Value().(int) + right.Value().(int))
			case value.IVOR_FLOAT:
				return value.NewFloat(float64(left.Value().(int)) + right.Value().(float64))
			}
		case value.IVOR_FLOAT:
			switch right.Type() {
			case value.IVOR_FLOAT:
				return value.NewFloat(left.Value().(float64) + right.Value().(float64))
			case value.IVOR_INT:
				return value.NewFloat(left.Value().(float64) + float64(right.Value().(int)))
			}
		case value.IVOR_STRING:
			if right.Type() == value.IVOR_STRING {
				return value.NewString(left.Value().(string) + right.Value().(string))
			}
		}
		if left.Type() == value.IVOR_INT && right.Type() == value.IVOR_FLOAT {
			return value.NewFloat(float64(left.Value().(int)) + right.Value().(float64))
		}
	case "-":
		switch left.Type() {
		case value.IVOR_INT:
			switch right.Type() {
			case value.IVOR_INT:
				return value.NewInt(left.Value().(int) - right.Value().(int))
			case value.IVOR_FLOAT:
				return value.NewFloat(float64(left.Value().(int)) - right.Value().(float64))
			}
		case value.IVOR_FLOAT:
			switch right.Type() {
			case value.IVOR_FLOAT:
				return value.NewFloat(left.Value().(float64) - right.Value().(float64))
			case value.IVOR_INT:
				return value.NewFloat(left.Value().(float64) - float64(right.Value().(int)))
			}
		}
		if left.Type() == value.IVOR_INT && right.Type() == value.IVOR_FLOAT {
			return value.NewFloat(float64(left.Value().(int)) - right.Value().(float64))
		}
	default:
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Operador no válido, solo se permiten + o -")
		return value.NewNil()
	}

	v.ErrorTable.NewSemanticError(ctx.GetStart(), "Operación no válida entre tipos incompatibles")
	return value.NewNil()
}

func (v *ReplVisitor) VisitId(ctx *parser.IdContext) interface{} {
	id := ctx.ID().GetText()
	variable := v.CurrentScope.GetVariable(id)
	if variable == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), fmt.Sprintf("Variable '%s' no declarada", id))
		return value.NewNil()
	}

	// NO copiar si es struct
	if structInst, ok := variable.Value.(*value.StructInstance); ok {
		return structInst
	}

	// NO copiar si es slice
	if slice, ok := variable.Value.(*value.SliceValue); ok {
		return slice
	}

	// Solo copia para primitivos
	return variable.Value.Copy()
}


func (v *ReplVisitor) VisitExpresionStatement(ctx *parser.ExpresionStatementContext) interface{} {
	fmt.Println("➡️ VisitExpresionStatement")
	return v.Visit(ctx.Expresion())
}



func (v *ReplVisitor) VisitAsignacionCompuestaSuma(ctx *parser.AsignacionCompuestaSumaContext) interface{} {
	id := ctx.ID().GetText()
	variable := v.CurrentScope.GetVariable(id)

	if variable == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), fmt.Sprintf("Variable '%s' no declarada", id))
		return value.NewNil()
	}

	right := v.Visit(ctx.Expresion()).(value.IVOR)
	if right.Type() == value.IVOR_NIL {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "No se puede usar nil en +=.")
		return value.NewNil()
	}

	switch variable.Type {
	case value.IVOR_INT:
		if right.Type() == value.IVOR_INT {
			val := variable.Value.Value().(int) + right.Value().(int)
			variable.Value = value.NewInt(val)
			return variable.Value
		}
	case value.IVOR_FLOAT:
		switch right.Type() {
		case value.IVOR_INT:
			val := variable.Value.Value().(float64) + float64(right.Value().(int))
			variable.Value = value.NewFloat(val)
			return variable.Value
		case value.IVOR_FLOAT:
			val := variable.Value.Value().(float64) + right.Value().(float64)
			variable.Value = value.NewFloat(val)
			return variable.Value
		}
	case value.IVOR_STRING:
		if right.Type() == value.IVOR_STRING {
			val := variable.Value.Value().(string) + right.Value().(string)
			variable.Value = value.NewString(val)
			return variable.Value
		}
	}

	v.ErrorTable.NewSemanticError(ctx.GetStart(), "Tipos incompatibles en '+='")
	return value.NewNil()
}



func (v *ReplVisitor) VisitAsignacionCompuestaResta(ctx *parser.AsignacionCompuestaRestaContext) interface{} {
	id := ctx.ID().GetText()
	variable := v.CurrentScope.GetVariable(id)

	if variable == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), fmt.Sprintf("Variable '%s' no declarada", id))
		return value.NewNil()
	}

	rightResult := v.Visit(ctx.Expresion())
	right, ok := rightResult.(value.IVOR)
	if !ok || right == nil || right.Type() == value.IVOR_NIL {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Expresión inválida en la resta compuesta")
		return value.NewNil()
	}

	// Validación de tipos
	switch variable.Type {
	case value.IVOR_INT:
		switch right.Type() {
		case value.IVOR_INT:
			val := variable.Value.Value().(int) - right.Value().(int)
			variable.Value = value.NewInt(val)
			return variable.Value
		default:
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "No se puede restar tipos diferentes a int desde una variable int")
			return value.NewNil()
		}

	case value.IVOR_FLOAT:
		switch right.Type() {
		case value.IVOR_INT:
			val := variable.Value.Value().(float64) - float64(right.Value().(int))
			variable.Value = value.NewFloat(val)
			return variable.Value
		case value.IVOR_FLOAT:
			val := variable.Value.Value().(float64) - right.Value().(float64)
			variable.Value = value.NewFloat(val)
			return variable.Value
		default:
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "No se puede restar tipos no numéricos desde una variable float64")
			return value.NewNil()
		}

	default:
		v.ErrorTable.NewSemanticError(ctx.GetStart(), fmt.Sprintf("La variable '%s' no es de tipo numérico", id))
		return value.NewNil()
	}
}


// igualdad y desigualdad
func (v *ReplVisitor) VisitIgualdad(ctx *parser.IgualdadContext) interface{} {
	leftResult := v.Visit(ctx.Expresion(0))
	rightResult := v.Visit(ctx.Expresion(1))

	left, ok1 := leftResult.(value.IVOR)
	right, ok2 := rightResult.(value.IVOR)

	if !ok1 || !ok2 || left == nil || right == nil || left.Type() == value.IVOR_NIL || right.Type() == value.IVOR_NIL {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Comparación inválida: uno o ambos operandos son nil")
		return value.NewNil()
	}

	op := ctx.GetOp().GetText()

	// Comparación entre tipos válidos
	switch left.Type() {
	case value.IVOR_INT:
		switch right.Type() {
		case value.IVOR_INT:
			return value.NewBool(compararNumeros(op, float64(left.Value().(int)), float64(right.Value().(int))))
		case value.IVOR_FLOAT:
			return value.NewBool(compararNumeros(op, float64(left.Value().(int)), right.Value().(float64)))
		}
	case value.IVOR_FLOAT:
		switch right.Type() {
		case value.IVOR_FLOAT:
			return value.NewBool(compararNumeros(op, left.Value().(float64), right.Value().(float64)))
		case value.IVOR_INT:
			return value.NewBool(compararNumeros(op, left.Value().(float64), float64(right.Value().(int))))
		}
	case value.IVOR_BOOL:
		if right.Type() == value.IVOR_BOOL {
			return value.NewBool(compararBool(op, left.Value().(bool), right.Value().(bool)))
		}
	case value.IVOR_STRING:
		switch right.Type() {
		case value.IVOR_STRING:
			return value.NewBool(compararTexto(op, left.Value().(string), right.Value().(string)))
		case value.IVOR_CHARACTER:
			return value.NewBool(compararTexto(op, left.Value().(string), string(right.Value().(rune))))
		}
	case value.IVOR_CHARACTER:
		switch right.Type() {
		case value.IVOR_CHARACTER:
			return value.NewBool(compararNumeros(op, float64(left.Value().(rune)), float64(right.Value().(rune))))
		case value.IVOR_STRING:
			return value.NewBool(compararTexto(op, string(left.Value().(rune)), right.Value().(string)))
		}
	}

	v.ErrorTable.NewSemanticError(ctx.GetStart(), fmt.Sprintf("Comparación no válida entre tipos '%s' y '%s'", left.Type(), right.Type()))
	return value.NewNil()
}




// operaciones relacionales
func (v *ReplVisitor) VisitRelacionales(ctx *parser.RelacionalesContext) interface{} {
	leftResult := v.Visit(ctx.Expresion(0))
	rightResult := v.Visit(ctx.Expresion(1))

	left, ok1 := leftResult.(value.IVOR)
	right, ok2 := rightResult.(value.IVOR)

	if !ok1 || !ok2 || left == nil || right == nil || left.Type() == value.IVOR_NIL || right.Type() == value.IVOR_NIL {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "No se puede comparar con valores nil")
		return value.NewNil()
	}

	op := ctx.GetOp().GetText()

	switch left.Type() {
	case value.IVOR_INT:
		switch right.Type() {
		case value.IVOR_INT:
			return value.NewBool(relacion(op, float64(left.Value().(int)), float64(right.Value().(int))))
		case value.IVOR_FLOAT:
			return value.NewBool(relacion(op, float64(left.Value().(int)), right.Value().(float64)))
		}
	case value.IVOR_FLOAT:
		switch right.Type() {
		case value.IVOR_INT:
			return value.NewBool(relacion(op, left.Value().(float64), float64(right.Value().(int))))
		case value.IVOR_FLOAT:
			return value.NewBool(relacion(op, left.Value().(float64), right.Value().(float64)))
		}
	case value.IVOR_CHARACTER:
		if right.Type() == value.IVOR_CHARACTER {
			return value.NewBool(relacion(op, float64(left.Value().(rune)), float64(right.Value().(rune))))
		}
	}

	v.ErrorTable.NewSemanticError(ctx.GetStart(), fmt.Sprintf(
		"Comparación no válida entre tipos '%s' y '%s'", left.Type(), right.Type()))
	return value.NewNil()
}


// operadores logicos
func (v *ReplVisitor) VisitAndExpr(ctx *parser.AndExprContext) interface{} {
	left := v.Visit(ctx.Expresion(0)).(value.IVOR)
	right := v.Visit(ctx.Expresion(1)).(value.IVOR)

	if left.Type() != value.IVOR_BOOL || right.Type() != value.IVOR_BOOL {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "El operador '&&' solo es válido para booleanos")
		return value.NewNil()
	}

	return value.NewBool(left.Value().(bool) && right.Value().(bool))
}

func (v *ReplVisitor) VisitOrExpr(ctx *parser.OrExprContext) interface{} {
	left := v.Visit(ctx.Expresion(0)).(value.IVOR)
	right := v.Visit(ctx.Expresion(1)).(value.IVOR)

	if left.Type() != value.IVOR_BOOL || right.Type() != value.IVOR_BOOL {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "El operador '||' solo es válido para booleanos")
		return value.NewNil()
	}

	return value.NewBool(left.Value().(bool) || right.Value().(bool))
}


// tipos estaticos
func (v *ReplVisitor) VisitAsignacionfor(ctx *parser.AsignacionforContext) interface{} {
	id := ctx.ID().GetText()
	variable := v.CurrentScope.GetVariable(id)

	if variable == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), fmt.Sprintf("Variable '%s' no declarada", id))
		return value.NewNil()
	}

	nuevoValor := v.Visit(ctx.Expresion()).(value.IVOR)

	if variable.Type != nuevoValor.Type() {
		v.ErrorTable.NewSemanticError(ctx.GetStart(),
			fmt.Sprintf("No se puede asignar un valor de tipo '%s' a la variable '%s' de tipo '%s'",
				nuevoValor.Type(), id, variable.Type))
		return value.NewNil()
	}

	variable.Value = nuevoValor
	return nil
}





// tipos de datos por default

func defaultValueForType(tipo string) value.IVOR {
	switch tipo {
	case value.IVOR_INT:
		return value.NewInt(0)
	case value.IVOR_FLOAT:
		return value.NewFloat(0.0)
	case value.IVOR_STRING:
		return value.NewString("")
	case value.IVOR_BOOL:
		return value.NewBool(false)
	case value.IVOR_CHARACTER:
		return value.NewChar('\x00') // valor nulo de rune

	default:
		return value.NewNil()
	}
}


func isReservedWord(id string) bool {
	reserved := map[string]bool{
		"if": true, "for": true, "while": true, "mut": true,
		"print": true, "nil": true, "true": true, "false": true,
	}
	return reserved[strings.ToLower(id)]
}

// Función para mapear nombres de tipo externos a constantes internas
func NormalizeTypeName(name string) string {
	switch name {
	case "int":
		return value.IVOR_INT
	case "float64":
		return value.IVOR_FLOAT
	case "string":
		return value.IVOR_STRING
	case "bool":
		return value.IVOR_BOOL
	case "char":
		return value.IVOR_CHARACTER
	case "[]int":
		return "slice_Int"
	case "[][]int":
		return "slice_slice_Int"
	default:
		return ""
	}
}


func validateNestedSliceType(val value.IVOR, expectedType string) bool {
	actual := val.Type()
	return actual == expectedType
}



// bloques de ambitos
func (v *ReplVisitor) VisitBloqueStatement(ctx *parser.BloqueStatementContext) interface{} {
	// Crear un nuevo ámbito local
	newScope := NewBaseScope("bloque", v.CurrentScope)
	oldScope := v.CurrentScope
	v.CurrentScope = newScope

	// Obtener el nodo real de tipo `bloque`
	bloque := ctx.Bloque().(*parser.BloqueContext)
	for _, decl := range bloque.AllDeclaraciones() {
		result := v.Visit(decl)

		// ✅ Propagar return si aparece
		if ret, ok := result.(ReturnSignal); ok {
			v.CurrentScope = oldScope
			return ret
		}

		// ✅ Propagar break o continue si aparecen
		if ctrl, ok := result.(string); ok {
			if ctrl == "BREAK" || ctrl == "CONTINUE" {
				v.CurrentScope = oldScope
				return ctrl
			}
		}
	}

	// Restaurar el ámbito anterior
	v.CurrentScope = oldScope
	return nil
}




func (v *ReplVisitor) VisitBloque(ctx *parser.BloqueContext) interface{} {
	for _, decl := range ctx.AllDeclaraciones() {
		result := v.Visit(decl)

		// 🚨 Propagar return
		if ret, ok := result.(ReturnSignal); ok {
			return ret
		}

		// 🚨 Propagar break o continue
		if ctrl, ok := result.(string); ok {
			if ctrl == "BREAK" || ctrl == "CONTINUE" {
				return ctrl
			}
		}
	}
	return nil
}



// igualdad y desigualdad

func compararNumeros(op string, a, b float64) bool {
	if op == "==" {
		return a == b
	}
	return a != b
}

func compararBool(op string, a, b bool) bool {
	if op == "==" {
		return a == b
	}
	return a != b
}

func compararTexto(op string, a, b string) bool {
	if op == "==" {
		return a == b
	}
	return a != b
}



// parser de operaciones relaciones
func relacion(op string, a, b float64) bool {
	switch op {
	case "<":
		return a < b
	case ">":
		return a > b
	case "<=":
		return a <= b
	case ">=":
		return a >= b
	default:
		return false
	}
}

////////////////////// 	VISITOR IF-ELSE ELIF  //////////////////////////////////////////////////////

func (v *ReplVisitor) VisitIfDcl(ctx *parser.IfDclContext) interface{} {
	cond := v.Visit(ctx.Expresion())
	boolVal, ok := cond.(value.IVOR)

	if !ok || boolVal.Type() != value.IVOR_BOOL {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "La condición del if debe ser booleana")
		return nil
	}

	// 🟢 IF principal
	if boolVal.Value().(bool) {
		bloque := ctx.Bloque().(*parser.BloqueContext)
		for _, decl := range bloque.AllDeclaraciones() {
			res := v.Visit(decl)
			if str, ok := res.(string); ok {
				if str == "BREAK" || str == "CONTINUE" {
					return str
				}
			}
		}
		return nil
	}

	// 🔁 ELSE IF
	for _, elseif := range ctx.AllElseifDcl() {
		condElse := v.Visit(elseif.Expresion())
		boolElse, ok := condElse.(value.IVOR)
		if ok && boolElse.Type() == value.IVOR_BOOL && boolElse.Value().(bool) {
			bloque := elseif.Bloque().(*parser.BloqueContext)
			for _, decl := range bloque.AllDeclaraciones() {
				res := v.Visit(decl)
				if str, ok := res.(string); ok {
					if str == "BREAK" || str == "CONTINUE" {
						return str
					}
				}
			}
			return nil
		}
	}

	// 🟠 ELSE final
	if ctx.ElseDcl() != nil {
		bloque := ctx.ElseDcl().Bloque().(*parser.BloqueContext)
		for _, decl := range bloque.AllDeclaraciones() {
			res := v.Visit(decl)
			if str, ok := res.(string); ok {
				if str == "BREAK" || str == "CONTINUE" {
					return str
				}
			}
		}
	}

	return nil
}



func (v *ReplVisitor) VisitControlStatement(ctx *parser.ControlStatementContext) interface{} {
	return v.Visit(ctx.Sentencias_control())
}

func (v *ReplVisitor) VisitIf_context(ctx *parser.If_contextContext) interface{} {
	return v.Visit(ctx.IfDcl())
}




////////////////////// 	VISITOR PARA SWITCH CASE //////////////////////////////////////////////////////

func (v *ReplVisitor) VisitSwitch_context(ctx *parser.Switch_contextContext) interface{} {
	switchExpr := v.Visit(ctx.SwitchDcl().Expresion())
	valSwitch, ok := switchExpr.(value.IVOR)

	if !ok || valSwitch == nil || valSwitch.Type() == value.IVOR_NIL {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Expresión inválida en 'switch'")
		return nil
	}

	// Evaluar cada case
	for _, caseBlock := range ctx.SwitchDcl().AllCaseBlock() {
		caseExpr := v.Visit(caseBlock.Expresion())
		valCase, ok := caseExpr.(value.IVOR)

		if !ok || valCase == nil || valCase.Type() == value.IVOR_NIL {
			v.ErrorTable.NewSemanticError(caseBlock.GetStart(), "Expresión inválida en 'case'")
			continue
		}

		if compararValores(valSwitch, valCase) {
			newScope := NewBaseScope("switch_case", v.CurrentScope)
			oldScope := v.CurrentScope
			v.CurrentScope = newScope

			for _, decl := range caseBlock.AllDeclaraciones() {
				res := v.Visit(decl)

				if str, ok := res.(string); ok && (str == "BREAK" || str == "CONTINUE") {
					v.CurrentScope = oldScope
					return str
				}
				if ret, ok := res.(ReturnSignal); ok {
					v.CurrentScope = oldScope
					return ret
				}
			}

			v.CurrentScope = oldScope
			return nil
		}
	}

	// Default block
	if ctx.SwitchDcl().DefaultBlock() != nil {
		newScope := NewBaseScope("switch_default", v.CurrentScope)
		oldScope := v.CurrentScope
		v.CurrentScope = newScope

		for _, decl := range ctx.SwitchDcl().DefaultBlock().AllDeclaraciones() {
			res := v.Visit(decl)
			if str, ok := res.(string); ok && (str == "BREAK" || str == "CONTINUE") {
				v.CurrentScope = oldScope
				return str
			}
			if ret, ok := res.(ReturnSignal); ok {
				v.CurrentScope = oldScope
				return ret
			}
		}

		v.CurrentScope = oldScope
	}

	return nil
}




func compararValores(a, b value.IVOR) bool {
	if a.Type() != b.Type() {
		return false
	}
	return fmt.Sprintf("%v", a.Value()) == fmt.Sprintf("%v", b.Value())
}



////////////////////// 	VISITOR CICLO FOR //////////////////////////////////////////////////////

func (v *ReplVisitor) VisitFor_context(ctx *parser.For_contextContext) interface{} {
	return v.Visit(ctx.ForDcl())
}


func (v *ReplVisitor) VisitForCondicional(ctx *parser.ForCondicionalContext) interface{} {
	for {
		result := v.Visit(ctx.Expresion())
		cond, ok := result.(value.IVOR)
		if !ok || cond == nil || cond.Type() != value.IVOR_BOOL {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "La condición del for debe ser booleana")
			break
		}

		if !cond.Value().(bool) {
			break
		}

		bloque := ctx.Bloque().(*parser.BloqueContext)
		newScope := NewBaseScope("for_cond", v.CurrentScope)
		oldScope := v.CurrentScope
		v.CurrentScope = newScope

		salir := false
		continuar := false

		for _, decl := range bloque.AllDeclaraciones() {
			res := v.Visit(decl)
			if str, ok := res.(string); ok {
				if str == "BREAK" {
					salir = true
					break
				}
				if str == "CONTINUE" {
					continuar = true
					break
				}
			}
		}

		v.CurrentScope = oldScope

		if salir {
			break
		}
		if continuar {
			continue
		}
	}
	return nil
}



func (v *ReplVisitor) VisitForClasico(ctx *parser.ForClasicoContext) interface{} {
	_ = v.Visit(ctx.Expresion(0)) // inicialización

	for {
		cond := v.Visit(ctx.Expresion(1)).(value.IVOR)
		if cond.Type() != value.IVOR_BOOL || !cond.Value().(bool) {
			break
		}

		bloque := ctx.Bloque().(*parser.BloqueContext)
		newScope := NewBaseScope("for_clasico", v.CurrentScope)
		oldScope := v.CurrentScope
		v.CurrentScope = newScope

		salir := false
		continuar := false

		for _, decl := range bloque.AllDeclaraciones() {
			result := v.Visit(decl)
			if str, ok := result.(string); ok {
				if str == "BREAK" {
					salir = true
					break
				}
				if str == "CONTINUE" {
					continuar = true
					break
				}
			}
		}

		v.CurrentScope = oldScope

		if salir {
			break
		}

		if continuar {
			v.Visit(ctx.Expresion(2)) // incremento
			continue
		}

		v.Visit(ctx.Expresion(2)) // incremento
	}

	return nil
}



func (v *ReplVisitor) VisitForForeach(ctx *parser.ForForeachContext) interface{} {
	indexName := ctx.ID(0).GetText()
	valueName := ctx.ID(1).GetText()
	sliceName := ctx.ID(2).GetText()

	variable := v.CurrentScope.GetVariable(sliceName)
	if variable == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), fmt.Sprintf("Variable '%s' no definida", sliceName))
		return nil
	}

	val := variable.Value

	switch t := val.(type) {
	case *value.StringValue:
		// Iteración sobre string (caracteres)
		texto := t.InternalValue
		for i, ch := range texto {
			v.runForeachBlock(ctx, indexName, valueName, value.NewInt(i), value.NewChar(ch))
		}
	case *value.SliceValue:
		// Iteración sobre slice
		for i, elem := range t.Elements {
			v.runForeachBlock(ctx, indexName, valueName, value.NewInt(i), elem)
		}
	default:
		v.ErrorTable.NewSemanticError(ctx.GetStart(), fmt.Sprintf("'%s' no es iterable (string o slice)", sliceName))
		return nil
	}

	return nil
}


func (v *ReplVisitor) runForeachBlock(ctx *parser.ForForeachContext, idxName, valName string, idxVal, val value.IVOR) {
	newScope := NewBaseScope("for_each", v.CurrentScope)
	v.CurrentScope = newScope

	v.CurrentScope.AddVariable(idxName, idxVal.Type(), idxVal, false, false, ctx.GetStart())
	v.CurrentScope.AddVariable(valName, val.Type(), val, false, false, ctx.GetStart())

	salir := false
	continuar := false

	for _, decl := range ctx.Bloque().AllDeclaraciones() {
		res := v.Visit(decl)
		if str, ok := res.(string); ok {
			if str == "BREAK" {
				salir = true
				break
			}
			if str == "CONTINUE" {
				continuar = true
				break
			}
		}
	}

	v.CurrentScope = v.CurrentScope.Parent

	if salir {
		panic("break_foreach") // se puede cambiar por retorno de control si prefieres
	}
	if continuar {
		return
	}
}




////////////////////// 	VISITOR INCREMENTO SUMA//////////////////////////////////////////////////////


func (v *ReplVisitor) VisitAsignacionSuma(ctx *parser.AsignacionSumaContext) interface{} {
	id := ctx.ID().GetText()
	variable := v.CurrentScope.GetVariable(id)
	if variable == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), fmt.Sprintf("Variable '%s' no declarada", id))
		return value.NewNil()
	}

	if variable.Type != value.IVOR_INT && variable.Type != value.IVOR_FLOAT {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "El operador '++' solo es válido para enteros o flotantes")
		return value.NewNil()
	}

	if variable.Type == value.IVOR_INT {
		val := variable.Value.Value().(int) + 1
		variable.Value = value.NewInt(val)
		return variable.Value
	} else {
		val := variable.Value.Value().(float64) + 1.0
		variable.Value = value.NewFloat(val)
		return variable.Value
	}
}

////////////////////// 	VISITOR INCREMENTO RESTO//////////////////////////////////////////////////////

func (v *ReplVisitor) VisitAsignacionResta(ctx *parser.AsignacionRestaContext) interface{} {
	id := ctx.ID().GetText()
	variable := v.CurrentScope.GetVariable(id)
	if variable == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), fmt.Sprintf("Variable '%s' no declarada", id))
		return value.NewNil()
	}

	if variable.Type != value.IVOR_INT && variable.Type != value.IVOR_FLOAT {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "El operador '++' solo es válido para enteros o flotantes")
		return value.NewNil()
	}

	if variable.Type == value.IVOR_INT {
		val := variable.Value.Value().(int) - 1
		variable.Value = value.NewInt(val)
		return variable.Value
	} else {
		val := variable.Value.Value().(float64) - 1.0
		variable.Value = value.NewFloat(val)
		return variable.Value
	}
}


////////////////////// 	VISITOR BREAK ///////////////////////////////////////////////////////

func (v *ReplVisitor) VisitBreakStatement(ctx *parser.BreakStatementContext) interface{} {
	// Verificar si el break está dentro de un for o switch
	scope := v.CurrentScope
	permitido := false

	for scope != nil {
		if strings.HasPrefix(scope.Name, "for_") || strings.HasPrefix(scope.Name, "switch_") {
			permitido = true
			break
		}
		scope = scope.Parent
	}

	if !permitido {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "La sentencia 'break' solo puede usarse dentro de un for o switch")
		return value.NewNil()
	}

	// Devolvemos un string especial como señal de control
	return "BREAK"
}


////////////////////// visitor continue ///////////////////////////////////////////////////////

func (v *ReplVisitor) VisitContinueStatement(ctx *parser.ContinueStatementContext) interface{} {
	scope := v.CurrentScope
	for scope != nil {
		if strings.HasPrefix(scope.Name, "for_") {
			return "CONTINUE"
		}
		scope = scope.Parent
	}
	v.ErrorTable.NewSemanticError(ctx.GetStart(), "La sentencia 'continue' solo puede usarse dentro de un bucle")
	return value.NewNil()
}




//////////////////////// visitor return ///////////////////////////////////////////////////

func (v *ReplVisitor) VisitReturnStatement(ctx *parser.ReturnStatementContext) interface{} {
	if ctx.Expresion() != nil {
		val := v.Visit(ctx.Expresion()).(value.IVOR)
		return ReturnSignal{Value: val}
	}
	return ReturnSignal{Value: value.NewNil()}
}


type ReturnSignal struct {
	Value value.IVOR
}

// visitor Slices de slices
func (v *ReplVisitor) VisitSliceCreacion(ctx *parser.SliceCreacionContext) interface{} {
	// Validar tipo
	if ctx.Tipo() == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Debe especificar el tipo del slice")
		return value.NewNil()
	}

	tipoTexto := NormalizeTypeName(ctx.Tipo().GetText())
	if tipoTexto == "" || (!value.IsValidType(tipoTexto) && value.GlobalStructs[tipoTexto] == nil) {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Tipo inválido para slice: "+tipoTexto)
		return value.NewNil()
	}

	var elementos []value.IVOR
	if ctx.Parametros() != nil {
		for _, expr := range ctx.Parametros().AllExpresion() {
			raw := v.Visit(expr)
			if raw == nil {
				v.ErrorTable.NewSemanticError(ctx.GetStart(), "Elemento inválido en slice (nil)")
				return value.NewNil()
			}

			val, ok := raw.(value.IVOR)
			if !ok {
				v.ErrorTable.NewSemanticError(ctx.GetStart(), "Elemento inválido en slice")
				return value.NewNil()
			}

			// Validar tipo del elemento
			if val.Type() != tipoTexto {
				if sliceVal, ok := val.(*value.SliceValue); ok && sliceVal.Type() == tipoTexto {
					elementos = append(elementos, sliceVal)
				} else {
					v.ErrorTable.NewSemanticError(ctx.GetStart(),
						fmt.Sprintf("Tipo incompatible. Se esperaba '%s', se recibió '%s'", tipoTexto, val.Type()))
					return value.NewNil()
				}
			} else {
				elementos = append(elementos, val.Copy())
			}
		}
	}

	fmt.Printf("📦 Slice creado (tipo: %s) con %d elementos\n", tipoTexto, len(elementos))

	return &value.SliceValue{
		ElementType: tipoTexto,
		Elements:    elementos,
	}
}





func (v *ReplVisitor) VisitSliceDobleAcceso(ctx *parser.SliceDobleAccesoContext) interface{} {
	id := ctx.ID().GetText()
	variable := v.CurrentScope.GetVariable(id)
	if variable == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), fmt.Sprintf("Variable '%s' no declarada", id))
		return value.NewNil()
	}

	outer, ok := variable.Value.(*value.SliceValue)
	if !ok {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), fmt.Sprintf("'%s' no es un slice", id))
		return value.NewNil()
	}

	i := v.Visit(ctx.Expresion(0)).(value.IVOR)
	j := v.Visit(ctx.Expresion(1)).(value.IVOR)

	if i.Type() != value.IVOR_INT || j.Type() != value.IVOR_INT {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Los índices deben ser enteros")
		return value.NewNil()
	}

	rowIdx := i.Value().(int)
	colIdx := j.Value().(int)

	if rowIdx < 0 || rowIdx >= len(outer.Elements) {
		v.ErrorTable.NewRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "Índice de fila fuera de rango")
		return value.NewNil()
	}

	rowVal := outer.Elements[rowIdx]
	row, ok := rowVal.(*value.SliceValue)
	if !ok {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Elemento no es un slice (fila inválida)")
		return value.NewNil()
	}

	if colIdx < 0 || colIdx >= len(row.Elements) {
		v.ErrorTable.NewRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "Índice de columna fuera de rango")
		return value.NewNil()
	}
	fmt.Printf("🔍 Lectura mtx[%d][%d]: %v\n", rowIdx, colIdx, row.Elements[colIdx].Value())

	return row.Elements[colIdx].Copy()
}



func (v *ReplVisitor) VisitSliceDobleAsignacion(ctx *parser.SliceDobleAsignacionContext) interface{} {
	id := ctx.ID().GetText()
	variable := v.CurrentScope.GetVariable(id)
	if variable == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), fmt.Sprintf("Variable '%s' no declarada", id))
		return value.NewNil()
	}

	outer, ok := variable.Value.(*value.SliceValue)
	if !ok {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), fmt.Sprintf("'%s' no es un slice", id))
		return value.NewNil()
	}

	i := v.Visit(ctx.Expresion(0)).(value.IVOR)
	j := v.Visit(ctx.Expresion(1)).(value.IVOR)
	newVal := v.Visit(ctx.Expresion(2)).(value.IVOR)

	if i.Type() != value.IVOR_INT || j.Type() != value.IVOR_INT {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Los índices deben ser enteros")
		return value.NewNil()
	}

	rowIdx := i.Value().(int)
	colIdx := j.Value().(int)

	if rowIdx < 0 || rowIdx >= len(outer.Elements) {
		v.ErrorTable.NewRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "Índice de fila fuera de rango")
		return value.NewNil()
	}

	raw := outer.Elements[rowIdx]
	row, ok := raw.(*value.SliceValue)
	if !ok {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Elemento no es un slice (fila inválida)")
		return value.NewNil()
	}

	if colIdx < 0 || colIdx >= len(row.Elements) {
		v.ErrorTable.NewRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "Índice de columna fuera de rango")
		return value.NewNil()
	}

	if newVal.Type() != row.ElementType {
		v.ErrorTable.NewSemanticError(ctx.GetStart(),
			fmt.Sprintf("Tipo incompatible. La fila contiene '%s' y se intentó asignar '%s'",
				row.ElementType, newVal.Type()))
		return value.NewNil()
	}

	// ✅ Asignar directamente sobre la referencia del slice interno
	row.Elements[colIdx] = newVal.Copy()

	fmt.Printf("🟢 Asignación mtx[%d][%d] = %v\n", rowIdx, colIdx, newVal.Value())

	return newVal

}




func (v *ReplVisitor) VisitLiteralSliceExpr(ctx *parser.LiteralSliceExprContext) interface{} {
	var elementos []value.IVOR

	if ctx.Parametros() != nil {
		for _, expr := range ctx.Parametros().AllExpresion() {
			val := v.Visit(expr)

			if slice, ok := val.(*value.SliceValue); ok {
				elementos = append(elementos, slice) // ✅ usar referencia directa
				fmt.Printf("INSERTANDO sub-slice en matriz: %p\n", slice)

			} else if ivor, ok := val.(value.IVOR); ok {
				elementos = append(elementos, ivor.Copy()) // ✅ copiar primitivos
			} else {
				v.ErrorTable.NewSemanticError(ctx.GetStart(), "Elemento inválido en literal de slice")
				return value.NewNil()
			}
		}
	}

	if len(elementos) == 0 {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "No se puede deducir el tipo de un slice vacío")
		return value.NewNil()
	}

	first := elementos[0]

	// Caso slice de slices (matriz)
	if nestedSlice, ok := first.(*value.SliceValue); ok {
		for _, e := range elementos[1:] {
			if s, ok := e.(*value.SliceValue); !ok || s.ElementType != nestedSlice.ElementType {
				v.ErrorTable.NewSemanticError(ctx.GetStart(), "Todos los elementos deben ser slices del mismo tipo")
				return value.NewNil()
			}
		}

		fmt.Printf("📦 Slice externo (matriz) creado con %d filas, tipo interno: %s\n",
			len(elementos), nestedSlice.ElementType)

		return &value.SliceValue{
			ElementType: nestedSlice.ElementType,
			Elements:    elementos,
		}
	}

	// Caso slice de valores primitivos
	elemType := first.Type()
	for _, e := range elementos[1:] {
		if e.Type() != elemType {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "Todos los elementos del slice deben ser del mismo tipo")
			return value.NewNil()
		}
	}

	fmt.Printf("📦 Slice de primitivos creado con %d elementos, tipo: %s\n",
		len(elementos), elemType)

	return &value.SliceValue{
		ElementType: elemType,
		Elements:    elementos,
	}
}




// visitor funcion indexof de slices

func (v *ReplVisitor) VisitIndexOfExpr(ctx *parser.IndexOfExprContext) interface{} {
	sliceVal := v.Visit(ctx.Expresion(0))
	targetVal := v.Visit(ctx.Expresion(1))

	slice, ok := sliceVal.(*value.SliceValue)
	if !ok {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "El primer parámetro de 'indexOf' debe ser un slice")
		return value.NewInt(-1)
	}

	for i, elem := range slice.Elements {
		if elem.Type() == targetVal.(value.IVOR).Type() &&
			fmt.Sprintf("%v", elem.Value()) == fmt.Sprintf("%v", targetVal.(value.IVOR).Value()) {
			return value.NewInt(i)
		}
	}

	return value.NewInt(-1)
}


// visitor funcion join de slices
func (v *ReplVisitor) VisitJoinExpr(ctx *parser.JoinExprContext) interface{} {
	sliceVal := v.Visit(ctx.Expresion(0))
	sepVal := v.Visit(ctx.Expresion(1))

	// Validar separador
	sep, ok := sepVal.(value.IVOR)
	if !ok || sep.Type() != value.IVOR_STRING {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "El segundo parámetro de 'join' debe ser una cadena")
		return value.NewNil()
	}

	// Validar slice de strings
	slice, ok := sliceVal.(*value.SliceValue)
	if !ok || slice.ElementType != value.IVOR_STRING {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "El primer parámetro de 'join' debe ser un slice de cadenas ([]string)")
		return value.NewNil()
	}

	// Construir resultado
	var parts []string
	for _, elem := range slice.Elements {
		strElem, ok := elem.(*value.StringValue)
		if !ok {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "Todos los elementos del slice deben ser cadenas")
			return value.NewNil()
		}
		parts = append(parts, strElem.InternalValue)
	}

	joined := strings.Join(parts, sep.Value().(string))
	return value.NewString(joined)
}


// visitor funcion len de slices
func (v *ReplVisitor) VisitLenExpr(ctx *parser.LenExprContext) interface{} {
	val := v.Visit(ctx.Expresion())

	// Slices válidos
	if slice, ok := val.(*value.SliceValue); ok {
		return value.NewInt(len(slice.Elements))
	}

	// También opcionalmente permitir len("texto")
	if str, ok := val.(*value.StringValue); ok {
		return value.NewInt(len(str.InternalValue))
	}

	v.ErrorTable.NewSemanticError(ctx.GetStart(), "La función 'len' solo acepta slices o cadenas")
	return value.NewNil()
}



// visitor append funcion slice
func (v *ReplVisitor) VisitAppendExpr(ctx *parser.AppendExprContext) interface{} {
	sliceVal := v.Visit(ctx.Expresion(0))
	elemVal := v.Visit(ctx.Expresion(1))

	slice, ok := sliceVal.(*value.SliceValue)
	if !ok {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "El primer parámetro de 'append' debe ser un slice")
		return value.NewNil()
	}

	elem, ok := elemVal.(value.IVOR)
	if !ok {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Elemento inválido para 'append'")
		return value.NewNil()
	}

	if elem.Type() != slice.ElementType {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), fmt.Sprintf(
			"Tipo incompatible. El slice contiene '%s' y se intentó añadir '%s'",
			slice.ElementType, elem.Type()))
		return value.NewNil()
	}

	// ✅ No copiar — mantener referencia
	newElements := append(slice.Elements, elem)

	newSlice := &value.SliceValue{
		ElementType: slice.ElementType,
		Elements:    newElements,
	}

	fmt.Printf("📌 append → tipo: %s, total: %d elementos\n", newSlice.ElementType, len(newSlice.Elements))
	return newSlice
}




// visitiro acceso y asignacion slices
func (v *ReplVisitor) VisitSliceAcceso(ctx *parser.SliceAccesoContext) interface{} {
	id := ctx.ID().GetText()
	variable := v.CurrentScope.GetVariable(id)

	if variable == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), fmt.Sprintf("Variable '%s' no declarada", id))
		return value.NewNil()
	}

	indexVal := v.Visit(ctx.Expresion()).(value.IVOR)
	if indexVal.Type() != value.IVOR_INT {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "El índice debe ser un entero")
		return value.NewNil()
	}

	sliceVal, ok := variable.Value.(*value.SliceValue)
	if !ok {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), fmt.Sprintf("'%s' no es un slice", id))
		return value.NewNil()
	}

	index := indexVal.Value().(int)
	if index < 0 || index >= len(sliceVal.Elements) {
		v.ErrorTable.NewRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "Índice fuera de rango")
		return value.NewNil()
	}

	return sliceVal.Elements[index].Copy()
}


func (v *ReplVisitor) VisitSliceAsignacion(ctx *parser.SliceAsignacionContext) interface{} {
	id := ctx.ID().GetText()
	variable := v.CurrentScope.GetVariable(id)

	if variable == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), fmt.Sprintf("Variable '%s' no declarada", id))
		return value.NewNil()
	}

	indexVal := v.Visit(ctx.Expresion(0)).(value.IVOR)
	valorVal := v.Visit(ctx.Expresion(1)).(value.IVOR)

	if indexVal.Type() != value.IVOR_INT {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "El índice debe ser un entero")
		return value.NewNil()
	}

	sliceVal, ok := variable.Value.(*value.SliceValue)
	if !ok {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), fmt.Sprintf("'%s' no es un slice", id))
		return value.NewNil()
	}

	if valorVal.Type() != sliceVal.ElementType {
		v.ErrorTable.NewSemanticError(ctx.GetStart(),
			fmt.Sprintf("Tipo incompatible. El slice contiene '%s' y se intentó asignar '%s'",
				sliceVal.ElementType, valorVal.Type()))
		return value.NewNil()
	}

	index := indexVal.Value().(int)
	if index < 0 || index >= len(sliceVal.Elements) {
		v.ErrorTable.NewRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "Índice fuera de rango")
		return value.NewNil()
	}

	sliceVal.Elements[index] = valorVal.Copy()
	return nil
}
//////////////////////////////////////////////////////////////////////////////////////



func formatSliceRecursively(s *value.SliceValue) string {
	var parts []string
	for _, elem := range s.Elements {
		if inner, ok := elem.(*value.SliceValue); ok {
			parts = append(parts, formatSliceRecursively(inner)) // Recursión para slices anidados
		} else {
			parts = append(parts, fmt.Sprintf("%v", elem.Value()))
		}
	}
	return "[" + strings.Join(parts, ", ") + "]"
}



// structs 
func (v *ReplVisitor) VisitStructDeclaration(ctx *parser.StructDeclarationContext) interface{} {
	name := ctx.ID().GetText()
	fmt.Printf("🔍 Visitando declaración de struct: %s\n", name)

	if _, exists := value.GlobalStructs[name]; exists {
		fmt.Printf("❌ Error: Struct '%s' ya definido\n", name)
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Struct ya definido: "+name)
		return nil
	}
	fmt.Printf("✅ Nombre de struct '%s' válido (no existe previamente)\n", name)

	var fields []value.StructField
	atributosCtx := ctx.Atributos()
	if atributosCtx == nil {
		fmt.Printf("❌ Error: Struct '%s' no tiene atributos\n", name)
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Struct sin atributos")
		return nil
	}
	fmt.Printf("📋 Procesando atributos de struct '%s'\n", name)

	for _, campo := range atributosCtx.AllAtributo() {
		tipo := NormalizeTypeName(campo.Tipo().GetText())
		campoName := campo.ID().GetText()
		fmt.Printf("🔍 Procesando campo: %s (tipo: %s)\n", campoName, tipo)
		fmt.Printf("🔎 Validando tipo: IsValidType(%s)=%v, GlobalStructs[%s]=%v\n", tipo, value.IsValidType(tipo), tipo, value.GlobalStructs[tipo])

		if !value.IsValidType(tipo) && value.GlobalStructs[tipo] == nil {
			fmt.Printf("❌ Error: Tipo no válido o struct no definido para campo '%s': %s\n", campoName, tipo)
			v.ErrorTable.NewSemanticError(campo.GetStart(), "Tipo no válido o struct no definido: "+tipo)
			return nil
		}
		fmt.Printf("✅ Campo %s (tipo: %s) válido\n", campoName, tipo)
		fields = append(fields, value.StructField{
			Name: campoName,
			Type: tipo,
		})
	}

	value.GlobalStructs[name] = &value.StructDefinition{
		Name:   name,
		Fields: fields,
	}
	fmt.Printf("📦 Struct '%s' registrado con %d campos\n", name, len(fields))
	fmt.Printf("📚 Estado de GlobalStructs: %v\n", value.GlobalStructs)
	return nil
}




// visitor instancias de estructuras
func (v *ReplVisitor) VisitInstanciaStruct(ctx *parser.InstanciaStructContext) interface{} {
	name := ctx.ID().GetText()
	fmt.Printf("🔨 Creando instancia de struct: %s\n", name)
	fmt.Printf("📚 Estado de GlobalStructs: %v\n", value.GlobalStructs)

	def := value.GlobalStructs[name]
	if def == nil {
		fmt.Printf("❌ Error: Struct '%s' no definido\n", name)
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Struct no definido: "+name)
		return value.NewNil()
	}
	fmt.Printf("✅ Struct '%s' encontrado, con %d campos definidos\n", name, len(def.Fields))

	fields := make(map[string]value.IVOR)
	if ctx.InicializadorStruct() != nil {
		fmt.Printf("📋 Inicializando campos de la instancia\n")
		for _, campo := range ctx.InicializadorStruct().AllInicializadorCampo() {
			campoNombre := campo.ID().GetText()
			fmt.Printf("🔍 Procesando campo inicializado: %s\n", campoNombre)

			rawVal := v.Visit(campo.Expresion())
			if rawVal == nil {
				fmt.Printf("❌ Error: Valor nil en campo '%s'\n", campoNombre)
				v.ErrorTable.NewSemanticError(campo.GetStart(), "Expresión inválida al inicializar campo '"+campoNombre+"'")
				return value.NewNil()
			}

			val, ok := rawVal.(value.IVOR)
			if !ok {
				fmt.Printf("❌ Error: Valor no es IVOR en campo '%s'\n", campoNombre)
				v.ErrorTable.NewSemanticError(campo.GetStart(), "Expresión no retorna un valor válido en campo '"+campoNombre+"'")
				return value.NewNil()
			}

			found := false
			for _, defField := range def.Fields {
				if defField.Name == campoNombre {
					if defField.Type != val.Type() {
						fmt.Printf("❌ Error: Tipo incompatible para campo '%s'. Esperado: %s, Obtenido: %s\n", campoNombre, defField.Type, val.Type())
						v.ErrorTable.NewSemanticError(campo.GetStart(),
							fmt.Sprintf("Tipo incompatible para el campo '%s'", campoNombre))
						return value.NewNil()
					}
					found = true
					break
				}
			}
			if !found {
				fmt.Printf("❌ Error: Campo '%s' no existe en struct '%s'\n", campoNombre, name)
				v.ErrorTable.NewSemanticError(campo.GetStart(), fmt.Sprintf("Campo '%s' no existe en struct '%s'", campoNombre, name))
				return value.NewNil()
			}
			fields[campoNombre] = val.Copy()
			fmt.Printf("✅ Campo %s asignado correctamente\n", campoNombre)
		}
	}

	// Completar campos no inicializados
	for _, defField := range def.Fields {
		if _, exists := fields[defField.Name]; !exists {
			fmt.Printf("⚠️ Advertencia: Campo '%s' no inicializado, se asigna valor nil\n", defField.Name)
			fields[defField.Name] = value.NewNil()
		}
	}

	fmt.Printf("🎉 Instancia de struct '%s' creada con %d campos inicializados\n", name, len(fields))
	return value.NewStructInstance(def, fields)
}





// visitor acceso atributos
func (v *ReplVisitor) VisitAccesoAtributo(ctx *parser.AccesoAtributoContext) interface{} {
	id := ctx.ID(0).GetText()
	attr := ctx.ID(1).GetText()
	fmt.Printf("🔍 Accediendo a atributo: %s.%s\n", id, attr)

	variable := v.CurrentScope.GetVariable(id)
	if variable == nil {
		fmt.Printf("❌ Error: Variable '%s' no declarada\n", id)
		v.ErrorTable.NewSemanticError(ctx.GetStart(), fmt.Sprintf("Variable '%s' no declarada", id))
		return value.NewNil()
	}

	inst, ok := variable.Value.(*value.StructInstance)
	if !ok {
		fmt.Printf("❌ Error: '%s' no es una instancia de struct\n", id)
		v.ErrorTable.NewSemanticError(ctx.GetStart(), fmt.Sprintf("'%s' no es una instancia de struct", id))
		return value.NewNil()
	}

	val, ok := inst.Fields[attr]
	if !ok {
		fmt.Printf("❌ Error: Campo '%s' no existe en struct '%s'\n", attr, inst.Def.Name)
		v.ErrorTable.NewSemanticError(ctx.GetStart(), fmt.Sprintf("Campo '%s' no existe en struct '%s'", attr, inst.Def.Name))
		return value.NewNil()
	}

	fmt.Printf("✅ Atributo %s.%s accedido, valor: %s\n", id, attr, val.Type())

	// ❗ Evita copiar si el campo es slice o struct
	if _, isSlice := val.(*value.SliceValue); isSlice {
		return val
	}
	if _, isStruct := val.(*value.StructInstance); isStruct {
		return val
	}
	return val.Copy()
}



// visit asignacion atributos
func (v *ReplVisitor) VisitAsignacionAtributo(ctx *parser.AsignacionAtributoContext) interface{} {
	id := ctx.ID(0).GetText()
	attr := ctx.ID(1).GetText()
	fmt.Printf("🔧 Asignando valor a atributo: %s.%s\n", id, attr)

	newVal := v.Visit(ctx.Expresion()).(value.IVOR)
	fmt.Printf("📢 Nuevo valor para %s.%s: %v (tipo: %s)\n", id, attr, newVal, newVal.Type())

	variable := v.CurrentScope.GetVariable(id)
	if variable == nil {
		fmt.Printf("❌ Error: Variable '%s' no declarada\n", id)
		v.ErrorTable.NewSemanticError(ctx.GetStart(), fmt.Sprintf("Variable '%s' no declarada", id))
		return value.NewNil()
	}

	inst, ok := variable.Value.(*value.StructInstance)
	if !ok {
		fmt.Printf("❌ Error: '%s' no es una instancia de struct\n", id)
		v.ErrorTable.NewSemanticError(ctx.GetStart(), fmt.Sprintf("'%s' no es una instancia de struct", id))
		return value.NewNil()
	}

	def := inst.Def
	for _, f := range def.Fields {
		if f.Name == attr {
			if f.Type != newVal.Type() {
				fmt.Printf("❌ Error: Tipo incompatible para '%s'. Esperado: %s, Obtenido: %s\n", attr, f.Type, newVal.Type())
				v.ErrorTable.NewSemanticError(ctx.GetStart(), fmt.Sprintf("Tipo incompatible. Campo '%s' es '%s' y se intentó asignar '%s'", attr, f.Type, newVal.Type()))
				return value.NewNil()
			}

			// ✅ Asignación segura: copia si es primitivo
			if _, isSlice := newVal.(*value.SliceValue); isSlice {
				inst.Fields[attr] = newVal
			} else if _, isStruct := newVal.(*value.StructInstance); isStruct {
				inst.Fields[attr] = newVal
			} else {
				inst.Fields[attr] = newVal.Copy()
			}

			fmt.Printf("✅ Asignado: %s.%s = %v\n", id, attr, newVal)
			return nil
		}
	}

	fmt.Printf("❌ Error: Campo '%s' no definido en struct '%s'\n", attr, def.Name)
	v.ErrorTable.NewSemanticError(ctx.GetStart(), fmt.Sprintf("Campo '%s' no definido en struct '%s'", attr, def.Name))
	return value.NewNil()
}




// funcion con structs
//// no funciona struc con funciones
func (v *ReplVisitor) VisitFuncionStruct(ctx *parser.FuncionStructContext) interface{} {
	receiverName := ctx.ID(0).GetText()
	receiverType := ctx.ID(1).GetText()
	funcName := ctx.ID(2).GetText()

	// Validar existencia del struct receptor
	if value.GlobalStructs[receiverType] == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Struct no definido: "+receiverType)
		return nil
	}

	// Procesar parámetros
	var params []value.FunctionParam
	if ctx.ParametrosFormales() != nil {
		for _, param := range ctx.ParametrosFormales().AllParametroFormal() {
			p := param.(*parser.ParametroFormalContext)

			// Verificación de nil para TipoConReferencia y su hijo
			var tipo string
			var passByRef bool
			if p.TipoConReferencia() != nil {
				tipoCtx := p.TipoConReferencia().GetTipoTipo()
				if tipoCtx != nil {
					tipo = NormalizeTypeName(tipoCtx.GetText())
				} else {
					v.ErrorTable.NewSemanticError(p.GetStart(), "Tipo no especificado en parámetro")
					return nil
				}
				passByRef = p.TipoConReferencia().GetRef() != nil
			} else {
				v.ErrorTable.NewSemanticError(p.GetStart(), "Referencia de tipo inválida en parámetro")
				return nil
			}

			nombre := p.GetTipoRef().GetText()

			if !value.IsValidType(tipo) && value.GlobalStructs[tipo] == nil {
				v.ErrorTable.NewSemanticError(p.GetStart(), "Tipo no válido: "+tipo)
				return nil
			}

			params = append(params, value.FunctionParam{
				Name:            nombre,
				Type:            tipo,
				PassByReference: passByRef,
			})
		}
	}

	// Registrar la función del struct
	if value.GlobalStructFunctions[receiverType] == nil {
		value.GlobalStructFunctions[receiverType] = make(map[string]*value.StructFunction)
	}
	if _, exists := value.GlobalStructFunctions[receiverType][funcName]; exists {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Función '"+funcName+"' ya definida para struct "+receiverType)
		return nil
	}

	value.GlobalStructFunctions[receiverType][funcName] = &value.StructFunction{
		ReceiverName: receiverName,
		ReceiverType: receiverType,
		Params:       params,
		Body:         ctx.Bloque(),
	}

	return nil
}




func (v *ReplVisitor) VisitLlamadaMetodoStruct(ctx *parser.LlamadaMetodoStructContext) interface{} {
	instName := ctx.ID(0).GetText()
	methodName := ctx.ID(1).GetText()

	variable := v.CurrentScope.GetVariable(instName)
	if variable == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), fmt.Sprintf("Variable '%s' no declarada", instName))
		return value.NewNil()
	}

	inst, ok := variable.Value.(*value.StructInstance)
	if !ok {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), fmt.Sprintf("'%s' no es instancia de un struct", instName))
		return value.NewNil()
	}

	methods := value.GlobalStructFunctions[inst.Def.Name]
	if methods == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), fmt.Sprintf("El struct '%s' no tiene métodos definidos", inst.Def.Name))
		return value.NewNil()
	}

	method := methods[methodName]
	if method == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), fmt.Sprintf("Método '%s' no definido en struct '%s'", methodName, inst.Def.Name))
		return value.NewNil()
	}

	newScope := NewBaseScope("method_"+methodName, v.CurrentScope)

	// Registrar el receptor por referencia
	_, err := newScope.AddVariable(method.ReceiverName, method.ReceiverType, inst, false, false, ctx.GetStart())
	if err != "" {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), err)
		return value.NewNil()
	}

	// Evaluar y registrar parámetros
	argExprs := []parser.IExpresionContext{}
	if ctx.Parametros() != nil {
		argExprs = ctx.Parametros().AllExpresion()
	}
	if len(argExprs) != len(method.Params) {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), fmt.Sprintf("Número de argumentos incorrecto en llamada a '%s'", methodName))
		return value.NewNil()
	}

	for i, param := range method.Params {
		val := v.Visit(argExprs[i]).(value.IVOR)

		if val.Type() != param.Type {
			v.ErrorTable.NewSemanticError(ctx.GetStart(),
				fmt.Sprintf("Tipo incorrecto para parámetro '%s'. Esperado: %s, recibido: %s",
					param.Name, param.Type, val.Type()))
			return value.NewNil()
		}

		if param.PassByReference {
			// ✅ Pasaje por referencia: no copiar
			_, err = newScope.AddVariable(param.Name, param.Type, val, false, false, ctx.GetStart())
		} else {
			// ✅ Por valor: copiar el dato
			_, err = newScope.AddVariable(param.Name, param.Type, val.Copy(), false, true, ctx.GetStart())
		}

		if err != "" {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), err)
			return value.NewNil()
		}
	}

	// Ejecutar el cuerpo
	prevScope := v.CurrentScope
	v.CurrentScope = newScope
	result := v.Visit(method.Body)
	v.CurrentScope = prevScope

	// Si la función retorna algo, capturarlo
	if ret, ok := result.(ReturnSignal); ok {
		return ret.Value
	}

	return value.NewNil()
}



//// no funciona struc con funciones






// visitor de funciones 
func (v *ReplVisitor) VisitFuncion(ctx *parser.FuncionContext) interface{} {
	defer func() {
		if r := recover(); r != nil {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), fmt.Sprintf("Error interno al procesar función: %v", r))
		}
	}()

	name := ctx.ID().GetText()
	fmt.Println("🔧 Registrando función:", name)

	// Verificar si el nombre ya existe
	if _, exists := value.GlobalFunctions[name]; exists ||
		value.GlobalStructs[name] != nil ||
		v.CurrentScope.GetVariable(name) != nil {
		v.ErrorTable.NewSemanticError(ctx.ID().GetSymbol(), "El nombre '"+name+"' ya está en uso.")
		return nil
	}

	// === Validación y construcción de parámetros ===
	paramCtx := ctx.ParametrosFormales()
	paramMap := make(map[string]bool)
	var params []value.FunctionParam

	// Tipos válidos conocidos
	validTypes := map[string]bool{
		"int": true, "float": true, "float64": true, "string": true,
		"bool": true, "char": true, "void": true,
		// Agrega más si tienes structs definidos
	}

	if paramCtx != nil {
		for _, p := range paramCtx.AllParametroFormal() {
			param := p.(*parser.ParametroFormalContext)

			tipoRef := param.GetTipoRef()
			idNode := param.ID()
			if tipoRef == nil || idNode == nil {
				v.ErrorTable.NewSemanticError(param.GetStart(), "Parámetro inválido: tipo o nombre ausente.")
				return nil
			}

			tipo := tipoRef.GetText()
			id := idNode.GetText()

			if paramMap[id] {
				v.ErrorTable.NewSemanticError(param.GetStart(), "Parámetro duplicado: '"+id+"'")
				return nil
			}
			paramMap[id] = true

			if !validTypes[tipo] && value.GlobalStructs[tipo] == nil {
				v.ErrorTable.NewSemanticError(param.GetStart(), "Tipo no válido o no definido: '"+tipo+"'")
				return nil
			}

			params = append(params, value.FunctionParam{Name: id, Type: tipo})
		}
	}

	// === Tipo de retorno ===
	retType := ""
	if ctx.Tipo() != nil {
		retType = ctx.Tipo().GetText()
		if !validTypes[retType] && value.GlobalStructs[retType] == nil {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "Tipo de retorno no válido o no definido: '"+retType+"'")
			return nil
		}
	}

	// === Registro de función global ===
	value.GlobalFunctions[name] = &value.Function{
		Name:       name,
		ReturnType: retType,
		Params:     params,
		Body:       ctx.Bloque(),
	}

	// === Agregar a tabla de símbolos ===
	retTypeStr := "void"
	if ctx.Tipo() != nil {
		retTypeStr = ctx.Tipo().GetText()
	}

	v.Symbols = append(v.Symbols, SymbolInfo{
		ID:      name,
		Tipo:    "función " + strings.Join(append([]string{retTypeStr}, funcParamTypes(ctx.ParametrosFormales())...), ", "),
		Valor:   "",
		Ambito:  v.CurrentScope.Name,
		Linea:   ctx.GetStart().GetLine(),
		Columna: ctx.GetStart().GetColumn(),
	})

	return nil
}


func funcParamTypes(params parser.IParametrosFormalesContext) []string {
	if params == nil {
		return []string{}
	}
	var types []string
	for _, p := range params.AllParametroFormal() {
		param := p.(*parser.ParametroFormalContext)
		if param.GetTipoRef() != nil {
			types = append(types, param.GetTipoRef().GetText())
		} else {
			types = append(types, "undefined")
		}
	}
	return types
}

func (v *ReplVisitor) VisitLlamadaFuncion(ctx *parser.LlamadaFuncionContext) interface{} {
	defer func() {
		if r := recover(); r != nil {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), fmt.Sprintf("Error interno en llamada a función: %v", r))
		}
	}()

	name := ctx.ID().GetText()

	// === Funciones embebidas ===
	switch name {
	case "Atoi":
		args := []parser.IExpresionContext{}
		if ctx.Parametros() != nil {
			args = ctx.Parametros().AllExpresion()
		}
		return v.visitAtoi(args, ctx)

	case "parseFloat":
		args := []parser.IExpresionContext{}
		if ctx.Parametros() != nil {
			args = ctx.Parametros().AllExpresion()
		}
		return v.visitParseFloat(args, ctx)

	case "typeOf":
		args := []parser.IExpresionContext{}
		if ctx.Parametros() != nil {
			args = ctx.Parametros().AllExpresion()
		}
		return v.visitTypeOf(args, ctx)
	}

	// === Buscar función definida ===
	function := value.GlobalFunctions[name]
	if function == nil {
		v.ErrorTable.NewSemanticError(ctx.ID().GetSymbol(), "Función '"+name+"' no está definida.")
		return value.NewNil()
	}

	// === Evaluar argumentos ===
	var argExprs []parser.IExpresionContext
	if ctx.Parametros() != nil {
		argExprs = ctx.Parametros().AllExpresion()
	}
	if len(argExprs) != len(function.Params) {
		v.ErrorTable.NewSemanticError(ctx.GetStart(),
			fmt.Sprintf("Número incorrecto de argumentos en llamada a '%s': esperados %d, recibidos %d",
				name, len(function.Params), len(argExprs)))
		return value.NewNil()
	}

	// === Crear nuevo scope para la función ===
	newScope := NewBaseScope("func_"+name, v.CurrentScope)

	for i, param := range function.Params {
		val := v.Visit(argExprs[i]).(value.IVOR)

		if !compareTypes(param.Type, val.Type()) {
			v.ErrorTable.NewSemanticError(ctx.GetStart(),
				fmt.Sprintf("Tipo incorrecto en parámetro '%s'. Esperado: %s, Recibido: %s",
					param.Name, param.Type, val.Type()))
			return value.NewNil()
		}

		if strings.HasPrefix(val.Type(), "slice_") || value.GlobalStructs[val.Type()] != nil {
			newScope.AddVariable(param.Name, param.Type, val, false, true, ctx.GetStart())
		} else {
			newScope.AddVariable(param.Name, param.Type, val.Copy(), false, true, ctx.GetStart())
		}
	}

	// === Ejecutar cuerpo de la función ===
	prev := v.CurrentScope
	v.CurrentScope = newScope
	result := v.Visit(function.Body)
	v.CurrentScope = prev

	// === Retorno ===
	if function.ReturnType != "" {
		if ret, ok := result.(ReturnSignal); ok {
			if ret.Value == nil {
				v.ErrorTable.NewSemanticError(ctx.GetStart(), "La función '"+name+"' debe retornar un valor.")
				return value.NewNil()
			}
			if !compareTypes(ret.Value.Type(), function.ReturnType) {
				v.ErrorTable.NewSemanticError(ctx.GetStart(),
					fmt.Sprintf("Tipo de retorno incorrecto en '%s'. Esperado: %s, Recibido: %s",
						name, function.ReturnType, ret.Value.Type()))
				return value.NewNil()
			}
			return ret.Value
		}
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "La función '"+name+"' debe retornar un valor.")
		return value.NewNil()
	}

	if ret, ok := result.(ReturnSignal); ok {
		return ret.Value
	}

	return value.NewNil()
}


func compareTypes(expected, actual string) bool {
	return strings.EqualFold(expected, actual)
}




func (v *ReplVisitor) visitParseFloat(args []parser.IExpresionContext, ctx antlr.ParserRuleContext) value.IVOR {
	if len(args) != 1 {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "parseFloat() espera exactamente 1 argumento")
		return value.NewNil()
	}

	val := v.Visit(args[0])
	ivor, ok := val.(value.IVOR)
	if !ok || ivor.Type() != value.IVOR_STRING {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "parseFloat() solo acepta strings como argumento")
		return value.NewNil()
	}

	str, ok := ivor.Value().(string)
	if !ok {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Valor no es string válido")
		return value.NewNil()
	}

	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "parseFloat() no pudo convertir '"+str+"' a float64")
		return value.NewNil()
	}

	return value.NewFloat(num)
}



func (v *ReplVisitor) visitTypeOf(args []parser.IExpresionContext, ctx antlr.ParserRuleContext) value.IVOR {
	defer func() {
		if r := recover(); r != nil {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), fmt.Sprintf("Error interno en typeOf: %v", r))
		}
	}()

	if len(args) != 1 {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "typeOf espera 1 argumento.")
		return value.NewNil()
	}

	val := v.Visit(args[0]).(value.IVOR)
	return value.NewString(val.Type())
}
