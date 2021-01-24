package main

import (
	"fmt"
	"net/http"
)

func sleepOrHibernateHandler(w http.ResponseWriter, r *http.Request) {
	err := sleepOrHibernate()
	if err != nil {
		elog.Error(1, fmt.Sprintf("failed to lock the computer: %v", err))
	}
}

func shutdownHandler(w http.ResponseWriter, r *http.Request) {
	err := shutdown()
	if err != nil {
		elog.Error(1, fmt.Sprintf("failed to lock the computer: %v", err))
	}
}

func restartHandler(w http.ResponseWriter, r *http.Request) {
	err := restart()
	if err != nil {
		elog.Error(1, fmt.Sprintf("failed to lock the computer: %v", err))
	}
}

func listen() error {
	http.HandleFunc("/sleep", sleepOrHibernateHandler)
	http.HandleFunc("/shutdown", shutdownHandler)
	http.HandleFunc("/restart", restartHandler)

	return http.ListenAndServe(":8899", nil)
}
