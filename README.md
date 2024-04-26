# rofi-browser

rofi-browser is a simple profile launcher for Mozilla Firefox-based browsers,
enabling you to quickly launch a new profile instance from a [rofi](https://github.com/davatorium/rofi) menu.

## Compatibility

Currently works with Mozilla Firefox and LibreWolf. I'll gladly accept PR's to integrate with other browsers.

## Installation

### Arch Linux (AUR)

Using your favorite AUR helper, install `rofi-browser`.

### Binaries

Get a binary directly from the [Github Releases](https://github.com/obvionaoe/rofi-browser/releases)

### Go

Install using go:

```bash
go install github.com/obvionaoe/rofi-browser
```

### Build from source

To build from source, simply run:

```bash
make install
```

## Usage

```bash
rofi-browser
```
A [rofi](https://github.com/davatorium/rofi) menu will appear, allowing you to select and launch your desired browser
(by default, Firefox) profile.
