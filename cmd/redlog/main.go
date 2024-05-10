package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/charmbracelet/log"

	"github.com/op/redlog/internal/logtheme"
	"github.com/op/redlog/internal/themes"
)

var (
	flagTheme   = flag.String("theme", "", "Theme to use")
	flagVariant = flag.String("variant", "", "Variant to use")

	exec = os.Args[0]
	prog = filepath.Base(exec)
)

func main() {
	flag.Parse()

	w := os.Stdout

	t, ok := themes.ByName(*flagTheme)
	if !ok {
		if *flagTheme != "" {
			fmt.Fprintf(os.Stderr, "%s: unknown theme: %s\n\n", prog, *flagTheme)
			os.Exit(1)
		}
		for _, t := range themes.Themes {
			writeJSON(w, struct {
				Name string `json:"name"`
			}{
				Name: strings.ToLower(t.Name),
			})
		}
		os.Exit(0)
	}

	v, ok := themes.VariantByName(t, *flagVariant)
	if !ok {
		if *flagVariant != "" {
			fmt.Fprintf(os.Stderr, "%s: unknown variant: %s\n\n", prog, *flagVariant)
			os.Exit(1)
		}
		for _, f := range t.Variants {
			writeJSON(w, struct {
				Name       string `json:"name"`
				Background string `json:"background"`
			}{
				Name:       strings.ToLower(f.Name),
				Background: f.Background,
			})
		}
		os.Exit(0)
	}

	writeLog(t, v)
}

func writeJSON(w io.Writer, v any) {
	e := json.NewEncoder(w)
	e.SetEscapeHTML(false)
	if err := e.Encode(v); err != nil {
		log.Fatal(err)
	}
}

func writeLog(t logtheme.Theme, v logtheme.Variant) {
	log.Helper()

	num := 0
	l := log.NewWithOptions(os.Stderr, log.Options{
		TimeFunction: func(t time.Time) time.Time {
			num = num + 1
			return time.Unix(1234567+int64(math.Pow10(num)), 0)
		},
		TimeFormat:      "15:04",
		Level:           log.DebugLevel,
		Prefix:          v.Name,
		ReportTimestamp: true,
		ReportCaller:    true,
		CallerFormatter: func(string, int, string) string {
			line := (num-1)*int(math.Exp(float64(num))) + 42
			caller := fmt.Sprintf("%s/%s:%d", t.Name, v.Name, line)
			return strings.ToLower(caller)
		},
	})
	l.SetStyles(v.Styles)

	l.Debug("Starting oven", "style", strings.ToLower(t.Name))
	l.Info("Mixing ingredients", "ingredients",
		strings.Join([]string{
			"1dl of butter",
			"2dl of chocolate",
			"3dl of flour",
		}, "\n"),
	)
	l.Warn("Baking cookies", "flavour", v.Name)
	l.Error("Cookie got burnt 💥")
	l.Log(log.FatalLevel, "The kitchen is on fire 🔥")
}
