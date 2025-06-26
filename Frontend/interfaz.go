package Frontend

import (
	"bytes"

	"fmt"
	"os/exec"

	"os"
	"strings"

	"compiler/arm"
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

	codigoARM := widget.NewMultiLineEntry()
	codigoARM.Disable()
	stylizeEntry(codigoARM)

	codigoARMScroll := container.NewScroll(codigoARM)
	codigoARMScroll.SetMinSize(fyne.NewSize(400, 300))

	codigoARMBox := container.NewBorder(
		styledLabel("Código ARM generado:", false),
		nil, nil, nil,
		codigoARMScroll,
	)

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
		if !strings.Contains(codigo, "fun main") {
			codigo = "fun main() {\n" + codigo + "\n}"
		}


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


		armVisitor := arm.NewArmVisitor()
		armVisitor.Visit(tree)

		armFullCode := armVisitor.Generator.GetFullCode()

		fmt.Println("=== ARM generado ===")
		fmt.Println(armFullCode)
		codigoARM.SetText(armFullCode)



	})


	btnReportes := widget.NewButton("Ver Reportes", func() {
		ventanaReportes := a.NewWindow("Reportes")
		ventanaReportes.Resize(fyne.NewSize(900, 400))

		reportScroll := container.NewScroll(reportTabs)
		ventanaReportes.SetContent(reportScroll)
		ventanaReportes.Show()
	})


	btnExportarARM := widget.NewButton("Exportar .s", func() {
		saveDialog := dialog.NewFileSave(func(wr fyne.URIWriteCloser, err error) {
			if err == nil && wr != nil {
				wr.Write([]byte(codigoARM.Text))
				wr.Close()
			}
		}, w)
		saveDialog.SetFileName("salida.s")
		saveDialog.Show()
	})


	toolbar := container.NewHBox(btnNuevo, btnAbrir, btnGuardar, btnEjecutar, btnReportes, btnExportarARM)

	horizontal := container.NewHSplit(
		entradaBox,
		container.NewVSplit(consolaBox, codigoARMBox),
	)




	
	reportScroll := container.NewScroll(reportTabs)
	reportScroll.SetMinSize(fyne.NewSize(800, 250))

	w.SetContent(container.NewBorder(toolbar, nil, nil, nil, horizontal))




	w.ShowAndRun()
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
