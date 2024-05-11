# catppuccin

Package catppuccin consists of 4 beautiful pastel color palettes, named
flavors. The theme comes in one light and three dark variants.

* 🌻 *Latte* is the lightest theme harmoniously inverting the essence of
  Catppuccin's dark themes.
  
* 🪴 *Frappé* is a less vibrant alternative using subdued colors for a muted
  aesthetic.

* 🪷 *Macchiato* has medium contrast with gentle colors creating a soothing
  atmosphere.

* 🍀 *Mocha* is the original and the darkest variant, offering a cozy feeling
  with color-rich accents.

## Usage

Use `go get` to download the dependency.

```bash
go get github.com/op/redlog/pkg/catppuccin@latest
```

Then, `import` it in Go files:

```go
import (
  "github.com/charmbracelet/log"
  "github.com/op/redlog/pkg/catppuccin"
)
```

The package comes with a `Default` you can use.

```go
log.SetStyles(catppuccin.Default.Styles)
slog.SetDefault(slog.New(log.Default()))
```

To use a specific variant, request it by name.

```go
// logger always uses Macchiato
logger := log.New(os.Stderr)
logger.SetStyles(catppuccin.Macchiato.Styles)

// update default styles to use an adaptive theme
variant := catppuccin.Adaptive(catppuccin.Latte, catppuccin.Mocha)
log.SetStyles(catppuccin.New(variant))
```
