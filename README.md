# Red Log

<p>
    <a href="https://github.com/op/redlog/releases"><img src="https://img.shields.io/github/release/op/redlog.svg" alt="Latest Release"></a>
    <a href="https://pkg.go.dev/github.com/op/redlog?tab=doc"><img src="https://godoc.org/github.com/golang/gddo?status.svg" alt="Go Docs"></a>
    <a href="https://github.com/op/redlog/actions"><img src="https://github.com/op/redlog/workflows/build/badge.svg" alt="Build Status"></a>
    <a href="https://goreportcard.com/report/github.com/op/redlog"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/op/redlog"></a>
</p>

Gloss styles for [logs][log]. 🪵

It contains a set of styles that can be used with the [log] package. The
following themes are provided:

* [catppucchin](https://catppuccin.com)

[log]: /charmbracelet/log

## Gallery

*Catppuccin Latte*

<picture>
    <img width="500" src="./assets/catppuccin-latte.svg" alt="catppuccin latte" />
</picture>

*Catppuccin Frappé*

<picture>
    <img width="500" src="./assets/catppuccin-frappe.svg" alt="catppuccin frappe" />
</picture>

*Catppuccin Macchiato*

<picture>
    <img width="500" src="./assets/catppuccin-macchiato.svg" alt="catppuccin macchiato" />
</picture>

*Catppuccin Mocha*

<picture>
    <img width="500" src="./assets/catppuccin-mocha.svg" alt="catppuccin mocha" />
</picture>

## Usage

Use `go get` to download the dependency.

```bash
go get github.com/op/redlog@latest
```

Then, `import` it in Go files:

```go
import (
  "github.com/charmbracelet/log"
  "github.com/op/redlog"
)
```

The Red Log package comes with a global `Default` style.

```go
log.SetStyles(redlog.Default)
slog.SetDefault(slog.New(log.Default()))
```

It is possible to request a specific theme.

```go
styles, _ := redlog.Theme("catppuccin")
logger := log.New(os.Stderr)
logger.SetStyles(styles)
```

And with a specific variation of the theme.

```go
styles, err := redlog.Theme("catppuccin", redlog.WithVariant("mocha"))
if err != nil {
    log.Fatal(err)
}
log.SetStyles(styles)
```

## Specific theme

If you want to limit the number of dependencies, or make some customisation of
the theme, you can directly import and use the theme.

```go
import (
  "github.com/charmbracelet/log"
  "github.com/op/redlog/pkg/catppuccin"
)
```

Then initiate the theme the way that suits your needs.

```go
// use an adaptive theme that switches between light and dark depending on the
// terminal's color scheme
variant := catppuccin.Adaptive(catppuccin.Latte, catppuccin.Mocha)
log.SetStyles(catppuccin.New(variant))
```

## Contributing

Initialise `go.work` with `make bootstrap`.

```bash
make bootstrap
make
```
