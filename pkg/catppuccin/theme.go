package catppuccin

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"

	"github.com/op/redlog/internal/logtheme"
)

// Theme consists of 4 beautiful pastel color palettes. The theme comes in one
// light and three dark variants.
var Theme = logtheme.New(
	"Catppuccin",
	newVariant(Adaptive(Latte, Mocha)),
	[]logtheme.Variant{
		newVariant(Latte),
		newVariant(Frappe),
		newVariant(Macchiato),
		newVariant(Mocha),
	},
)

func hexColor(c lipgloss.TerminalColor) string {
	if c, ok := c.(lipgloss.Color); ok {
		return string(c)
	}
	// TODO: fix
	r, g, b, _ := c.RGBA()
	return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}

func newVariant(f Variant) logtheme.Variant {
	return logtheme.Variant{
		Name:       f.Name(),
		Background: hexColor(f.Base()),
		Styles:     New(f),
	}
}
