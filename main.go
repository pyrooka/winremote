package main

import (
	"log"

	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/debug"
)

var prelog debug.Log

func main() {
	inService, err := svc.IsWindowsService()
	if err != nil {
		log.Fatalf("Failed to determine if we are running in service: %v", err)
		return
	}
	if inService {
		runService(serviceName)
		return
	}

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
