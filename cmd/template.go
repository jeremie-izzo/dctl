/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var outputPath string
var templateType string

// templateCmd represents the template command
var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "Generate templates from the active profile",
	Long: `Generate various templates from the active profile.

Supported template types:
- docker-compose: Generate docker-compose.yaml
- tiltfile: Generate Tiltfile

EXAMPLES:
  godev template --type docker-compose --output docker-compose.yaml
  godev template --type tiltfile --output Tiltfile`,
	Run: func(cmd *cobra.Command, args []string) {
		//
		//composeFile, err := orchestrator.Template()
		//if err != nil {
		//	cmd.PrintErrf("❌ Failed to generate template: %v\n", err)
		//	return
		//}
		//
		//yaml.NewEncoder(os.Stdout).Encode(composeFile)
		//os.Stdout.WriteString("\n")

	},
}

func init() {
	templateCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Output path for the generated template (optional)")
	templateCmd.Flags().StringVarP(&templateType, "type", "t", "", "Template type: docker-compose or tiltfile (required)")

	rootCmd.AddCommand(templateCmd)
}
