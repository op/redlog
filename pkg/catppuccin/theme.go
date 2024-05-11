package catppuccin

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"

	"github.com/op/redlog/internal/logtheme"
)

// Theme consists of 4 beautiful pastel color palettes. The theme comes in one
// light and three dark variants.
var Theme = logtheme.New(
	"catppuccin",
	newVariant(Adaptive(Latte, Mocha)),
	[]logtheme.Variant{
		newVariant(Latte),
		newVariant(Frappe),
		newVariant(Macchiato),
		newVariant(Mocha),
	},
)

// Default is the default Catppuccin variant which is an adaptive flavour using
// Latte for light backgrounds and Mocha for dark backgrounds.
var Default = Theme.Default

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

func newVariant(f Variant) logtheme.Variant {
	return logtheme.Variant{
		Name:       f.Name(),
		Background: hexColor(f.Base()),
		Styles:     New(f),
	}
}
