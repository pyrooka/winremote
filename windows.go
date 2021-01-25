package main

import (
	"os/exec"
)

func sleepOrHibernate() error {
	// NOTE: if hibernation is off this will sleep the computer else hibernate it.
	_, err := exec.Command("rundll32.exe", "powrprof.dll,SetSuspendState", "0,1,0").Output()
	return err
}

func restart() error {
	_, err := exec.Command("shutdown", "/r").Output()
	return err
}

func shutdown() error {
	// Using the /f switch to force close all the running application.
	_, err := exec.Command("shutdown", "/s", "/f").Output()
	return err
}
