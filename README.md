# Installer TUI 

TUI-based project installer. Loads a catalog of projects from a JSON config and provides an interactive terminal UI to select and install them.

## Technology

- **Go** 1.26
- **Bubble Tea v2** — TUI framework
- **Lipgloss** — styling
- **Bubbles** — reusable TUI components

## Usage

```sh
go run .
```

### Watch mode (requires [watchexec](https://github.com/watchexec/watchexec))

```sh
watchexec -e go --restart --clear -- go run .
```
