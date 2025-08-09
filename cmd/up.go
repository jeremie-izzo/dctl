/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// upCmd represents the up command
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Start your development environment with Tilt",
	Long: `Start your complete development environment using Tilt based on the current profile.

This command generates a Tiltfile from your profile configuration and starts Tilt
to manage your development environment with live reloading and monitoring.

EXAMPLES:
  godev up                    # Start development environment with Tilt

Tilt provides a better development experience with live reloading and a web UI.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🚀 Starting development environment with Tilt...")

		// Generate docker-compose.yaml
		// compose, err := internal.LoadProfileTemplate("docker-compose")
		// if err != nil {
		// 	fmt.Printf("❌ Failed to load profile: %v\n", err)
		// 	return
		// }

		// // Write docker-compose.yaml to current directory
		// composePath := "docker-compose.yaml"
		// file, err := os.Create(composePath)
		// if err != nil {
		// 	log.Fatalf("❌ Failed to create docker-compose.yaml: %v", err)
		// }
		// defer file.Close()

		// // Write the template output directly
		// fmt.Fprint(file, compose)

		// // Generate Tiltfile
		// tiltfileContent, err := internal.LoadProfileTemplate("tiltfile")
		// if err != nil {
		// 	fmt.Printf("❌ Failed to load tiltfile template: %v\n", err)
		// 	return
		// 	}

		// // Write Tiltfile to current directory
		// tiltfilePath := "Tiltfile"
		// tiltfile, err := os.Create(tiltfilePath)
		// if err != nil {
		// 	log.Fatalf("❌ Failed to create Tiltfile: %v", err)
		// }
		// defer tiltfile.Close()

		// // tiltfileContent is now a string
		// fmt.Fprint(tiltfile, tiltfileContent)

		// fmt.Println("📄 Generated docker-compose.yaml and Tiltfile")
		// fmt.Println("🚀 Starting Tilt...")

		// // Run Tilt
		// cmdTilt := exec.Command("tilt", "up")
		// cmdTilt.Stdout = os.Stdout
		// cmdTilt.Stderr = os.Stderr

		// if err := cmdTilt.Run(); err != nil {
		// 	log.Fatalf("❌ Failed to run Tilt: %v", err)
		// }
	},
}

func init() {
	rootCmd.AddCommand(upCmd)
}
