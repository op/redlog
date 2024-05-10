package catppuccin

import (
	"os"

	"github.com/charmbracelet/log"
)

func Example() {
	logger := log.New(os.Stdout) // use os.Stderr in your application
	// TODO: submit pr to add Styles to the Options struct
	logger.SetStyles(New(Adaptive(Latte, Mocha)))
	logger.Info("purrr 🐾")
	// Output: INFO purrr 🐾
}
