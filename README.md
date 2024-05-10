# Red Log

<p>
    <a href="https://github.com/op/redlog/releases"><img src="https://img.shields.io/github/release/op/redlog.svg" alt="Latest Release"></a>
    <a href="https://pkg.go.dev/github.com/op/redlog?tab=doc"><img src="https://godoc.org/github.com/golang/gddo?status.svg" alt="Go Docs"></a>
    <a href="https://github.com/op/redlog/actions"><img src="https://github.com/op/redlog/workflows/build/badge.svg" alt="Build Status"></a>
</p>

Gloss styles for [logs][log]. 🪵

It contains a set of styles that can be used with the [log] package. The
following themes are provided:

* [catppucchin](https://catppuccin.com)

[log]: /charmbracelet/log

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
```

It is possible to request a specific theme.

```go
log.SetStyles(redlog.Theme("catppuccin"))
```

And with a specific variation of the theme.

```go
log.SetStyles(redlog.Theme("catppuccin", redlog.WithVariant("mocha")))
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
variant := catppuccin.Adaptive(catppuccin.Latte, catppuccin.Mocha)
log.SetStyles(catppuccin.New(variant))
```
