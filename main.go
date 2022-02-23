package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/websocket"
)

var myconn []*websocket.Conn

func main() {
	port := os.Getenv("PORT")

	http.Handle("/", websocket.Handler(Echo))

	log.Fatal(http.ListenAndServe(":"+port, nil))

	fmt.Println("server start on port : " + port)
}

func Echo(ws *websocket.Conn) {

	for {

		fmt.Println(ws.RemoteAddr())
		//falta fazer tratamento do address

		if Contains(ws) == false {
			myconn = append(myconn, ws)
		}

		fmt.Println(myconn)

		var reply string

		err := websocket.Message.Receive(ws, &reply)
		if err != nil {
			fmt.Println("Error receive : " + err.Error())
			ind := IndexOf(ws)
			if ind != -1 {
				Remove(ind)
			}
			break
		}

		//reply = "Echo from server " + reply

		for _, conn := range myconn {
			err = websocket.Message.Send(conn, reply)
			if err != nil {
				fmt.Println("Can't send")
			}
		}

	}
}

func Contains(x *websocket.Conn) bool {
	for _, n := range myconn {
		if x == n {
			return true
		}
	}
	return false
}

func IndexOf(x *websocket.Conn) int {
	var count int
	count = 0
	for _, n := range myconn {
		if x == n {
			return count
		}
		count = count + 1
	}
	return -1
}

func Remove(i int) {
	myconn = append(myconn[:i], myconn[i+1:]...)
}

/*func main() {
	port := os.Getenv("PORT")

	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}*/

/*func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}*/
