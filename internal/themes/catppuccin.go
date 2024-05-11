package themes

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/op/redlog/pkg/catppuccin"
)

// Catppuccin consists of 4 beautiful pastel color palettes, named flavors.
var Catppuccin = New(
	"catppuccin",
	[]*Variant{
		catppucinVariant(catppuccin.Latte),
		catppucinVariant(catppuccin.Frappe),
		catppucinVariant(catppuccin.Macchiato),
		catppucinVariant(catppuccin.Mocha),
	},
)

func catppucinVariant(f catppuccin.Flavour) *Variant {
	return &Variant{
		Name:       f.Name(),
		Background: hexColor(f.Base()),
		Styles:     catppuccin.Styles(f),
	}
}

func hexColor(c lipgloss.TerminalColor) string {
	if c, ok := c.(lipgloss.Color); ok {
		hex := string(c)
		if len(hex) == 7 {
			return hex
		}
	}
	r, g, b, _ := c.RGBA()
	return fmt.Sprintf("#%02x%02x%02x", r/0xff, g/0xff, b/0xff)
}
