package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jmaeso/parser-luna/app"
	lunahttp "github.com/jmaeso/parser-luna/infrastructure/http"
	"github.com/jmaeso/parser-luna/infrastructure/storage/memory"
)

func main() {
	messagesStore := memory.NewMessagesStore()

	messageHandler := lunahttp.MessageHandler{
		MessagesStorage: messagesStore,
	}

	rocketsHandler := lunahttp.RocketsHandler{
		RocketStateService: app.NewRocketStateService(messagesStore),
	}

	http.HandleFunc("POST /messages", messageHandler.PostMessage)
	http.HandleFunc("GET /rockets", rocketsHandler.ListRockets)
	http.HandleFunc("GET /rockets/{id}", rocketsHandler.GetRocketByID)

	port := ":8088"
	fmt.Printf("Server starting on port %s\n", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
