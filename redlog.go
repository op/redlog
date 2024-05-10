package redlog

import (
	"github.com/charmbracelet/log"
	"github.com/op/redlog/internal/themes"
)

// Default is the default theme style.
var Default = themes.Default.Default.Styles

// WithVariant prefers the specific variant of a theme.
func WithVariant(name string) func(*option) {
	return func(o *option) { o.variant = name }
}

// Theme returns the most suitable theme given the provided arguments.
//
// This function will fallback to the default theme or the found themes
// variant, if the requested was not found.
func Theme(name string, opts ...Option) *log.Styles {
	o := option{name: name}
	for _, opt := range opts {
		opt(&o)
	}

	t, ok := themes.ByName(o.name)
	if !ok {
		return Default
	}
	v, ok := themes.VariantByName(t, o.variant)
	if !ok {
		return t.Default.Styles
	}
	return v.Styles
}

type Option func(*option)

type option struct {
	name    string
	variant string
}
