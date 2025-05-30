package main

import (
	"strings"

	"github.com/fatih/color"
)

func ColoredText(c, text string) string {
	switch strings.ToLower(c) {
	case "red":
		red := color.New(color.FgRed).SprintFunc()
		return red(text)
	case "blue":
		blue := color.New(color.FgHiBlue).SprintFunc()
		return blue(text)
	case "magenta":
		magenta := color.New(color.FgMagenta).SprintFunc()
		return magenta(text)
	case "cyan":
		cyan := color.New(color.FgCyan).SprintFunc()
		return cyan(text)
	case "green":
		green := color.New(color.FgGreen).SprintFunc()
		return green(text)
	default:
		return text
	}
}
