package main

import (
	"log"

	"github.com/spf13/cobra"
)

const name = "winremote"

func init() {
	rootCmd.AddCommand(installCmd)
	rootCmd.AddCommand(uninstallCmd)
}

var rootCmd = &cobra.Command{
	Use:   "winremote",
	Short: "Turn off/hibernate/sleep/lock your computer remotely via HTTP",
}

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install the service",
	Run: func(cmd *cobra.Command, args []string) {
		if err := installService(name, name); err != nil {
			log.Fatalf("Cannot install the %s service: %s", name, err)
		}
	},
}

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall the service",
	Run: func(cmd *cobra.Command, args []string) {
		if err := uninstallService(name); err != nil {
			log.Fatalf("Cannot uninstall the %s service: %s", name, err)
		}
	},
}
