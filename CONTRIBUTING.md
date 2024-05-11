# Contributing

Clone the repository using [`git dolly`](https://github.com/op/dolly) or `git
clone`.

```bash
git dolly https://github.com/op/redlog
```

Running `make` is usually sufficient since it does everything. There is help
available with `make help`.

```bash
make clean && make
```

This also updates all example assets which you will find in
[assets/](./assets/). You can also test this in your terminal by running the
`redlog` command.

```bash
# use make to build
make build/redlog
# run with default theme (adaptive catppuccin)
./build/redlog -default
# run directly with go and use ctppuccin only
go run ./internal/cmd/redlog -theme catppuccin -default
# run with a specific variant
go run ./internal/cmd/redlog -theme catppuccin -variant latte
```
