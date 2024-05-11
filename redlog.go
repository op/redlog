package redlog

import (
	"errors"

	"github.com/charmbracelet/log"
	"github.com/op/redlog/internal/themes"
)

var (
	// ErrUnknownTheme is returned when a theme could not found.
	ErrUnknownTheme = errors.New("unknown theme")
	// ErrUnknownVariant is returned when a theme variant could not found.
	ErrUnknownVariant = errors.New("unknown variant")
)

// Default is the default theme style.
var Default = themes.Default.Default.Styles

// WithVariant prefers the specific variant of a theme.
func WithVariant(name string) func(*option) {
	return func(o *option) { o.variant = name }
}

// Theme returns the most suitable theme given the provided arguments.
//
// This function will always return a usable style, together with an error if
// the requested theme or variant was not found.
func Theme(name string, opts ...Option) (*log.Styles, error) {
	o := option{name: name}
	for _, opt := range opts {
		opt(&o)
	}

	t, ok := themes.ByName(o.name)
	if !ok {
		return Default, ErrUnknownTheme
	}
	v, ok := themes.VariantByName(t, o.variant)
	if !ok {
		return t.Default.Styles, ErrUnknownVariant
	}
	return v.Styles, nil
}

type Option func(*option)

type option struct {
	name    string
	variant string
}
