package arm

import (
	"fmt"
	"strings"
)

// Objeto que nos permita
// llevar la cuenta de las instrucciones
// este objeto se usa desde otro visitor
type ArmGenerator struct {
	Instructions []string
	Variables    map[string]SymbolInfo // ← ahora usa SymbolInfo para tipo + dirección
	Output       string
	VarOffset    int
	StdLib       *StandardLibrary

	tempRegs    []string
	tempIndex   int
	StringData  map[string]string
}


type SymbolInfo struct {
	Location string // Dirección en memoria (e.g. [SP, #8])
	Type     string // Tipo de dato: "int", "float64", "string", "bool"
}


var stringLabelCounter = 0

func (g *ArmGenerator) GenerateStringLabel() string {
	label := fmt.Sprintf(".str_%d", stringLabelCounter)
	stringLabelCounter++
	return label
}

func (g *ArmGenerator) AddData(label string, content string) {
	if g.StringData == nil {
		g.StringData = make(map[string]string)
	}

	g.StringData[label] = content
}





// Agregamos las cosas al string y el string es la traduccion

// Agregamos operaciones que o metan a un stack
// o agreguen al string lo que vamos haciendo
//

// funcion para agregar comnetarios al codigo
func (g *ArmGenerator) Comment(comment string) {
    g.Instructions = append(g.Instructions, "# "+comment)
}

/*
Agregar instrucciones
*/
func (g *ArmGenerator) Add(instruction string) {
	g.Instructions = append(g.Instructions, instruction)
}

/*
MUL ->
DIV ->
SUB ->
ADD ->
MOV ->
SVC -> Syscalls
*/

/*


 */

/*

 */

func (g *ArmGenerator) EndProgram() {
	g.StdLib.Use("end_program")
	g.Add("MOV X0, 0")  // Código de salida 0
	g.Add("MOV X8, 93") // Número de syscall para exit
	g.Add("SVC #0")     // Realizar syscall
}




/*
Funcion para sumar registros
*/

/*
// - Memory operations
*/
func Mov(g *ArmGenerator, rd, imm string) {
	g.Add(fmt.Sprintf("MOV %s, #%s", rd, imm))
}

func Ldr(g *ArmGenerator, rd, rs string) {
	g.Add(fmt.Sprintf("LDR %s, [%s]", rd, rs))
}
func Str(g *ArmGenerator, rs, rd string) {
	g.Add(fmt.Sprintf("STR %s, [%s]", rs, rd))
}

/*
push - pop
*/

func Push(g *ArmGenerator, rs string) {
	g.Add(fmt.Sprintf("STR %s, [SP, #-8]!", rs))
}

func Pop(g *ArmGenerator, rd string) {
	g.Add(fmt.Sprintf("LDR %s, [SP], #8", rd))
}

func Sub(g *ArmGenerator, rd, rs1, rs2 string) {
	g.Add(fmt.Sprintf("SUB %s, %s, %s", rd, rs1, rs2))
}

func Add(g *ArmGenerator, rd, rs1, rs2 string) {
	g.Add(fmt.Sprintf("ADD %s, %s, %s", rd, rs1, rs2))
}


/*
Impresion de las cosas
*/

func PrintInteger(g *ArmGenerator, rs string) {
	g.StdLib.Use("print_integer")
	g.Add(fmt.Sprintf("MOV X0, %s", rs))
	g.Add("BL print_integer")
}


func (g *ArmGenerator) String() string {
	return g.GetFullCode()
}


/*
Funcion para terminar el programa
syscall EndProgram

*/

func (g *ArmGenerator) GetOutput() string {
	return strings.Join(g.Instructions, "\n")
}


func (g *ArmGenerator) GetFullCode() string {
	header := `.global _start
.section .text
_start:
    BL main
    mov x8, #93
    svc #0
`

	textSection := g.GetOutput()
	libFunctions := g.StdLib.GetFunctionDefinitions()

	dataSection := ""
	if len(g.StringData) > 0 {
		dataSection += "\n.section .data\n"
		for label, content := range g.StringData {
			escaped := strings.ReplaceAll(content, `\`, `\\`)
			escaped = strings.ReplaceAll(escaped, `"`, `\"`)
			escaped = strings.ReplaceAll(escaped, "\n", `\n`)
			escaped = strings.ReplaceAll(escaped, "\t", `\t`)
			dataSection += fmt.Sprintf("%s:\n\t.asciz \"%s\"\n", label, escaped)
		}
	}

	// 🔁 Mueve funciones justo después del código principal
	return header + "\n" + textSection + "\n\n" + libFunctions + "\n" + dataSection
}



