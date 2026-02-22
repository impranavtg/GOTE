# 📝 Gote

A minimal, keyboard-driven note-taking app for the terminal — built with Go and [Bubble Tea](https://github.com/charmbracelet/bubbletea).

Gote keeps your notes as Markdown files in `~/.gote`, so they're always plain-text and easy to find.

---

## ✨ Features

- **Create notes** — quickly spin up a new `.md` file from the TUI
- **List & browse notes** — fuzzy-filterable list of all your saved notes
- **Edit in-place** — open any note in a full textarea editor
- **Save & close** — persist changes back to disk with a single keystroke
- **Lightweight** — no database, no config, just files

## ⌨️ Keyboard Shortcuts

| Key      | Action                             |
| -------- | ---------------------------------- |
| `Ctrl+N` | Create a new note                  |
| `Ctrl+L` | List all notes                     |
| `Ctrl+S` | Save the current note              |
| `Esc`    | Go back / close current view       |
| `Enter`  | Confirm input / open selected note |
| `Ctrl+C` | Quit                               |


## 🛠️ Built With

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) — TUI framework
- [Bubbles](https://github.com/charmbracelet/bubbles) — TUI components (text input, textarea, list)
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) — Terminal styling

