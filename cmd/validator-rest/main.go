package main

import (
	"cards/validator/internal"
	"net/http"
	"time"
)

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/validate", internal.MethodHandler)

	serv := &http.Server{
		Addr:         "0.0.0.0:3000",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 15,
		Handler:      r,
	}

	panic(serv.ListenAndServe())
}
