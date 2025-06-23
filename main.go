package main



import (
	"fmt"
	"os"

	compiler "compiler/parser"
	repl "compiler/repl"

	"compiler/cst"
	"compiler/errors"
	// "main/repl"

	"github.com/antlr4-go/antlr/v4"
	"compiler/Frontend" 
)

func main() {
	// Mostrar interfaz web (si aplica)
	Frontend.MostrarIDE()

	// Leer código fuente desde stdin
	fmt.Println("Escribe el código de VLang (Ctrl+D para terminar):")
	inputCode, err := readStdin()
	if err != nil {
		fmt.Println("❌ Error leyendo entrada:", err)
		return
	}

	// Canal paralelo para CST (opcional)
	resultChannel := make(chan string)
	go func() {
		resultChannel <- cst.CstReport(inputCode)
	}()

	// === Análisis léxico ===
	input := antlr.NewInputStream(inputCode)
	lexer := compiler.NewVlangLexer(input)

	// Crear listener de errores léxicos
	lexicalErrorListener := errors.NewLexicalErrorListener()
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexicalErrorListener)

	// === Tokens ===
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// === Parser ===
	parser := compiler.NewVlangParser(stream)
	parser.BuildParseTrees = true

	// Listener y estrategia personalizada de errores sintácticos
	syntaxErrorListener := errors.NewSyntaxErrorListener(lexicalErrorListener.ErrorTable)
	parser.RemoveErrorListeners()
	parser.SetErrorHandler(errors.NewCustomErrorStrategy())
	parser.AddErrorListener(syntaxErrorListener)

	// Parsear el árbol desde el axioma principal
	tree := parser.Programa()

	// === Crear consola y visitor ===
	console := repl.NewConsole()
	visitor := repl.NewReplVisitor(console)

	// Visitar árbol
	visitor.Visit(tree)

	// === Mostrar errores léxicos ===
	if len(lexicalErrorListener.ErrorTable.Errors) > 0 {
		fmt.Println("🛑 Errores léxicos:")
		for _, err := range lexicalErrorListener.ErrorTable.Errors {
			fmt.Println(err.String())
		}
		return
	}

	// === Mostrar errores sintácticos/semánticos ===
	if len(visitor.ErrorTable.Errors) > 0 {
		fmt.Println("🛑 Errores sintácticos o semánticos:")
		for _, err := range visitor.ErrorTable.Errors {
			fmt.Println(err.String())
		}
		return
	}

	// === Mostrar salida final (si no hay errores) ===
	fmt.Println("✅ Ejecución finalizada sin errores. Salida:")
	fmt.Println(console.GetOutput())
}

func readStdin() (string, error) {
	input, err := os.ReadFile("/dev/stdin")
	return string(input), err
}

// Funciones para visualizar nuestro arbol
func PrintVerticalTree(node antlr.Tree, ruleNames []string) {
	printVerticalNode(node, ruleNames, "", true)
}

func printVerticalNode(node antlr.Tree, ruleNames []string, prefix string, isLast bool) {
	connector := "+-- "
	if !isLast {
		connector = "|-- "
	}

	var label string
	switch n := node.(type) {
	case antlr.RuleNode:
		label = ruleNames[n.GetRuleContext().GetRuleIndex()]
	case antlr.TerminalNode:
		label = fmt.Sprintf("\"%s\"", n.GetText())
	default:
		label = fmt.Sprintf("%T", n)
	}

	fmt.Printf("%s%s%s\n", prefix, connector, label)

	// Actualizar el prefijo para los hijos
	childCount := node.GetChildCount()
	for i := 0; i < childCount; i++ {
		child := node.GetChild(i)
		newPrefix := prefix
		if isLast {
			newPrefix += "    "
		} else {
			newPrefix += "|   "
		}
		printVerticalNode(child, ruleNames, newPrefix, i == childCount-1)
	}
}