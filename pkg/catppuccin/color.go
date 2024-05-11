package catppuccin

import (
	"fmt"

	catppuccin "github.com/catppuccin/go"
	"github.com/charmbracelet/lipgloss"
)

// Flavour is an interface implemented by all Catppuccin flavours.
//
// This is the catppuccin.Flavour interface but with lipgloss.TerminalColor.
type Flavour interface {
	Rosewater() lipgloss.TerminalColor
	Flamingo() lipgloss.TerminalColor
	Pink() lipgloss.TerminalColor
	Mauve() lipgloss.TerminalColor
	Red() lipgloss.TerminalColor
	Maroon() lipgloss.TerminalColor
	Peach() lipgloss.TerminalColor
	Yellow() lipgloss.TerminalColor
	Green() lipgloss.TerminalColor
	Teal() lipgloss.TerminalColor
	Sky() lipgloss.TerminalColor
	Sapphire() lipgloss.TerminalColor
	Blue() lipgloss.TerminalColor
	Lavender() lipgloss.TerminalColor
	Text() lipgloss.TerminalColor
	Subtext1() lipgloss.TerminalColor
	Subtext0() lipgloss.TerminalColor
	Overlay2() lipgloss.TerminalColor
	Overlay1() lipgloss.TerminalColor
	Overlay0() lipgloss.TerminalColor
	Surface2() lipgloss.TerminalColor
	Surface1() lipgloss.TerminalColor
	Surface0() lipgloss.TerminalColor
	Crust() lipgloss.TerminalColor
	Mantle() lipgloss.TerminalColor
	Base() lipgloss.TerminalColor
	Name() string
}

// flavour implements the Flavour interface for a single flavour.
type flavour struct {
	c catppuccin.Flavour
}

// FromFlavour converts a catppuccin flavour to the Variant in this package.
func FromFlavour(c catppuccin.Flavour) Flavour {
	return &flavour{c}
}

// color converts a catppuccin color to a lipgloss color.
func color(c catppuccin.Color) lipgloss.TerminalColor {
	// TODO: use CompleteColor
	return lipgloss.Color(c.Hex)
}

func (v flavour) Rosewater() lipgloss.TerminalColor {
	return color(v.c.Rosewater())
}

func (v flavour) Flamingo() lipgloss.TerminalColor {
	return color(v.c.Flamingo())
}

func (v flavour) Pink() lipgloss.TerminalColor {
	return color(v.c.Pink())
}

func (v flavour) Mauve() lipgloss.TerminalColor {
	return color(v.c.Mauve())
}

func (v flavour) Red() lipgloss.TerminalColor {
	return color(v.c.Red())
}

func (v flavour) Maroon() lipgloss.TerminalColor {
	return color(v.c.Maroon())
}

func (v flavour) Peach() lipgloss.TerminalColor {
	return color(v.c.Peach())
}

func (v flavour) Yellow() lipgloss.TerminalColor {
	return color(v.c.Yellow())
}

func (v flavour) Green() lipgloss.TerminalColor {
	return color(v.c.Green())
}

func (v flavour) Teal() lipgloss.TerminalColor {
	return color(v.c.Teal())
}

func (v flavour) Sky() lipgloss.TerminalColor {
	return color(v.c.Sky())
}

func (v flavour) Sapphire() lipgloss.TerminalColor {
	return color(v.c.Sapphire())
}

func (v flavour) Blue() lipgloss.TerminalColor {
	return color(v.c.Blue())
}

func (v flavour) Lavender() lipgloss.TerminalColor {
	return color(v.c.Lavender())
}

func (v flavour) Text() lipgloss.TerminalColor {
	return color(v.c.Text())
}

func (v flavour) Subtext1() lipgloss.TerminalColor {
	return color(v.c.Subtext1())
}

func (v flavour) Subtext0() lipgloss.TerminalColor {
	return color(v.c.Subtext0())
}

func (v flavour) Overlay2() lipgloss.TerminalColor {
	return color(v.c.Overlay2())
}

func (v flavour) Overlay1() lipgloss.TerminalColor {
	return color(v.c.Overlay1())
}

func (v flavour) Overlay0() lipgloss.TerminalColor {
	return color(v.c.Overlay0())
}

func (v flavour) Surface2() lipgloss.TerminalColor {
	return color(v.c.Surface2())
}

func (v flavour) Surface1() lipgloss.TerminalColor {
	return color(v.c.Surface1())
}

func (v flavour) Surface0() lipgloss.TerminalColor {
	return color(v.c.Surface0())
}

func (v flavour) Crust() lipgloss.TerminalColor {
	return color(v.c.Crust())
}

func (v flavour) Mantle() lipgloss.TerminalColor {
	return color(v.c.Mantle())
}

func (v flavour) Base() lipgloss.TerminalColor {
	return color(v.c.Base())
}

func (v flavour) Name() string {
	return v.c.Name()
}

// adaptive implements the Flavour interface using two flavours, one
// light and one dark, and adapts to the terminal background color.
type adaptive struct {
	light    Flavour
	dark     Flavour
	renderer *lipgloss.Renderer
}

type AdaptiveOption func(a *adaptive)

func WithRenderer(r *lipgloss.Renderer) AdaptiveOption {
	return func(a *adaptive) { a.renderer = r }
}

// Adaptive creates an adaptive variant based on the provided flavours.
//
// The Catppuccin theme comes with one light and three dark flavours. Use this
// function to create your own adaptive variants.
//
// Latte has the lightest palette and is a light flavour, followed by Frappe,
// Macchiato, and Mocha -- which all are dark flavours.
func Adaptive(light, dark Flavour, opts ...AdaptiveOption) Flavour {
	f := &adaptive{light: light, dark: dark}

	for _, opt := range opts {
		opt(f)
	}

	if f.renderer == nil {
		f.renderer = lipgloss.DefaultRenderer()
	}

	return f
}

func (v adaptive) color(light, dark lipgloss.TerminalColor) lipgloss.TerminalColor {
	if v.renderer.HasDarkBackground() {
		return dark
	}
	return light
}

func (v adaptive) Rosewater() lipgloss.TerminalColor {
	return v.color(v.light.Rosewater(), v.dark.Rosewater())
}

func (v adaptive) Flamingo() lipgloss.TerminalColor {
	return v.color(v.light.Flamingo(), v.dark.Flamingo())
}

func (v adaptive) Pink() lipgloss.TerminalColor {
	return v.color(v.light.Pink(), v.dark.Pink())
}

func (v adaptive) Mauve() lipgloss.TerminalColor {
	return v.color(v.light.Mauve(), v.dark.Mauve())
}

func (v adaptive) Red() lipgloss.TerminalColor {
	return v.color(v.light.Red(), v.dark.Red())
}

func (v adaptive) Maroon() lipgloss.TerminalColor {
	return v.color(v.light.Maroon(), v.dark.Maroon())
}

func (v adaptive) Peach() lipgloss.TerminalColor {
	return v.color(v.light.Peach(), v.dark.Peach())
}

func (v adaptive) Yellow() lipgloss.TerminalColor {
	return v.color(v.light.Yellow(), v.dark.Yellow())
}

func (v adaptive) Green() lipgloss.TerminalColor {
	return v.color(v.light.Green(), v.dark.Green())
}

func (v adaptive) Teal() lipgloss.TerminalColor {
	return v.color(v.light.Teal(), v.dark.Teal())
}

func (v adaptive) Sky() lipgloss.TerminalColor {
	return v.color(v.light.Sky(), v.dark.Sky())
}

func (v adaptive) Sapphire() lipgloss.TerminalColor {
	return v.color(v.light.Sapphire(), v.dark.Sapphire())
}

func (v adaptive) Blue() lipgloss.TerminalColor {
	return v.color(v.light.Blue(), v.dark.Blue())
}

func (v adaptive) Lavender() lipgloss.TerminalColor {
	return v.color(v.light.Lavender(), v.dark.Lavender())
}

func (v adaptive) Text() lipgloss.TerminalColor {
	return v.color(v.light.Text(), v.dark.Text())
}

func (v adaptive) Subtext1() lipgloss.TerminalColor {
	return v.color(v.light.Subtext1(), v.dark.Subtext1())
}

func (v adaptive) Subtext0() lipgloss.TerminalColor {
	return v.color(v.light.Subtext0(), v.dark.Subtext0())
}

func (v adaptive) Overlay2() lipgloss.TerminalColor {
	return v.color(v.light.Overlay2(), v.dark.Overlay2())
}

func (v adaptive) Overlay1() lipgloss.TerminalColor {
	return v.color(v.light.Overlay1(), v.dark.Overlay1())
}

func (v adaptive) Overlay0() lipgloss.TerminalColor {
	return v.color(v.light.Overlay0(), v.dark.Overlay0())
}

func (v adaptive) Surface2() lipgloss.TerminalColor {
	return v.color(v.light.Surface2(), v.dark.Surface2())
}

func (v adaptive) Surface1() lipgloss.TerminalColor {
	return v.color(v.light.Surface1(), v.dark.Surface1())
}

func (v adaptive) Surface0() lipgloss.TerminalColor {
	return v.color(v.light.Surface0(), v.dark.Surface0())
}

func (v adaptive) Crust() lipgloss.TerminalColor {
	return v.color(v.light.Crust(), v.dark.Crust())
}

func (v adaptive) Mantle() lipgloss.TerminalColor {
	return v.color(v.light.Mantle(), v.dark.Mantle())
}

func (v adaptive) Base() lipgloss.TerminalColor {
	return v.color(v.light.Base(), v.dark.Base())
}

func (v adaptive) Name() string {
	return fmt.Sprintf("%s/%s", v.light.Name(), v.dark.Name())
}
