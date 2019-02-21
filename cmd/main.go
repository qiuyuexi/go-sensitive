package main

import (
	"go-sensitive/internal/app/api"
	"go-sensitive/internal/pkg/server"
	"net/http"
)

func main() {
	server.Start()
	http.HandleFunc("/filter",api.FilterHandel);
	http.ListenAndServe(":8081",nil)
}
