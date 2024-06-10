//go:generate fyne bundle -o bundled.go assets

package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type fysionTheme struct {
	fyne.Theme
}

func newFysionTheme() fyne.Theme {
	return &fysionTheme{Theme: theme.DefaultTheme()}
}

func (t *fysionTheme) Font(style fyne.TextStyle) fyne.Resource {
	if style.Symbol || style.Monospace {
		return t.Theme.Font(style)
	}

	if style.Bold && style.Italic{	
		return resourcePoppinsBoldItalicTtf
	} else if style.Bold {
		return resourcePoppinsBoldTtf
	} else if style.Italic {
		return resourcePoppinsItalicTtf
	} else {
		return resourcePoppinsRegularTtf
	}
}

func (t *fysionTheme) Color(name fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
	return t.Theme.Color(name, theme.VariantLight) // or theme.VariantDark
}

func (t *fysionTheme)  Size(name fyne.ThemeSizeName) float32 {
	if name == theme.SizeNameText {
		return 12
	}
	return t.Theme.Size(name)
}
