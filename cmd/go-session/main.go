package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/thanawatpetchuen/gos/internal/gos"
)

func main() {
	router := gos.NewRouter()
	server := gos.NewServer(router)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	exit := make(chan os.Signal, 1)

	signal.Notify(exit, os.Interrupt)

	<-exit
	if err := server.Close(); err != nil {
		log.Fatal(err)
	}
	log.Println("Shutting down the Server")
}
