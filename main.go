package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		log.Print(r.Header.Get("Origin"))
		if ValidateAddress(r.Header.Get("Origin")) != true {
			return false
		}
		return true
	},
}
var addressAut [2]string = [2]string{"https://meiras.outsystemscloud.com", "https://www.piesocket.com"}

var myconn []*websocket.Conn

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
	//defer conn.Close()

	//log.Print(conn)
	if Contains(conn) == false {
		myconn = append(myconn, conn)
		msg := []byte("Welcome to Miguel Websocket Server")
		err = conn.WriteMessage(1, msg)
		if err != nil {
			log.Print("Can't send welcome message")
		}
	}

	// The event loop
	for {

		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error during message reading:", err)
			break
		}
		log.Printf("Received: %s", message)

		log.Print(messageType)
		for _, ws := range myconn {
			err := ws.WriteMessage(messageType, message)
			if err != nil {
				log.Println("Error during message writing:", err)
				break
			}
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

func Contains(x *websocket.Conn) bool {
	for _, n := range myconn {
		if x == n {
			return true
		}
	}
	return false
}
