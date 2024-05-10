package themes

import (
	"strings"

	"github.com/op/redlog/internal/logtheme"
	"github.com/op/redlog/pkg/catppuccin"
)

var Themes = []logtheme.Theme{
	catppuccin.Theme,
}

var Default = catppuccin.Theme

func ByName(name string) (logtheme.Theme, bool) {
	for _, t := range Themes {
		if strings.EqualFold(t.Name, name) {
			return t, true
		}
	}
	return logtheme.Theme{}, false
}

func VariantByName(t logtheme.Theme, name string) (logtheme.Variant, bool) {
	for _, v := range t.Variants {
		if strings.EqualFold(v.Name, name) {
			return v, true
		}
	}
	return logtheme.Variant{}, false
}
