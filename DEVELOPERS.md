# DEVELOPERS.md

## Introduction

`tview-command` is a flexible keybinding manager designed to be integrated into terminal user interfaces (TUIs) built with [tview](https://github.com/rivo/tview). This package allows developers to manage keybindings using a context-sensitive system, where different key mappings can be activated depending on the current context (e.g., player queue management, text fields, modals).

## Getting Started

To test and experiment with `tview-command`, you can use the `cmd/example1` directory as a test rig. This will allow you to explore how keybindings work with contexts in a live application.

### Step 1: Clone and Navigate to the Example

Make sure you have cloned the project and navigate to the example directory:

```bash
git clone https://github.com/spezifisch/tview-command
cd tview-command/cmd/example1
```

### Step 2: Run the Example

You can run the example using the following command:

```bash
go run -tags example main.go
```

This will launch a simple TUI where you can test the keybindings. The bindings are configured using `config_example1.toml` in the same directory.
