package main

import (
	"log"

	"github.com/spf13/cobra"
	"golang.org/x/sys/windows/svc"
)

const serviceName = "winremote"

func init() {
	rootCmd.AddCommand(installCmd)
	rootCmd.AddCommand(uninstallCmd)
}

var rootCmd = &cobra.Command{
	Use:   "winremote",
	Short: "Turn off/restart/hibernate/sleep your computer remotely via HTTP",
	Run: func(cmd *cobra.Command, args []string) {
		inService, err := svc.IsWindowsService()
		if err != nil {
			log.Fatalf("Failed to determine if we are running in service: %v", err)
			return
		}
		if inService {
			runService(serviceName)
			return
		}
	},
}

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install the service",
	Run: func(cmd *cobra.Command, args []string) {
		if err := installService(serviceName, serviceName); err != nil {
			log.Fatalf("Cannot install the %s service: %v", serviceName, err)
		}
	},
}

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall the service",
	Run: func(cmd *cobra.Command, args []string) {
		if err := uninstallService(serviceName); err != nil {
			log.Fatalf("Cannot uninstall the %s service: %v", serviceName, err)
		}
	},
}
