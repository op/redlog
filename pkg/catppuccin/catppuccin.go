// Package catppuccin consists of 4 beautiful pastel color palettes, named
// flavors. The theme comes in one light and three dark variants.
package catppuccin

import (
	catppuccin "github.com/catppuccin/go"
	"github.com/charmbracelet/log"
)

var (
	// Latte is the lightest theme harmoniously inverting the essence of
	// Catppuccin's dark themes.
	Latte = FromFlavour(catppuccin.Latte)
	// Frappe is a less vibrant alternative using subdued colors for a muted
	// aesthetic.
	Frappe = FromFlavour(catppuccin.Frappe)
	// Macchiato has medium contrast with gentle colors creating a soothing
	// atmosphere.
	Macchiato = FromFlavour(catppuccin.Macchiato)
	// Mocha is the original and the darkest variant, offering a cozy feeling
	// with color-rich accents.
	Mocha = FromFlavour(catppuccin.Mocha)
)

// New creates new log styles based the provided flavor of Catppuccin.
func New(f Variant) *log.Styles {
	d := log.DefaultStyles()

	// message metadata
	d.Timestamp = d.Timestamp.Foreground(f.Overlay1()) // subtle
	d.Caller = d.Caller.Foreground(f.Overlay1())       // line number
	d.Prefix = d.Prefix.Foreground(f.Mauve())          // keyword / syntax
	d.Message = d.Message.Foreground(f.Text())         // body copy

	// structured logging
	d.Key = d.Key.Foreground(f.Subtext0())             // labels
	d.Value = d.Value.Foreground(f.Maroon())           // parameters
	d.Separator = d.Separator.Foreground(f.Overlay2()) // braces, delimiters

	// log levels uses rainbow highlight / errors, warnings and informative
	d.Levels[log.DebugLevel] = d.Levels[log.DebugLevel].Foreground(f.Sapphire())
	d.Levels[log.InfoLevel] = d.Levels[log.InfoLevel].Foreground(f.Green())
	d.Levels[log.WarnLevel] = d.Levels[log.WarnLevel].Foreground(f.Yellow())
	d.Levels[log.ErrorLevel] = d.Levels[log.ErrorLevel].Foreground(f.Peach())
	d.Levels[log.FatalLevel] = d.Levels[log.FatalLevel].Foreground(f.Red())

	return d
}
