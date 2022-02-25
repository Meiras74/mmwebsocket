package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func main() {
	port := os.Getenv("PORT")
	//port := "3000"

	http.HandleFunc("/", socketHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func socketHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade our raw HTTP connection to a websocket based one

	fmt.Println(r.Header.Get("Origin"))
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Error during connection upgradation:", err)
		return
	}
	defer conn.Close()

	// The event loop
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error during message reading:", err)
			break
		}
		log.Printf("Received: %s", message)
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Println("Error during message writing:", err)
			break
		}
	}
}
