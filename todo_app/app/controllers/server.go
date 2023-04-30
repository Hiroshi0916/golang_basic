package controllers

import (
	"net/http"
	"testing/todo_app/config"
)

func StartmainServer() error {
	http.HandleFunc("/", top)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
