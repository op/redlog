package redlog

import (
	"testing"

	"github.com/op/redlog/internal/logtheme"
	"github.com/op/redlog/internal/themes"
)

func TestTheme(t *testing.T) {
	var (
		// catppuccin
		catp          = theme(t, themes.Themes[0], "catppuccin")
		catpLatte     = variant(t, catp.Variants[0], "latte")
		catpFrappe    = variant(t, catp.Variants[1], "frappe")
		catpMacchiato = variant(t, catp.Variants[2], "macchiato")
		catpMocha     = variant(t, catp.Variants[3], "mocha")
	)

	tests := []struct {
		theme   string
		variant string
		want    logtheme.Variant
		wantErr error
	}{
		{"", "", catp.Default, ErrUnknownTheme},

		// catppuccin
		{"catppuccin", "", catp.Default, ErrUnknownVariant},
		{"catppuccin", "latte", catpLatte, nil},
		{"catppuccin", "frappe", catpFrappe, nil},
		{"catppuccin", "macchiato", catpMacchiato, nil},
		{"catppuccin", "mocha", catpMocha, nil},

		// casing
		{"catppuccin", "LATTE", catpLatte, nil},
		{"CaTpPuCcIN", "lAtTe", catpLatte, nil},
		{"CATPPUCCIN", "latte", catpLatte, nil},

		// uknown
		{"CATPPUCCIN", "unknown", catp.Default, ErrUnknownVariant},
		{"dogppuccin", "", catp.Default, ErrUnknownTheme},
	}

	for _, tt := range tests {
		t.Run(tt.theme+"/"+tt.variant, func(t *testing.T) {
			var opts []Option
			if tt.variant != "" {
				opts = append(opts, WithVariant(tt.variant))
			}

			got, err := Theme(tt.theme, opts...)
			if err != tt.wantErr {
				t.Fatalf("got = %v, want = %v", err, tt.wantErr)
			} else if got != tt.want.Styles {
				t.Errorf("got = %v, want = %v", got, tt.want.Styles)
			}
		})
	}
}

func theme(t *testing.T, th logtheme.Theme, name string) logtheme.Theme {
	t.Helper()
	if th.Name != name {
		t.Fatalf("got = %s, want = %s", th.Name, name)
	}
	return th
}

func variant(t *testing.T, v logtheme.Variant, name string) logtheme.Variant {
	t.Helper()
	if v.Name != name {
		t.Fatalf("got = %s, want = %s", v.Name, name)
	}
	return v
}
