package catppuccin

import (
	"fmt"

	catppuccin "github.com/catppuccin/go"
	"github.com/charmbracelet/lipgloss"
)

// Variant is an interface implemented by all Catppuccin variations.
//
// This is the catppuccin.Flavour interface but with lipgloss.TerminalColor.
type Variant interface {
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

// variant implements the Flavour interface for a single variant.
type variant struct {
	c catppuccin.Flavour
}

// FromFlavour converts a catppuccin flavour to the Variant in this package.
func FromFlavour(c catppuccin.Flavour) Variant {
	return &variant{c}
}

// color converts a catppuccin color to a lipgloss color.
func color(c catppuccin.Color) lipgloss.TerminalColor {
	// TODO: use CompleteColor
	return lipgloss.Color(c.Hex)
}

func (v variant) Rosewater() lipgloss.TerminalColor {
	return color(v.c.Rosewater())
}

func (v variant) Flamingo() lipgloss.TerminalColor {
	return color(v.c.Flamingo())
}

func (v variant) Pink() lipgloss.TerminalColor {
	return color(v.c.Pink())
}

func (v variant) Mauve() lipgloss.TerminalColor {
	return color(v.c.Mauve())
}

func (v variant) Red() lipgloss.TerminalColor {
	return color(v.c.Red())
}

func (v variant) Maroon() lipgloss.TerminalColor {
	return color(v.c.Maroon())
}

func (v variant) Peach() lipgloss.TerminalColor {
	return color(v.c.Peach())
}

func (v variant) Yellow() lipgloss.TerminalColor {
	return color(v.c.Yellow())
}

func (v variant) Green() lipgloss.TerminalColor {
	return color(v.c.Green())
}

func (v variant) Teal() lipgloss.TerminalColor {
	return color(v.c.Teal())
}

func (v variant) Sky() lipgloss.TerminalColor {
	return color(v.c.Sky())
}

func (v variant) Sapphire() lipgloss.TerminalColor {
	return color(v.c.Sapphire())
}

func (v variant) Blue() lipgloss.TerminalColor {
	return color(v.c.Blue())
}

func (v variant) Lavender() lipgloss.TerminalColor {
	return color(v.c.Lavender())
}

func (v variant) Text() lipgloss.TerminalColor {
	return color(v.c.Text())
}

func (v variant) Subtext1() lipgloss.TerminalColor {
	return color(v.c.Subtext1())
}

func (v variant) Subtext0() lipgloss.TerminalColor {
	return color(v.c.Subtext0())
}

func (v variant) Overlay2() lipgloss.TerminalColor {
	return color(v.c.Overlay2())
}

func (v variant) Overlay1() lipgloss.TerminalColor {
	return color(v.c.Overlay1())
}

func (v variant) Overlay0() lipgloss.TerminalColor {
	return color(v.c.Overlay0())
}

func (v variant) Surface2() lipgloss.TerminalColor {
	return color(v.c.Surface2())
}

func (v variant) Surface1() lipgloss.TerminalColor {
	return color(v.c.Surface1())
}

func (v variant) Surface0() lipgloss.TerminalColor {
	return color(v.c.Surface0())
}

func (v variant) Crust() lipgloss.TerminalColor {
	return color(v.c.Crust())
}

func (v variant) Mantle() lipgloss.TerminalColor {
	return color(v.c.Mantle())
}

func (v variant) Base() lipgloss.TerminalColor {
	return color(v.c.Base())
}

func (v variant) Name() string {
	return v.c.Name()
}

// adaptiveVariant implements the Variant interface using two flavours, one
// light and one dark, and adapts to the terminal background color.
type adaptiveVariant struct {
	light    Variant
	dark     Variant
	renderer *lipgloss.Renderer
}

type VariantOption func(f *adaptiveVariant)

func WithRenderer(r *lipgloss.Renderer) VariantOption {
	return func(f *adaptiveVariant) { f.renderer = r }
}

// Adaptive creates an adaptive variant based on the provided variants.
//
// The Catppuccin theme comes with one light and three dark variants. Use this
// function to create your own adaptive variants.
//
// Latte has the lightest palette and is a light variant, followed by Frappe,
// Macchiato, and Mocha -- which all are dark variants.
func Adaptive(light, dark Variant, opts ...VariantOption) Variant {
	f := &adaptiveVariant{light: light, dark: dark}

	for _, opt := range opts {
		opt(f)
	}

	if f.renderer == nil {
		f.renderer = lipgloss.DefaultRenderer()
	}

	return f
}

func (v adaptiveVariant) color(light, dark lipgloss.TerminalColor) lipgloss.TerminalColor {
	if v.renderer.HasDarkBackground() {
		return dark
	}
	return light
}

func (v adaptiveVariant) Rosewater() lipgloss.TerminalColor {
	return v.color(v.light.Rosewater(), v.dark.Rosewater())
}

func (v adaptiveVariant) Flamingo() lipgloss.TerminalColor {
	return v.color(v.light.Flamingo(), v.dark.Flamingo())
}

func (v adaptiveVariant) Pink() lipgloss.TerminalColor {
	return v.color(v.light.Pink(), v.dark.Pink())
}

func (v adaptiveVariant) Mauve() lipgloss.TerminalColor {
	return v.color(v.light.Mauve(), v.dark.Mauve())
}

func (v adaptiveVariant) Red() lipgloss.TerminalColor {
	return v.color(v.light.Red(), v.dark.Red())
}

func (v adaptiveVariant) Maroon() lipgloss.TerminalColor {
	return v.color(v.light.Maroon(), v.dark.Maroon())
}

func (v adaptiveVariant) Peach() lipgloss.TerminalColor {
	return v.color(v.light.Peach(), v.dark.Peach())
}

func (v adaptiveVariant) Yellow() lipgloss.TerminalColor {
	return v.color(v.light.Yellow(), v.dark.Yellow())
}

func (v adaptiveVariant) Green() lipgloss.TerminalColor {
	return v.color(v.light.Green(), v.dark.Green())
}

func (v adaptiveVariant) Teal() lipgloss.TerminalColor {
	return v.color(v.light.Teal(), v.dark.Teal())
}

func (v adaptiveVariant) Sky() lipgloss.TerminalColor {
	return v.color(v.light.Sky(), v.dark.Sky())
}

func (v adaptiveVariant) Sapphire() lipgloss.TerminalColor {
	return v.color(v.light.Sapphire(), v.dark.Sapphire())
}

func (v adaptiveVariant) Blue() lipgloss.TerminalColor {
	return v.color(v.light.Blue(), v.dark.Blue())
}

func (v adaptiveVariant) Lavender() lipgloss.TerminalColor {
	return v.color(v.light.Lavender(), v.dark.Lavender())
}

func (v adaptiveVariant) Text() lipgloss.TerminalColor {
	return v.color(v.light.Text(), v.dark.Text())
}

func (v adaptiveVariant) Subtext1() lipgloss.TerminalColor {
	return v.color(v.light.Subtext1(), v.dark.Subtext1())
}

func (v adaptiveVariant) Subtext0() lipgloss.TerminalColor {
	return v.color(v.light.Subtext0(), v.dark.Subtext0())
}

func (v adaptiveVariant) Overlay2() lipgloss.TerminalColor {
	return v.color(v.light.Overlay2(), v.dark.Overlay2())
}

func (v adaptiveVariant) Overlay1() lipgloss.TerminalColor {
	return v.color(v.light.Overlay1(), v.dark.Overlay1())
}

func (v adaptiveVariant) Overlay0() lipgloss.TerminalColor {
	return v.color(v.light.Overlay0(), v.dark.Overlay0())
}

func (v adaptiveVariant) Surface2() lipgloss.TerminalColor {
	return v.color(v.light.Surface2(), v.dark.Surface2())
}

func (v adaptiveVariant) Surface1() lipgloss.TerminalColor {
	return v.color(v.light.Surface1(), v.dark.Surface1())
}

func (v adaptiveVariant) Surface0() lipgloss.TerminalColor {
	return v.color(v.light.Surface0(), v.dark.Surface0())
}

func (v adaptiveVariant) Crust() lipgloss.TerminalColor {
	return v.color(v.light.Crust(), v.dark.Crust())
}

func (v adaptiveVariant) Mantle() lipgloss.TerminalColor {
	return v.color(v.light.Mantle(), v.dark.Mantle())
}

func (v adaptiveVariant) Base() lipgloss.TerminalColor {
	return v.color(v.light.Base(), v.dark.Base())
}

func (v adaptiveVariant) Name() string {
	return fmt.Sprintf("%s/%s", v.light.Name(), v.dark.Name())
}
