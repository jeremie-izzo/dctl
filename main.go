/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/jeremie-izzo/dctl/cmd"
	_ "github.com/jeremie-izzo/dctl/internal/templater/compose"
)

func main() {
	// Initialize values at startup
	cmd.Execute()
}
