/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/jeremie-izzo/dctl/pkg/registry"
	"github.com/jeremie-izzo/dctl/pkg/tilt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// templateCmd represents the template command
var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "Generate templates ",
	Run: func(cmd *cobra.Command, args []string) {
		path := viper.GetString("config")
		cmd.Println(path)
		file, err := config.Load(path)
		if err != nil {
			cmd.PrintErrln("Error loading config: %v", err)
			return
		}
		reg := registry.Global

		// This for now is hardcoded to "tilt" since its only supported template engine
		engine, ok := reg.TemplateEngine("tilt")
		if !ok {
			fmt.Println("No engine found for compose")
			os.Exit(1)
		}

		// This is the provider that will be used to generate the artifacts based on provider
		// hardcoded to "compose" for now, but can be extended to support other providers
		pro, ok := reg.ArtifactProvider("compose")
		if !ok {
			fmt.Println("No provider for compose")
			os.Exit(1)
		}

		b := tilt.NewEmitter()
		for name, svc := range file.Services {
			// not sure if this is the correct way to use the provider at this point
			preset, exists := reg.Preset(name)
			if exists {
				for k, v := range svc.{
			}
			err := pro.AddService(name, svc)

			// This is a builder for tilt plugins it faciliates the creation of plugins by adding redefine resources to the plan
			if svc.Plugins != nil {
				// Iterate over the plugins defined in the service and apply them
				for _, plugin := range svc.Plugins {
					p, ok := reg.TiltPlugin(plugin)
					if !ok {
						fmt.Printf("Plugin %s not found\n", plugin)
						os.Exit(1)
					}
					p.AddService()
					if err := p.Build(svc, b); err != nil {
						fmt.Printf("Error preparing plugin %s for service %s: %v\n", plugin, name, err)
						os.Exit(1)
					}
				}
			}
		}
		pro.Finalize()

		plan := BuiltPlan(b.Build()) // Build the plan using the builder

		content := engine.Render(plan) // Render the plan using the selected engine (this generates tiltfile content)
	},
}

func init() {
	// sensible defaults + bindings
	templateCmd.Flags().StringP("config", "c", "devctl.yaml", "Path to config file")
	templateCmd.Flags().StringP("out", "o", "", "Output directory for rendered files")

	_ = viper.BindPFlag("config", templateCmd.Flags().Lookup("config"))
	_ = viper.BindPFlag("out", templateCmd.Flags().Lookup("out"))

	rootCmd.AddCommand(templateCmd)
	rootCmd.AddCommand(templateCmd)
}
