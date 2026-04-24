package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stderr, "server ", log.LstdFlags)

	// создаём сервер, передавая ему логгер
	srv := server.NewServer(logger)
	http.ListenAndServe(":8080", srv.Server.Handler)
	// запускаем сервер
	log.Println("starting server on :8080")
	if err := srv.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
