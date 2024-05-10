package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/op/redlog"
	"github.com/op/redlog/internal/themes"
)

var (
	flagFontFamily = flag.String("font.family", "Noto Mono", "Font family")
	flagFontSize   = flag.Int("font.size", 16, "Font size")
)

func main() {
	flag.Parse()

	log.SetStyles(redlog.Default)

	for _, t := range themes.Themes {
		for _, v := range t.Variants {

			execute := fmt.Sprintf(
				`build/redlog -theme "%s" -variant "%s"`,
				t.Name, v.Name,
			)
			filename := strings.ToLower(fmt.Sprintf("%s-%s.svg", t.Name, v.Name))
			output := filepath.Join("assets", filename)

			cmd := exec.Command(
				"freeze",
				"--background", v.Background,
				"--font.family", *flagFontFamily,
				"--font.size", strconv.Itoa(*flagFontSize),
				"--border.radius", "16",
				"--border.width", "1",
				"--padding", "20,20,0,20",
				"--execute", execute,
				"--output", output,
			)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			if err := cmd.Run(); err != nil {
				log.Fatal(err)
			}
		}
	}
}
