module github.com/op/logstyles

go 1.22.1

replace (
	github.com/op/logstyles/internal/logtheme => ./internal/logtheme
	github.com/op/logstyles/internal/themes => ./internal/themes
	github.com/op/logstyles/pkg/catppuccin => ./pkg/catppuccin
)

require (
	github.com/charmbracelet/log v0.4.0
	github.com/op/logstyles/internal/logtheme v0.0.0-00010101000000-000000000000
	github.com/op/logstyles/internal/themes v0.0.0-00010101000000-000000000000
)

require (
	github.com/aymanbagabas/go-osc52/v2 v2.0.1 // indirect
	github.com/catppuccin/go v0.2.0 // indirect
	github.com/charmbracelet/lipgloss v0.10.0 // indirect
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.18 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/muesli/reflow v0.3.0 // indirect
	github.com/muesli/termenv v0.15.2 // indirect
	github.com/op/logstyles/pkg/catppuccin v0.0.0-00010101000000-000000000000 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	golang.org/x/exp v0.0.0-20231006140011-7918f672742d // indirect
	golang.org/x/sys v0.13.0 // indirect
)
