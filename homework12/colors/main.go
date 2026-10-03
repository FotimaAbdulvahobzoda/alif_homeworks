package main

import (
	"github.com/fatih/color"
)

func PrintCheck(name string, passed bool) {
	if passed {
		color.Green("OK: %s", name)
	} else {
		color.Red("FAIL: %s", name)
	}
}

func main() {
	PrintCheck("Login test", true)
	PrintCheck("Database test", false)
	PrintCheck("Password test", true)
	PrintCheck("Email test", false)
}
