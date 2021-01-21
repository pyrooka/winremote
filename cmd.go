package main

import (
	"fmt"
	"log"
	"os"

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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(os.Args)
	},
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

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the service",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the service",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var pauseCmd = &cobra.Command{
	Use:   "pause",
	Short: "Pause the service",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var continueCmd = &cobra.Command{
	Use:   "continue",
	Short: "Continue the service",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
