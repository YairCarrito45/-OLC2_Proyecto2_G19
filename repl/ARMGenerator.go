package repl

import (
	"fmt"
	"os"
	"strings"
)

type ARMGenerator struct {
	Console      *Console
	Instructions []string
}

// === Instrucciones de bajo nivel ===

func (g *ARMGenerator) AddInstruction(instruction string) {
	g.Instructions = append(g.Instructions, instruction)
}

func (g *ARMGenerator) Comment(comment string) {
	g.AddInstruction(fmt.Sprintf("// %s", comment))
}

func (g *ARMGenerator) Mov(rd string, imm int) {
	g.AddInstruction(fmt.Sprintf("MOV %s, #%d", rd, imm))
}

func (g *ARMGenerator) Add(rd, rs1, rs2 string) {
	g.AddInstruction(fmt.Sprintf("ADD %s, %s, %s", rd, rs1, rs2))
}

func (g *ARMGenerator) Sub(rd, rs1, rs2 string) {
	g.AddInstruction(fmt.Sprintf("SUB %s, %s, %s", rd, rs1, rs2))
}

func (g *ARMGenerator) Mul(rd, rs1, rs2 string) {
	g.AddInstruction(fmt.Sprintf("MUL %s, %s, %s", rd, rs1, rs2))
}

func (g *ARMGenerator) Div(rd, rs1, rs2 string) {
	g.AddInstruction(fmt.Sprintf("SDIV %s, %s, %s", rd, rs1, rs2))
}

func (g *ARMGenerator) Addi(rd, rs1 string, imm int) {
	g.AddInstruction(fmt.Sprintf("ADDI %s, %s, #%d", rd, rs1, imm))
}

// === Memoria y pila ===

func (g *ARMGenerator) Str(rs1, rs2 string, offset int) {
	g.AddInstruction(fmt.Sprintf("STR %s, [%s, #%d]", rs1, rs2, offset))
}

func (g *ARMGenerator) Ldr(rd, rs1 string, offset int) {
	g.AddInstruction(fmt.Sprintf("LDR %s, [%s, #%d]", rd, rs1, offset))
}

func (g *ARMGenerator) Push(rs string) {
	g.AddInstruction(fmt.Sprintf("STR %s, [SP, #-8]!", rs))
}

func (g *ARMGenerator) Pop(rd string) {
	g.AddInstruction(fmt.Sprintf("LDR %s, [SP], #8", rd))
}

// === Syscalls ===

func (g *ARMGenerator) Svc() {
	g.AddInstruction("SVC #0")
}

// === Utilidades ===

func (g *ARMGenerator) EndProgram() {
	g.Mov("X0", 0)
	g.Mov("X8", 93) // syscall exit
	g.Svc()
}

// Ejemplo simple de impresión (puedes conectar esto a stdlib)
func (g *ARMGenerator) PrintInteger(rs string) {
	g.AddInstruction(fmt.Sprintf("MOV X0, %s", rs))
	g.AddInstruction("BL print_integer") // debe existir una función externa
}

// === Exportar código a archivo ===

func (g *ARMGenerator) ExportToFile(filename string) error {
	var sb strings.Builder

	sb.WriteString(".text\n.global _start\n_start:\n")

	for _, instr := range g.Instructions {
		sb.WriteString(instr + "\n")
	}

	// Epílogo opcional (comentado si ya usaste EndProgram manualmente)
	// g.EndProgram()

	return os.WriteFile(filename, []byte(sb.String()), 0644)
}
