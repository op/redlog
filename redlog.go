package redlog

import (
	"github.com/charmbracelet/log"

	"github.com/op/redlog/pkg/catppuccin"
)

// Catppuccin consists of three adaptive variants mixed from Catppuccin's
// beautiful pastel palettes.
var Catppuccin = struct {
	Frappe    *log.Styles
	Macchiato *log.Styles
	Mocha     *log.Styles
}{
	Frappe:    newCatppuccin(catppuccin.Latte, catppuccin.Frappe),
	Macchiato: newCatppuccin(catppuccin.Latte, catppuccin.Macchiato),
	Mocha:     newCatppuccin(catppuccin.Latte, catppuccin.Mocha),
}

func newCatppuccin(light, dark catppuccin.Flavour) *log.Styles {
	return catppuccin.Styles(catppuccin.Adaptive(light, dark))
}
