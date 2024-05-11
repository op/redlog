package themes

import (
	"testing"
)

func TestByName(t *testing.T) {
	tests := []struct {
		theme string
		want  *Theme
	}{
		// casing
		{"catppuccin", Catppuccin},
		{"CaTpPuCcIN", Catppuccin},
		{"CATPPUCCIN", Catppuccin},

		// uknown
		{"dogppuccin", nil},
	}

	for _, tt := range tests {
		t.Run(tt.theme, func(t *testing.T) {
			got, ok := ByName(tt.theme)
			if !ok && tt.want != nil {
				t.Errorf("got = %v, want = %v", ok, true)
			} else if got != tt.want {
				t.Fatalf("got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestVariant(t *testing.T) {
	var (
		// catppuccin
		catpLatte     = variant(t, Catppuccin.Variants[0], "latte")
		catpFrappe    = variant(t, Catppuccin.Variants[1], "frappe")
		catpMacchiato = variant(t, Catppuccin.Variants[2], "macchiato")
		catpMocha     = variant(t, Catppuccin.Variants[3], "mocha")
	)

	tests := []struct {
		theme   *Theme
		variant string
		want    *Variant
	}{
		// catppuccin
		{Catppuccin, "latte", catpLatte},
		{Catppuccin, "frappe", catpFrappe},
		{Catppuccin, "macchiato", catpMacchiato},
		{Catppuccin, "mocha", catpMocha},

		// casing
		{Catppuccin, "LATTE", catpLatte},
		{Catppuccin, "lAtTe", catpLatte},
		{Catppuccin, "latte", catpLatte},

		// uknown
		{Catppuccin, "unknown", nil},
	}

	for _, tt := range tests {
		t.Run(tt.variant, func(t *testing.T) {
			got, ok := tt.theme.Variant(tt.variant)
			if !ok && tt.want != nil {
				t.Errorf("got = %v, want = %v", ok, true)
			} else if got != tt.want {
				t.Fatalf("got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func variant(t *testing.T, v *Variant, name string) *Variant {
	t.Helper()
	if v.Name != name {
		t.Fatalf("got = %s, want = %s", v.Name, name)
	}
	return v
}
