package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/websocket"
)

func main() {
	port := os.Getenv("PORT")

	fmt.Println("server start in port : " + port)

	http.Handle("/", websocket.Handler(Echo))

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func Echo(ws *websocket.Conn) {

	for {
		var reply string
		err := websocket.Message.Receive(ws, &reply)
		if err != nil {
			fmt.Println("Error receive : " + err.Error())
			break
		}

		reply = "Echo from server " + reply

		err = websocket.Message.Send(ws, reply)
		if err != nil {
			fmt.Println("Can't send")
		}
	}
}

/*func main() {
	port := os.Getenv("PORT")

	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}*/

/*func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}*/
