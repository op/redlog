package catppuccin

import (
	"log/slog"
	"os"

	"github.com/charmbracelet/log"
)

func Example() {
	logger := log.New(os.Stdout) // use os.Stderr in your application
	// TODO: submit pr to add Styles to the Options struct
	logger.SetStyles(Styles(Adaptive(Latte, Mocha)))
	logger.Info("purrr 🐾")

	log.SetDefault(logger)
	log.Info("purrr 🐱")

	slog.SetDefault(slog.New(logger))
	slog.Info("purrr 😸")

	// Output:
	// INFO purrr 🐾
	// INFO purrr 🐱
	// INFO purrr 😸
}
