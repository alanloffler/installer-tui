# Installer TUI 

TUI-based project installer. Loads a catalog of projects from a JSON config and provides an interactive terminal UI to select and install them.

## Technology

- [**Go**](https://go.dev) 1.26
- [**Bubble Tea v2**](https://github.com/charmbracelet/bubbletea) — TUI framework
- [**Lipgloss**](https://github.com/charmbracelet/lipgloss) — styling
- [**Bubbles**](https://github.com/charmbracelet/bubbles) — reusable TUI components

## Usage

```sh
go run .
```

### Watch mode (requires [watchexec](https://github.com/watchexec/watchexec))

```sh
watchexec -e go --restart --clear -- go run .
```
