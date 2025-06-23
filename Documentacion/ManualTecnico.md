# Manual Tecnico - Proyecto #1 VLangCherry 

### Grupo 19
- Estiben Yair Lopez Leveron - 202204578
- Harold Alejandro Sánchez Hernández - 202200100
- Johan Moises Cardona Rosales - 202201405
-------
## 1. Introduccion 

Este documento tiene como objetivo explicar la implementación técnica del intérprete VLangCherry, un sistema diseñado para permitir la edición, análisis y ejecución de programas escritos en dicho lenguaje, simulando el flujo de trabajo de compiladores reales. Está dirigido a profesores, auxiliares y desarrolladores interesados en comprender, mantener o extender el sistema, proporcionando una visión detallada de su arquitectura, componentes y funcionamiento interno.

## 2. Arquitectura del Sistema

- Lexer (Scanner): Generado con ANTLR4.

- Parser (Parser): Gramática definida en ANTLR4 para analizar la - estructura del código.

- AST (Árbol de Sintaxis Abstracta): Estructura de nodos que representan el código analizado.

- Análisis Semántico: Verifica tipos, declaraciones, accesos válidos, etc.

- Intérprete: Recorre el AST y ejecuta sentencias.

- Reportes: Errores, Tabla de Símbolos, AST.

- Interfaz (si aplica): Editor, consola de salida, área de reportes.

## 3. Tecnologías Utilizadas
- Lenguaje: **Go**
- Generador de gramática: **ANTLR4**
- Sistema operativo: **Linux**
- GUI: **Fyne**

## 4. Estructura del Proyecto

#### parser:
- **Vlang.g4:** Archivo fuente de gramática en ANTLR4.

- **Vlang.tokens, Vlang.interp:** Archivos generados por ANTLR para el parser.

- **vlang_lexer.go, vlang_parser.go:** Código fuente generado para el lexer y parser en Go.

- **vlang_base_visitor.go, vlang_visitor.go:** Implementación del patrón Visitor para recorrer el árbol sintáctico.

- **vlang_listener.go, vlang_base_listener.go:**  Implementación opcional del patrón Listener (no utilizada si se trabaja solo con Visitor).

#### repl:
- **argument.go:** Gestión de argumentos de funciones.

- **console.go:** Implementación de salida por consola.

- **error_table.go:** Registro y generación del reporte de errores.

- **scope.go:** Gestión de entornos y visibilidad de variables (manejo de ámbitos).

- **variable.go:** Declaración, acceso y actualización de variables.

- **visitor.go:** Visitor principal que recorre el AST y ejecuta instrucciones.

##### value:
- **base_value.go:** Interfaz común para todos los tipos de valores.

- **int_value.go, float_value.go, bool_value.go, char_value.go, string_value.go:** Tipos primitivos y su comportamiento.

- **nil_value.go:** Representación del valor nil.

- **slice_value.go:** Implementación de slices y slices multidimensionales.


## 5. Explicación de Estructura

#### Vlang.g4 — Definición de la gramática

```bash
programa : declaraciones* EOF ;
declaraciones : varDcl | stmt ;
```

* Programa es la regla inicial del lenguaje. Indica que un programa válido está compuesto por múltiples declaraciones y finaliza con EOF (fin de archivo).

- Cada declaraciones puede ser una declaración de variable o una sentencia ejecutable.


**Declaración de variables**
```bash
varDcl
    : MUT ID tipo? ((ASSIGN | ASSIGN_DECL) expresion)? #variableDeclaration
    ;
```

* Define la sintaxis para declarar una variable.

- Palabra clave obligatoria: mut.

- Puede tener un tipo explícito (int, string, etc.) y un valor inicial.

- Soporta asignación estándar (=) o declaración con inferencia (:=).

**Tipos de datos permitidos**
```bash
tipo : sliceTipo | 'int' | 'float64' | 'string' | 'bool' | 'char' ;
sliceTipo : LSQUARE RSQUARE tipo ;
```
* Se permiten tipos primitivos como int, float64, string, bool, char.

* También se admite la definición de slices (arreglos dinámicos), incluyendo slices de slices para estructuras multidimensionales.


**Sentencias de control y bloques**
```bash
sentencias_control
    : ifDcl         #if_context 
    | forDcl        #for_context 
    | whileDcl      #while_context
    | switchDcl     #switch_context 
    ;
```
* Se agrupan todas las estructuras de control de flujo: if, for, while, switch.


**Estructura If**
```bash
ifDcl
    : IF LPAREN? expresion RPAREN? bloque (elseifDcl)* (elseDcl)?
    ;
```

**Estructura for: clásico, condicional, foreach**
```bash
forDcl
    : FOR expresion bloque                                 #forCondicional
    | FOR expresion SEMICOLON expresion SEMICOLON expresion bloque   #forClasico
    | FOR ID COMMA ID IN ID bloque                         #forForeach
    ;
```

El lenguaje soporta tres variantes:

- Condicional (como while)

- Clásico (for init; cond; inc)

- Iteración sobre slices (for i, v in arreglo)

**Expresiones**
```bash
expresion
    : ID PLUSEQ expresion                                #asignacionCompuestaSuma
    | expresion op=(MUL | DIV | MOD) expresion           #multdivmod
    | expresion op=(EQ | NEQ) expresion                  #igualdad
    | valor                                              #valorexpr
    | ID                                                 #id
    ;
```
- La regla expresion permite construir operaciones aritméticas, lógicas, llamadas a funciones, expresiones anidadas, etc.

- Soporta operadores compuestos (+=, -=), unarios, de comparación (==, !=), y más.

**Literales y valores**
```bash
valor
    : ENTERO    #valorEntero
    | DECIMAL   #valorDecimal
    | CADENA    #valorCadena
    | BOOLEANO  #valorBooleano
    | CARACTER  #valorCaracter
    ;
```

**Tokens léxicos**
```bash
ID       : [_a-zA-Z][_a-zA-Z0-9]* ;
BOOLEANO : 'true' | 'false' ;
ENTERO   : [0-9]+ ;
CADENA   : '"' (~["\\] | '\\' .)* '"' ;
```

- Definen los tokens léxicos del lenguaje.

- El identificador (ID) acepta letras, números y guiones bajos, pero debe comenzar con letra o guion bajo.


**Ignorar comentarios y espacios**
```bash
WS            : [ \t\r\n]+ -> skip ;
LINE_COMMENT  : '//' ~[\r\n]* -> skip ;
BLOCK_COMMENT : '/*' .*? '*/' -> skip ;
```

- Espacios en blanco y comentarios de línea o bloque son ignorados durante el análisis.


#### Visitor.go (Interprete)

**Fragmento clave: constructor e inicialización**
```bash
type ReplVisitor struct {
	parser.BaseVlangVisitor
	CurrentScope *BaseScope
	ErrorTable   *ErrorTable
	Console      *Console
}

func NewReplVisitor(console *Console) *ReplVisitor {
	return &ReplVisitor{
		CurrentScope: NewBaseScope("global", nil),
		ErrorTable:   NewErrorTable(),
		Console:      console,
	}
}

```
**Recorre todas las declaraciones del programa y las ejecuta una por una.**
```bash
func (v *ReplVisitor) VisitPrograma(ctx *parser.ProgramaContext) interface{} {
	for _, decl := range ctx.AllDeclaraciones() {
		v.Visit(decl)
	}
	return nil
}
```
**Declaraciones y expresiones** 

***Variables***

- Valida si el nombre está reservado.

- Infieren o verifican el tipo.

- Se almacenan en el ámbito actual.

```bash
func (v *ReplVisitor) VisitVariableDeclaration(ctx *parser.VariableDeclarationContext) interface{} {
	// Declaración con tipo y/o valor opcional
	// Manejo de error si el tipo no coincide
}
```
**Evaluación de expresiones**

Se implementan operaciones con coerción y validación de tipos:

- Sumas, restas, multiplicaciones, divisiones, módulos

- Comparaciones relacionales y lógicas

- Negación unaria (-, !)

- Igualdad (==, !=)

- Acceso a variables (VisitId)

- Asignaciones compuestas (+=, -=)

```bash
func (v *ReplVisitor) VisitSumres(ctx *parser.SumresContext) interface{} {
	// Permite suma entre int, float y string según casos válidos
}
```
**Control de flujo**

If - Else If - Else
```bash
func (v *ReplVisitor) VisitIfDcl(ctx *parser.IfDclContext) interface{} {
	// Evalúa condición, ejecuta bloque correspondiente
}

func (v *ReplVisitor) VisitSwitch_context(ctx *parser.Switch_contextContext) interface{} {
	// Evalúa expresión y ejecuta el case que coincide
}

func (v *ReplVisitor) VisitForClasico(ctx *parser.ForClasicoContext) interface{}
```

**Control de ejecución**

Soporte para sentencias:

- break: sólo válido en for y switch.

- continue: válido solo dentro de for.

- return: retorna un valor (si aplica) desde funciones.

**Slices**
```bash
func (v *ReplVisitor) VisitSliceCreacion(ctx *parser.SliceCreacionContext) interface{}
```

**Gestión de tipos**
Define funciones de ayuda:

```bash
func defaultValueForType(tipo string) value.IVOR
func NormalizeTypeName(name string) string
```

#### Scope.go

Cada scope (ámbito) es un entorno que contiene variables declaradas localmente. También puede tener un scope padre (herencia léxica).

```bash
type BaseScope struct {
	Name      string
	Parent    *BaseScope
	Variables map[string]*Variable
}
```
- Name: identificador del scope (ej. "global", "for_clasico", etc.).

- Parent: referencia al scope padre (si es anidado).

- Variables: mapa con las variables declaradas en ese ámbito.


**Creación de scopes**

El método NewBaseScope permite crear un nuevo entorno:
```bash
func NewBaseScope(name string, parent *BaseScope) *BaseScope {
	return &BaseScope{
		Name:      name,
		Parent:    parent,
		Variables: make(map[string]*Variable),
	}
}
```

**Validación de existencia de variables**
```bash
func (s *BaseScope) variableExists(name string) bool {
	_, exists := s.Variables[name]
	return exists
}

```

- Verifica si la variable ya existe dentro del mismo scope (no en el padre).

- Se usa para evitar redefiniciones locales.


**Agregar una variable**
```bash
func (s *BaseScope) AddVariable(name string, varType string, val value.IVOR, isConst bool, allowNil bool, token antlr.Token) (*Variable, string)

```

- Crea una instancia del tipo Variable.

- La agrega al mapa Variables.

- La variable creada (*Variable).

- Un mensaje de error (string) si ya estaba declarada.

**Buscar una variable**
```bash
func (s *BaseScope) GetVariable(name string) *Variable {
	scope := s
	for scope != nil {
		if v, ok := scope.Variables[name]; ok {
			return v
		}
		scope = scope.Parent
	}
	return nil
}
```
- Busca la variable de forma recursiva desde el scope actual hacia los padres.

- Soporta sombras de variables (variable local oculta una global con el mismo nombre).

- Retorna nil si no se encuentra en ningún nivel.


#### variable.go

Define la estructura base de una variable dentro del intérprete de VLangCherry. Es el modelo que encapsula toda la información necesaria para que una variable sea reconocida, evaluada y validada semánticamente en un ámbito 

```bash
type Variable struct {
	Name     string
	Value    value.IVOR
	Type     string
	IsConst  bool
	AllowNil bool
	Token    antlr.Token
}
```
- **Name:** Nombre de la variable (identificador en el código fuente).
- **Value:** Valor actual de la variable (implementado usando la interfaz value.IVOR).
- **Type:** Tipo de dato declarado o inferido (ej. int, string, []int).
- **IsConst:** Indica si la variable es constante (no se puede modificar tras la asignación inicial).
- **AllowNil:** Permite valores nil si es true (útil para variables opcionales o no inicializadas).
- **Token:** Token original ANTLR, útil para manejo de errores (incluye posición en el código fuente).


#### error_table.go

Define el sistema de manejo y registro de errores que puede surgir durante la ejecución del intérprete, centraliza el control de errores léxicos, sintácticos, semánticos y de tiempo de ejecución.

```bash 
type Error struct {
	Line   int
	Column int
	Msg    string
	Type   string
}
```

Cada error registrado contiene:

- Línea y columna de aparición.

- Mensaje explicativo del error.

- Tipo de error (léxico, sintáctico, semántico, runtime).

**Tabla de errores: ErrorTable**

```bash
type ErrorTable struct {
	Errors []Error
}
```
- La tabla mantiene una lista acumulativa de errores capturados, lo que permite reportar múltiples problemas en una misma ejecución.



#### base_value

```bash
type IVOR interface {
	Value() interface{}
	Type() string
	Copy() IVOR
}
```

Esta interfaz permite que todos los tipos (int, float, string, bool, etc.) puedan ser manipulados homogéneamente por el intérprete. Cada implementación concreta proporciona:

- Value(): el valor crudo (ej. 42, true, "Hola").

- Type(): el tipo como string (ej. "Int").

- Copy(): una copia profunda del valor, evitando aliasing o efectos colaterales.

**Constantes de tipo**
```bash
const (
	IVOR_INT = "Int"
	IVOR_FLOAT = "Float"
	...
)
```
Estas constantes se usan como etiquetas internas de tipo. Incluyen tanto tipos primitivos como estructuras más avanzadas


**Constructores de tipos**

El archivo también define funciones para crear valores:

```bash
func NewInt(val int) IVOR         // Devuelve IntValue
func NewFloat(val float64) IVOR   // Devuelve FloatValue
func NewString(val string) IVOR   // Devuelve StringValue
func NewBool(val bool) IVOR       // Devuelve BoolValue
func NewNil() IVOR                // Devuelve NilValue (único)
```
Estas funciones son utilizadas por el visitor (ReplVisitor) cuando se declaran o evalúan expresiones.


#### int_value.go — Tipo entero (int)
```bash
type IntValue struct {
	InternalValue int
}
```
Métodos implementados:
- **Value():** Devuelve el valor entero como interface{}.

- **Type():** Retorna "Int" (constante IVOR_INT).

- **Copy():** Retorna una nueva instancia del mismo valor, evitando aliasing.

Se utiliza para representar variables, resultados de expresiones y literales enteros en tiempo de ejecución.


#### float_value.go — Tipo flotante (float64)

```bash
type FloatValue struct {
	InternalValue float64
}
```

Métodos implementados:
- **Value():** Retorna el número flotante.

- **Type():** Devuelve "Float" (constante IVOR_FLOAT).

- **Copy():** Crea una copia idéntica del valor.

Se usa para expresiones decimales, conversiones implícitas y operaciones aritméticas mixtas con int.


#### slice_value.go 

```bash
type SliceValue struct {
	ElementType string
	Elements    []IVOR
}
```

Métodos implementados:

- **Value():** Devuelve el arreglo como []IVOR.

- **Type():** Retorna "slice_<tipo>", por ejemplo "slice_Int".

- **Copy():** Clona el slice y todos sus elementos usando Copy() recursivo.

- **Append(val IVOR):** Agrega un elemento si el tipo coincide; de lo contrario, lanza un error.

Este tipo permite trabajar con slices homogéneos, incluidos slices de slices, gracias al uso del sistema de tipos dinámicos.


## 6. Conclusión

El desarrollo del intérprete VLangCherry representa una implementación completa de un lenguaje de programación estructurado, diseñado desde cero utilizando herramientas modernas como ANTLR4 y Go. A lo largo de este manual, se ha documentado detalladamente la arquitectura del sistema, los módulos funcionales, la semántica de tipos, el control de flujo y la gestión de errores, todos elementos esenciales para un lenguaje dinámico, seguro y extensible.


- Interpretar código fuente escrito en VLangCherry mediante un árbol de sintaxis abstracta (AST).

- Ejecutar estructuras de control (if, for, while, switch) y operaciones complejas con slices y expresiones.

- Gestionar variables con un sistema robusto de ámbitos (scopes) y tipos (int, float, string, bool, char, slice).

- Proveer retroalimentación clara mediante un sistema de consola interna y una tabla de errores léxicos, sintácticos, semánticos y de ejecución.

- Integrarse con una interfaz visual o usarse en consola, haciendo el entorno flexible tanto para el usuario final como para fines educativos.