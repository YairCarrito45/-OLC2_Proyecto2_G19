package Frontend

import (
	"bytes"

	"fmt"
	"os/exec"

	"os"
	"strings"

	"compiler/errors"
	parser "compiler/parser"
	"compiler/repl"
	_ "image/png"


	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/antlr4-go/antlr/v4"
)

var monoStyle = fyne.TextStyle{Monospace: true}

func styledLabel(text string, bold bool) *widget.Label {
	return widget.NewLabelWithStyle(text, fyne.TextAlignLeading, fyne.TextStyle{Bold: bold, Monospace: true})
}

func stylizeEntry(entry *widget.Entry) {
	entry.TextStyle = monoStyle
}

func MostrarIDE() {
	a := app.New()
a.Settings().SetTheme(DarkMonoTheme{})


	w := a.NewWindow("VLan Cherry IDE")
	w.Resize(fyne.NewSize(1100, 750))

	editor := widget.NewMultiLineEntry()
	editor.SetPlaceHolder("Escribe tu código VLan Cherry aquí...")
	editor.Wrapping = fyne.TextWrapWord
	stylizeEntry(editor)

	editorScroll := container.NewScroll(editor)
	editorScroll.SetMinSize(fyne.NewSize(400, 300))

	entradaBox := container.NewBorder(
		styledLabel("Entrada:", true),
		nil, nil, nil,
		editorScroll,
	)

	console := widget.NewMultiLineEntry()
	console.SetText("#------ Estiben Yair Lopez Leveron ------ 202204578 ----\n")
	console.SetText(console.Text + "#------ Harold Alejandro Sanchez Hernandez------ 202200100 ----\n")
	console.SetText(console.Text + "#------ Johan Moises Cardona Rosales  ------ 202201405 ----")
	console.Disable()
	stylizeEntry(console)

	consoleScroll := container.NewScroll(console)
	consoleScroll.SetMinSize(fyne.NewSize(400, 300))

	consolaBox := container.NewBorder(
		styledLabel("Consola:", false),
		nil, nil, nil,
		consoleScroll,
	)

	var currentFile fyne.URI

	btnNuevo := widget.NewButton("Nuevo", func() {
		editor.SetText("")
		currentFile = nil
	})

	btnAbrir := widget.NewButton("Abrir", func() {
		openDialog := dialog.NewFileOpen(func(r fyne.URIReadCloser, err error) {
			if err == nil && r != nil {
				data, _ := os.ReadFile(r.URI().Path())
				editor.SetText(string(data))
				currentFile = r.URI()
			}
		}, w)
		openDialog.SetFilter(storage.NewExtensionFileFilter([]string{".vch"}))
		openDialog.Show()
	})

	btnGuardar := widget.NewButton("Guardar", func() {
		if currentFile == nil {
			saveDialog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
				if writer != nil {
					writer.Write([]byte(editor.Text))
					currentFile = writer.URI()
				}
			}, w)
			saveDialog.SetFileName("archivo.vch")
			saveDialog.Show()
		} else {
			f, err := os.Create(currentFile.Path())
			if err == nil {
				f.WriteString(editor.Text)
				f.Close()
			}
		}
	})

	erroresEntry := widget.NewMultiLineEntry()
	erroresEntry.SetText("No se detectaron errores.")
	erroresEntry.Disable()
	stylizeEntry(erroresEntry)

	var erroresData = [][]string{
		{"#", "Tipo", "Línea", "Col", "Descripción"},
	}

	erroresTabla := widget.NewTable(
		func() (int, int) { return len(erroresData), len(erroresData[0]) },
		func() fyne.CanvasObject {
			lbl := widget.NewLabel("")
			lbl.TextStyle = fyne.TextStyle{Monospace: true}
			return lbl
		},
		func(id widget.TableCellID, o fyne.CanvasObject) {
			lbl := o.(*widget.Label)
			lbl.SetText(erroresData[id.Row][id.Col])
			if id.Row > 0 { // Filas de datos (no encabezados)
				errType := erroresData[id.Row][1]
				switch errType {
				case "Léxico":
					lbl.TextStyle = fyne.TextStyle{Monospace: true, Bold: true}
				case "Sintáctico":
					lbl.TextStyle = fyne.TextStyle{Monospace: true, Italic: true}
				default:
					lbl.TextStyle = fyne.TextStyle{Monospace: true}
				}
			}
		},
	)

	var columnWidths = []float32{50, 200, 80, 80, 400}
	for i, width := range columnWidths {
		erroresTabla.SetColumnWidth(i, width)
	}

	var tablaSimbolosData = [][]string{
		{"ID", "Tipo", "Valor", "Ámbito", "Línea", "Columna"},
	}

	tablaSimbolos := widget.NewTable(
		func() (int, int) { return len(tablaSimbolosData), len(tablaSimbolosData[0]) },
		func() fyne.CanvasObject {
			lbl := widget.NewLabel("")
			lbl.TextStyle = monoStyle
			return lbl
		},
		func(id widget.TableCellID, o fyne.CanvasObject) {
			lbl := o.(*widget.Label)
			if id.Row < len(tablaSimbolosData) && id.Col < len(tablaSimbolosData[id.Row]) {
				lbl.SetText(tablaSimbolosData[id.Row][id.Col])
				if id.Row == 0 {
					lbl.TextStyle = fyne.TextStyle{Monospace: true, Bold: true}
				} else {
					lbl.TextStyle = monoStyle
				}
			}
		},
	)

	symbolColumnWidths := []float32{150, 200, 200, 80, 80}
	for i, width := range symbolColumnWidths {
		tablaSimbolos.SetColumnWidth(i, width)
	}



	reportTabs := container.NewAppTabs(
		container.NewTabItem("Errores", container.NewScroll(erroresTabla)),
		container.NewTabItem("Símbolos", container.NewScroll(tablaSimbolos)),
		container.NewTabItem("AST", container.NewVBox()),
	)

	btnEjecutar := widget.NewButton("Ejecutar", func() {
		codigo := editor.Text

		lexListener := errors.NewLexicalErrorListener()
		lexer := parser.NewVlangLexer(antlr.NewInputStream(codigo))
		lexer.RemoveErrorListeners()
		lexer.AddErrorListener(lexListener)

		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		syntListener := errors.NewSyntaxErrorListener(lexListener.ErrorTable)
		p := parser.NewVlangParser(stream)
		p.BuildParseTrees = true
		p.RemoveErrorListeners()
		p.SetErrorHandler(errors.NewCustomErrorStrategy())
		p.AddErrorListener(syntListener)

		tree := p.Programa()

		consoleObj := repl.NewConsole()
		visitor := repl.NewReplVisitor(consoleObj)
		visitor.Visit(tree)

		console.SetText(consoleObj.GetOutput())

		errores := append(lexListener.ErrorTable.Errors, syntListener.ErrorTable.Errors...)
		errores = append(errores, visitor.ErrorTable.Errors...)

		erroresData = [][]string{{"#", "Tipo", "Línea", "Col", "Descripción"}}
		if len(errores) == 0 {
			erroresEntry.SetText("No se encontraron errores.")
			erroresData = append(erroresData, []string{"", "", "", "", "No se detectaron errores."})
		} else {
			var builder strings.Builder
			builder.WriteString("Se detectaron errores:\n\n")
			unique := make(map[string]bool)
			for i, err := range errores {
				key := fmt.Sprintf("%s:%d:%d:%s", err.Type, err.Line, err.Column, err.Msg)
				if !unique[key] {
					unique[key] = true
					builder.WriteString(fmt.Sprintf("🔸 [%s] Línea %d, Col %d → %s\n", err.Type, err.Line, err.Column, err.Msg))
					erroresData = append(erroresData, []string{
						fmt.Sprintf("%d", i+1),
						err.Type,
						fmt.Sprintf("%d", err.Line),
						fmt.Sprintf("%d", err.Column),
						err.Msg,
					})
				}
			}
			erroresEntry.SetText(builder.String())
		}
		erroresTabla.Refresh()

		tablaSimbolosData = [][]string{{"ID", "Tipo", "Valor", "Ámbito", "Línea", "Columna"}}
		for _, symbol := range visitor.Symbols {
			tablaSimbolosData = append(tablaSimbolosData, []string{
				symbol.ID,
				symbol.Tipo,
				symbol.Valor,
				symbol.Ambito,
				fmt.Sprintf("%d", symbol.Linea),
				fmt.Sprintf("%d", symbol.Columna),
			})
		}
		tablaSimbolos.Refresh()



		// === Generar AST localmente con DOT + Graphviz ===
		dot := GenerateTreeDOT(tree, p)

		if err := os.WriteFile("ast.dot", []byte(dot), 0644); err != nil {
			fmt.Println("❌ Error escribiendo archivo .dot:", err)
			return
		}

		cmd := exec.Command("dot", "-Tpng", "-Gdpi=150", "ast.dot", "-o", "ast_local.png") // resolución alta
		if err := cmd.Run(); err != nil {
			fmt.Println("❌ Error ejecutando Graphviz:", err)
			return
		}

		imgData, err := os.ReadFile("ast_local.png")
		if err != nil {
			fmt.Println("❌ Error leyendo imagen:", err)
			return
		}

		r := bytes.NewReader(imgData)
		img := canvas.NewImageFromReader(r, "ast_local.png")
		img.FillMode = canvas.ImageFillOriginal // Mostrar tamaño real

		// === Zoom dinámico ===
		scaleFactor := float32(1.0)
			img.Resize(fyne.NewSize(1000*scaleFactor, 700*scaleFactor))

		zoomInBtn := widget.NewButton("+ Zoom", func() {
			scaleFactor += 0.2
			img.Resize(fyne.NewSize(1000*scaleFactor, 700*scaleFactor))
		})

		zoomOutBtn := widget.NewButton("- Zoom", func() {
			if scaleFactor > 0.4 {
				scaleFactor -= 0.2
				img.Resize(fyne.NewSize(1000*scaleFactor, 700*scaleFactor))
			}
		})

		// Contenedor de imagen con scroll
		scroll := container.NewScroll(img)
		scroll.SetMinSize(fyne.NewSize(1000, 700))

		zoomControls := container.NewHBox(zoomInBtn, zoomOutBtn)
		astContainer := container.NewVBox(zoomControls, scroll)

		reportTabs.Items[2].Content = astContainer
		reportTabs.Refresh()

	})



	toolbar := container.NewHBox(btnNuevo, btnAbrir, btnGuardar, btnEjecutar)

	horizontal := container.NewHSplit(entradaBox, consolaBox)
	horizontal.Offset = 0.5

	principal := container.NewVSplit(horizontal, reportTabs)
	principal.Offset = 0.4

	w.SetContent(container.NewBorder(toolbar, nil, nil, nil, principal))
	w.ShowAndRun()
}

func PrintVerticalTree(node antlr.Tree, ruleNames []string, prefix string, isLast bool, out *strings.Builder) {
	connector := "+-- "
	if !isLast {
		connector = "|-- "
	}

	var label string
	switch n := node.(type) {
	case antlr.RuleNode:
		ruleIndex := n.GetRuleContext().GetRuleIndex()
		if ruleIndex >= 0 && ruleIndex < len(ruleNames) {
			label = ruleNames[ruleIndex]
			switch label {
			case "programa":
				label = "programa"
			case "funcion":
				label = "funcion"
			case "declaraciones":
				label = "declaraciones"
			case "sentencia":
				label = "sentencia"
			case "expresion":
				label = "expresion"
			case "asignacion":
				label = "asignacion"
			case "retorno":
				label = "retorno"
			default:
				if strings.HasPrefix(label, "valor") || strings.HasPrefix(label, "tipo") {
					label = "valor"
				} else {
					label = "nodo"
				}
			}
		} else {
			label = "nodo"
		}
	case antlr.TerminalNode:
		text := n.GetText()
		if text == ";" || text == "(" || text == ")" || text == "{" || text == "}" {
			return
		}
		label = fmt.Sprintf("\"%s\"", text)
	default:
		label = "nodo"
	}

	out.WriteString(fmt.Sprintf("%s%s%s\n", prefix, connector, label))

	for i := 0; i < node.GetChildCount(); i++ {
		child := node.GetChild(i)
		newPrefix := prefix
		if isLast {
			newPrefix += "    "
		} else {
			newPrefix += "|   "
		}
		PrintVerticalTree(child, ruleNames, newPrefix, i == node.GetChildCount()-1, out)
	}
}



func GenerateASTSVG(tree antlr.Tree, ruleNames []string) string {
	fmt.Println("🔍 Iniciando generación de SVG...")
	fmt.Printf("🔍 Tree: %T, RuleNames: %d\n", tree, len(ruleNames))
	
	// Verificar que el árbol no sea nil
	if tree == nil {
		fmt.Println("❌ Tree es nil")
		return ""
	}
	
	var svg strings.Builder

	// SVG header con fondo de prueba más visible
	svg.WriteString(`<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<svg xmlns="http://www.w3.org/2000/svg" width="2000" height="1200" viewBox="0 0 2000 1200">
<defs>
<style>
.node-rect { fill: #4a4a4a; stroke: #2a2a2a; stroke-width: 1.5; rx: 5; filter: url(#shadow); }
.terminal-rect { fill: #2ecc71; stroke: #27ae60; stroke-width: 1.5; rx: 5; filter: url(#shadow); }
.node-text {
	font-family: 'Arial', sans-serif;
	font-size: 13px;
	fill: #ffffff;
	text-anchor: middle;
	dominant-baseline: central;
	paint-order: stroke;
	stroke: #000000;
	stroke-width: 0.5px;
}

.edge-line { stroke: #34495e; stroke-width: 1.5; }
.bg { fill: #f0f0f0; stroke: #cccccc; stroke-width: 2; }
</style>
<filter id="shadow" x="-50%" y="-50%" width="200%" height="200%">
  <feGaussianBlur in="SourceAlpha" stdDeviation="2"/>
  <feOffset dx="1" dy="1" result="offsetblur"/>
  <feComponentTransfer>
    <feFuncA type="linear" slope="0.5"/>
  </feComponentTransfer>
  <feMerge>
    <feMergeNode/>
    <feMergeNode in="SourceGraphic"/>
  </feMerge>
</filter>
</defs>
<!-- Fondo visible para debugging -->
<rect x="0" y="0" width="2000" height="1200" class="bg"/>
<!-- Marcador visual de prueba -->
<rect x="10" y="10" width="100" height="30" fill="red" stroke="black"/>
<text x="60" y="25" font-family="Arial" font-size="12" fill="white" text-anchor="middle">TEST</text>
`)

	nodeCounter := 0
	nodesGenerated := 0

	var generateNodes func(node antlr.Tree, x, y, level int) int
	generateNodes = func(node antlr.Tree, x, y, level int) int {
		if node == nil {
			fmt.Printf("⚠️ Nodo nil en nivel %d\n", level)
			return nodeCounter
		}
		
		fmt.Printf("🔍 Procesando nodo en (%d,%d) nivel %d: %T\n", x, y, level, node)
		
		var label string
		var isTerminal bool

		switch n := node.(type) {
		case antlr.RuleNode:
			ruleIndex := n.GetRuleContext().GetRuleIndex()
			fmt.Printf("🔍 RuleNode - RuleIndex: %d\n", ruleIndex)
			if ruleIndex >= 0 && ruleIndex < len(ruleNames) {
				label = ruleNames[ruleIndex]
			} else {
				label = fmt.Sprintf("rule_%d", ruleIndex)
			}
			isTerminal = false
		case antlr.TerminalNode:
			text := n.GetText()
			fmt.Printf("🔍 TerminalNode - Text: '%s'\n", text)
			// Solo omitir tokens específicos, no todos los símbolos
			if text == ";" || text == "(" || text == ")" || text == "{" || text == "}" {
				fmt.Printf("⏭️ Omitiendo símbolo: '%s'\n", text)
				return nodeCounter
			}
			// Escape special characters for SVG
			label = strings.ReplaceAll(text, "&", "&amp;")
			label = strings.ReplaceAll(label, "<", "&lt;")
			label = strings.ReplaceAll(label, ">", "&gt;")
			label = strings.ReplaceAll(label, "\"", "&quot;")
			isTerminal = true
		default:
			fmt.Printf("🔍 Nodo desconocido: %T\n", node)
			label = "unknown"
			isTerminal = false
		}

		if label == "" {
			label = "empty"
		}

		fmt.Printf("✅ Generando nodo: '%s' (terminal: %v)\n", label, isTerminal)

		// Dynamic width based on text length, with minimum and maximum bounds
		textLength := len(label) * 7
		width := max(60, min(textLength+20, 200))
		height := 28
		rectClass := "node-rect"
		if isTerminal {
			rectClass = "terminal-rect"
		}

		// Generar el rectángulo y texto del nodo
		svg.WriteString(fmt.Sprintf(`<rect x="%d" y="%d" width="%d" height="%d" class="%s"/>
<text x="%d" y="%d" class="node-text">%s</text>
`, x-width/2, y-height/2, width, height, rectClass, x, y, label))

		nodesGenerated++

		childCount := node.GetChildCount()
		fmt.Printf("🔍 Nodo '%s' tiene %d hijos\n", label, childCount)
		
		if childCount > 0 {
			// Dynamic spacing based on level and number of children
			childSpacing := max(120-level*15, 80) // Más espaciado
			startX := x - (childCount-1)*childSpacing/2
			childY := y + 100 + level*20 // Más separación vertical

			for i := 0; i < childCount; i++ {
				child := node.GetChild(i)
				if child == nil {
					fmt.Printf("⚠️ Hijo %d es nil\n", i)
					continue
				}
				
				childX := startX + i*childSpacing

				// Dibujar línea de conexión
				svg.WriteString(fmt.Sprintf(`<line x1="%d" y1="%d" x2="%d" y2="%d" class="edge-line"/>
`, x, y+height/2, childX, childY-height/2))

				generateNodes(child, childX, childY, level+1)
			}
		}

		nodeCounter++
		return nodeCounter
	}

	// Start from the root, centered
	fmt.Println("🔍 Iniciando generación desde la raíz...")
	generateNodes(tree, 1000, 100, 0)

	svg.WriteString("</svg>")
	
	result := svg.String()
	fmt.Printf("✅ SVG generado: %d caracteres, %d nodos\n", len(result), nodesGenerated)
	
	// Debug: mostrar inicio del SVG
	if len(result) > 200 {
		fmt.Printf("🔍 Inicio del SVG: %s...\n", result[:200])
	}
	
	// Guardar SVG para debugging (opcional)
	if err := os.WriteFile("debug_ast.svg", []byte(result), 0644); err == nil {
		fmt.Println("🔍 SVG guardado en debug_ast.svg")
	}
	
	return result
}

// Función auxiliar para verificar el árbol
func DebugTree(tree antlr.Tree, level int) {
	if tree == nil {
		return
	}
	
	indent := strings.Repeat("  ", level)
	
	switch n := tree.(type) {
	case antlr.RuleNode:
		fmt.Printf("%sRULE: %d (children: %d)\n", indent, n.GetRuleContext().GetRuleIndex(), tree.GetChildCount())
	case antlr.TerminalNode:
		fmt.Printf("%sTERM: '%s'\n", indent, n.GetText())
	default:
		fmt.Printf("%sUNKNOWN: %T (children: %d)\n", indent, tree, tree.GetChildCount())
	}
	
	for i := 0; i < tree.GetChildCount(); i++ {
		DebugTree(tree.GetChild(i), level+1)
	}
}


func GenerateTreeDOT(root antlr.Tree, parser *parser.VlangParser) string {
	var builder strings.Builder
	builder.WriteString("digraph G {\n")
	builder.WriteString("  node [shape=box, style=\"rounded,filled\", fontname=\"Arial\"];\n")
	builder.WriteString("  edge [arrowhead=none];\n")

	nodeIds := make(map[antlr.Tree]int)
	currentId := 0

	var traverse func(node antlr.Tree) int
	traverse = func(node antlr.Tree) int {
		if node == nil {
			return -1
		}
		if id, exists := nodeIds[node]; exists {
			return id
		}

		id := currentId
		currentId++
		nodeIds[node] = id

		var label string
		switch n := node.(type) {
		case *antlr.TerminalNodeImpl:
			label = fmt.Sprintf("%q", n.GetText())
			builder.WriteString(fmt.Sprintf("  %d [label=%s, fillcolor=\"#e8f5e9\"];\n", id, label))
		case antlr.RuleContext:
			label = parser.GetRuleNames()[n.GetRuleIndex()]
			builder.WriteString(fmt.Sprintf("  %d [label=\"%s\", fillcolor=\"#bbdefb\"];\n", id, label))
		}

		for i := 0; i < node.GetChildCount(); i++ {
			childId := traverse(node.GetChild(i))
			if childId != -1 {
				builder.WriteString(fmt.Sprintf("  %d -> %d;\n", id, childId))
			}
		}
		return id
	}

	traverse(root)
	builder.WriteString("}")
	return builder.String()
}
