package main

import (
	"fmt"
	"net/http"
)

func sleepOrHibernateHandler(w http.ResponseWriter, r *http.Request) {
	err := sleepOrHibernate()
	if err != nil {
		elog.Error(1, fmt.Sprintf("failed to sleep/hibernate the computer: %v", err))
	}
}

func shutdownHandler(w http.ResponseWriter, r *http.Request) {
	err := shutdown()
	if err != nil {
		elog.Error(1, fmt.Sprintf("failed to shutdown the computer: %v", err))
	}
}

func restartHandler(w http.ResponseWriter, r *http.Request) {
	err := restart()
	if err != nil {
		elog.Error(1, fmt.Sprintf("failed to restart the computer: %v", err))
	}
}

func listen() {
	http.HandleFunc("/sleep", sleepOrHibernateHandler)
	http.HandleFunc("/shutdown", shutdownHandler)
	http.HandleFunc("/restart", restartHandler)

	err := http.ListenAndServe(":8899", nil)
	if err != nil {
		elog.Error(1, fmt.Sprintf("cannot start the webserver %v", err))
	}
}
