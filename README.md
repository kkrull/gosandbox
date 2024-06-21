# Go Sandbox

A place to try out ideas in Go.

## Future Work

### GitHub Actions

- <https://docs.github.com/en/actions/automating-builds-and-tests/building-and-testing-go>
- <https://github.com/actions/setup-go>

### `godog`

<https://github.com/cucumber/godog>

### Cobra

<https://github.com/spf13/cobra>

### Bubbletea

<https://github.com/charmbracelet/bubbletea/blob/master/examples/simple/main.go>

## Task automation

Sub-projects each have their own `Makefile`, and there is a top-level `Makefile` that ties them all
together.  This is so it's easier to keep them all up to date.

### `make` (default task)

Runs a main program or tests, whichever is used in that sub-project.

### `make tidy`

Runs `go mod tidy` in a sub-project.
