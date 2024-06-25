# Go Sandbox

A place to try out ideas in Go.

## Continuous Integration

This repository includes an example of how to do some basic Continuous Integration with GitHub
Actions.  See `.github/workflows/` for details.

## Future Work

### `godog`

<https://github.com/cucumber/godog>

### Bubbletea

<https://github.com/charmbracelet/bubbletea/blob/master/examples/simple/main.go>

## Task automation

Sub-projects each have their own `Makefile`, and there is a top-level `Makefile` that ties them all
together.  This is so it's easier to keep them all up to date.

### `make` (default task)

Runs a main program or tests, whichever is used in that sub-project.

### `make tidy`

Runs `go mod tidy` in a sub-project.
