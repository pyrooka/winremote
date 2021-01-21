package main

import (
	"errors"
	"os/exec"
)

func lock() error {
	_, err := exec.Command("rundll32.exe", "user32.dll,LockWorkStation").Output()
	return err
}

func sleep() error {
	return errors.New("sleep is not implemented yet")
}

func hibernate() error {
	_, err := exec.Command("rundll32.exe", "powrprof.dll,SetSuspendState", "0,1,0").Output()
	return err
}

func restart() error {
	_, err := exec.Command("shutdown", "/r").Output()
	return err
}

func turnOff() error {
	// Using the /f switch to force close all the running application.
	_, err := exec.Command("shutdown", "/s", "/f").Output()
	return err
}
