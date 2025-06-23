package Frontend

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type DarkMonoTheme struct{}

func (DarkMonoTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
    switch n {
    case theme.ColorNameForeground:
        return color.RGBA{R: 144, G: 238, B: 144, A: 255} // Light green
    case theme.ColorNameDisabled:
        return color.RGBA{R: 144, G: 238, B: 144, A: 255} // Igual que el foreground
    case theme.ColorNameBackground:
        return color.RGBA{R: 30, G: 30, B: 30, A: 255}
    default:
        return theme.DefaultTheme().Color(n, v)
    }
}

func (DarkMonoTheme) Font(s fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(fyne.TextStyle{Monospace: true})
}

func (DarkMonoTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (DarkMonoTheme) Size(n fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(n)
}
