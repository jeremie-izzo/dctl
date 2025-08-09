/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"github.com/spf13/cobra"
)

// downCmd represents the down command
var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Stop your development environment",
	Long: `Stop your complete development environment including Docker containers and Tilt.

This command:
- Stops all Docker containers managed by godev
- Terminates Tilt if it's running
- Removes orphaned containers  
- Cleans up temporary files

EXAMPLES:
  godev down                    # Stop everything and cleanup`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🛑 Stopping Tilt...")

		// Stop Tilt
		cmdTilt := exec.Command("tilt", "down")
		cmdTilt.Stdout = os.Stdout
		cmdTilt.Stderr = os.Stderr
	
		if err := cmdTilt.Run(); err != nil {
			log.Printf("⚠️  Warning: Failed to stop Tilt: %v", err)
		} else {
			fmt.Println("✅ Tilt stopped successfully!")
		}	},
}

func init() {
	rootCmd.AddCommand(downCmd)
}
