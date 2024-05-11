package themes

import (
	"strings"

	"github.com/charmbracelet/log"
)

// Themes contains all available themes.
var Themes = []*Theme{
	Catppuccin,
}

// ByName returns a theme by name.
func ByName(name string) (*Theme, bool) {
	for _, t := range Themes {
		if strings.EqualFold(t.Name, name) {
			return t, true
		}
	}
	return nil, false
}

// Theme represents a theme and its variants.
type Theme struct {
	// Name is the name of the theme.
	Name string

	// Variants are the different styles that exists for a specific theme. These
	// are the standard styles and themes might allow you to customise them
	// further if you use them directly.
	Variants []*Variant
}

// New creates a New theme with the given name and variants.
func New(name string, variants []*Variant) *Theme {
	return &Theme{Name: name, Variants: variants}
}

// Variant represents a single style within a theme.
type Variant struct {
	// Name is the name of the flavour.
	Name string

	// Background is the background color of the flavour. This is only used while
	// rendering the example asset for the flavour.
	Background string

	// Styles are the actual styles that are used to render the log lines.
	Styles *log.Styles
}

// VariantByName returns a variant of a theme by name.
func (t *Theme) Variant(name string) (*Variant, bool) {
	for _, v := range t.Variants {
		if strings.EqualFold(v.Name, name) {
			return v, true
		}
	}
	return nil, false
}
