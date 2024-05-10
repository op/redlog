package logtheme

import "github.com/charmbracelet/log"

// Theme represents a theme and its variants.
type Theme struct {
	// Name is the name of the theme.
	Name string

	// Default is the default variant of the theme.
	Default Variant

	// Variants are the different styles that exists for a specific theme. These
	// are the standard styles and themes might allow you to customise them
	// further if you use them directly.
	Variants []Variant
}

// New creates a new theme with the given name and variants.
func New(name string, def Variant, variants []Variant) Theme {
	return Theme{Name: name, Default: def, Variants: variants}
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
