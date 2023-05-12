package controllers

import (
	"log"
	"net/http"
	"testing/todo_app/config"
)

func StartMainServer() {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/", Top)
	err := http.ListenAndServe(":"+config.Config.Port, nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
