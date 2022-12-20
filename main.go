package main

import (
    "errors"
    "log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(201)
		w.Write([]byte("hello there"))
	})

	server := http.Server{
        Addr: ":8080",
        Handler: mux,
    }

    err := server.ListenAndServe()
    if err != nil {
        if errors.Is(err, http.ErrServerClosed) {
            return
        }
        log.Fatalln("An error occured:", err.Error())
    }
}
