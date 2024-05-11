package redlog

import (
	"github.com/charmbracelet/log"

	"github.com/op/redlog/pkg/catppuccin"
)

// Catppuccin consists of three adaptive variants mixed from Catppuccin's
// beautiful pastel palettes.
var Catppuccin = struct {
	// Frappe is an less vibrant alternative using subdued colors for a muted
	// aesthetic. It is an adaptive variant using Latte in light mode and Frappe
	// elsewhere.
	Frappe *log.Styles

	// Macchiato has medium contrast with gentle colors creating a soothing
	// atmosphere. It is an adaptive variant using Latte in light mode
	// and Macchiato elsewhere.
	Macchiato *log.Styles

	// Mocha is the original and the darkest variant, offering a cozy feeling
	// with color-rich accents. It is an adaptive variant using Latte in light
	// mode and Mocha elsewhere.
	Mocha *log.Styles
}{
	Frappe:    newCatppuccin(catppuccin.Latte, catppuccin.Frappe),
	Macchiato: newCatppuccin(catppuccin.Latte, catppuccin.Macchiato),
	Mocha:     newCatppuccin(catppuccin.Latte, catppuccin.Mocha),
}

func newCatppuccin(light, dark catppuccin.Flavour) *log.Styles {
	return catppuccin.Styles(catppuccin.Adaptive(light, dark))
}
