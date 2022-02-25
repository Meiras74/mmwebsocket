package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		fmt.Println(r.Header.Get("Origin"))
		if ValidateAddress(r.Header.Get("Origin")) != true {
			http.Error(w, "Origin not allowed", http.StatusForbidden)
			return false
		}
		return true
	},
}
var addressAut [2]string = [2]string{"https://meiras.outsystemscloud.com"}

func main() {
	port := os.Getenv("PORT")
	//port := "3000"

	http.HandleFunc("/", socketHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func socketHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade our raw HTTP connection to a websocket based one

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

func ValidateAddress(a string) bool {

	for i := 0; i < len(addressAut); i++ {
		if addressAut[i] == a {
			return true
		}
	}
	return false
}
