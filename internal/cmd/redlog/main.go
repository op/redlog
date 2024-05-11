package main

import (
	"flag"
	"fmt"
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
	flagDefault = flag.Bool("default", false, "Allow fallback to default")

	exec = os.Args[0]
	prog = filepath.Base(exec)
)

func main() {
	flag.Parse()

	t, ok := themes.ByName(*flagTheme)
	if !ok && !*flagDefault {
		fmt.Fprintf(os.Stderr, "%s: unknown theme: %s\n\n", prog, *flagTheme)
		os.Exit(1)
	} else if !ok {
		t = themes.Default
	}

	v, ok := themes.VariantByName(t, *flagVariant)
	if !ok && !*flagDefault {
		fmt.Fprintf(os.Stderr, "%s: unknown variant: %s\n\n", prog, *flagVariant)
		os.Exit(1)
	} else if !ok {
		v = t.Default
	}

	writeLog(t, v)
}

func writeLog(t logtheme.Theme, v logtheme.Variant) {
	log.Helper()

	num := 0
	l := log.NewWithOptions(os.Stderr, log.Options{
		TimeFunction: func(_ time.Time) time.Time {
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
