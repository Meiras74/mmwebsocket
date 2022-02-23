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

	http.Handle("/", websocket.Handler(Echo))

	log.Fatal(http.ListenAndServe(":"+port, nil))

	fmt.Println("server start on port : " + port)
}

func Echo(ws *websocket.Conn) {

	for {

		fmt.Println(websocket.RemoteAddr())

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
