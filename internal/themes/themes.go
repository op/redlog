package themes

import (
	"strings"

	"github.com/op/redlog/internal/logtheme"
	"github.com/op/redlog/pkg/catppuccin"
)

// Themes contains all available themes.
var Themes = []logtheme.Theme{
	catppuccin.Theme,
}

// Default theme is catppuccin.
var Default = catppuccin.Theme

// ByName returns a theme by name.
func ByName(name string) (logtheme.Theme, bool) {
	for _, t := range Themes {
		if strings.EqualFold(t.Name, name) {
			return t, true
		}
	}
	return logtheme.Theme{}, false
}

// VariantByName returns a variant of a theme by name.
func VariantByName(t logtheme.Theme, name string) (logtheme.Variant, bool) {
	for _, v := range t.Variants {
		if strings.EqualFold(v.Name, name) {
			return v, true
		}
	}
	return logtheme.Variant{}, false
}
