package main

import (
	"fmt"
	"os"

	compiler "compiler/parser"
	repl "compiler/repl"

	"compiler/cst"
	"compiler/errors"
	"compiler/arm"
	"github.com/antlr4-go/antlr/v4"
	"compiler/Frontend"
)

func main() {
	// Mostrar interfaz gráfica (opcional)
	Frontend.MostrarIDE()

	// Leer código fuente desde stdin
	fmt.Println("Escribe el código de VLang (Ctrl+D para terminar):")
	inputCode, err := readStdin()
	if err != nil {
		fmt.Println("❌ Error leyendo entrada:", err)
		return
	}

	// Iniciar generación del CST en paralelo (opcional)
	cstOutput := make(chan string)

	go func() {
		cstOutput <- cst.CstReport(inputCode)
	}()

	// === Análisis léxico ===
	input := antlr.NewInputStream(inputCode)
	lexer := compiler.NewVlangLexer(input)

	// Escuchar errores léxicos
	lexicalErrorListener := errors.NewLexicalErrorListener()
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexicalErrorListener)

	// Tokenizar
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// === Parser ===
	parser := compiler.NewVlangParser(stream)
	parser.BuildParseTrees = true

	// Escuchar errores sintácticos
	syntaxErrorListener := errors.NewSyntaxErrorListener(lexicalErrorListener.ErrorTable)
	parser.RemoveErrorListeners()
	parser.SetErrorHandler(errors.NewCustomErrorStrategy())
	parser.AddErrorListener(syntaxErrorListener)

	// Parsear el árbol desde el axioma principal
	tree := parser.Programa()

	// === Crear consola y visitor REPL ===
	console := repl.NewConsole()
	visitor := repl.NewReplVisitor(console)
	visitor.Visit(tree)

	// === Validar errores antes de traducir ===
	if len(lexicalErrorListener.ErrorTable.Errors) > 0 {
		fmt.Println("🛑 Errores léxicos:")
		for _, err := range lexicalErrorListener.ErrorTable.Errors {
			fmt.Println(err.String())
		}
		return
	}

	if len(visitor.ErrorTable.Errors) > 0 {
		fmt.Println("🛑 Errores sintácticos o semánticos:")
		for _, err := range visitor.ErrorTable.Errors {
			fmt.Println(err.String())
		}
		return
	}

// ... (todo lo anterior igual)

	// === Iniciar traducción a ARM64 ===
	armVisitor := arm.NewArmVisitor()
	armVisitor.Visit(tree)

	// Obtener instrucciones y funciones estándar necesarias
	armInstructions := armVisitor.Generator.GetOutput()
	stdLibCode := armVisitor.Generator.StdLib.GetFunctionDefinitions()

	if len(armVisitor.Generator.Instructions) == 0 {
		fmt.Println("⚠️ No se generaron instrucciones ARM.")
		return
	}

	// Encabezado ARM64 correcto
	armHeader := `.global _start
.section .text
_start:
    BL main
    mov x8, #93     // syscall: exit
    svc #0

`

	// Concatenar todo
	armFullCode := armHeader + armInstructions + "\n\n" + stdLibCode

	// Escribir archivo
	outputFile := "salida.s"
	err = os.WriteFile(outputFile, []byte(armFullCode), 0644)
	if err != nil {
		fmt.Println("❌ Error escribiendo archivo ARM:", err)
		return
	}

	fmt.Println("✅ Código ARM generado correctamente en", outputFile)

	// Mostrar salida interpretada
	fmt.Println("✅ Ejecución finalizada sin errores. Salida:")
	fmt.Println(console.GetOutput())

}

// Utilidad para leer stdin
func readStdin() (string, error) {
	input, err := os.ReadFile("/dev/stdin")
	return string(input), err
}
