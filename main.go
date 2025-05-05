package main

import (
	"log"
	"net/http"
	"time"

	messagingHttp "alexandre-gerault.fr/gochat-server/internal/messaging/ui/http"
	shared_infrastructure "alexandre-gerault.fr/gochat-server/internal/shared/infrastructure"
)



func main() {
	app := shared_infrastructure.Application {}

	app.Boot()
	shared_infrastructure.RunMigrations()

	log.Println("Start http server...")

	router := http.NewServeMux()

	router.HandleFunc("POST /messages/", messagingHttp.SendMessageEndpoint)

	httpServer := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(httpServer.ListenAndServe())

	log.Println("Http server started successfully")
}
