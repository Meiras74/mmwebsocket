package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/websocket"
)

var myconn []*websocket.Conn
var addressAut [3]string = [3]string{"https://meiras.outsystemscloud.com", "https://www.piesocket.com"}

func main() {
	port := os.Getenv("PORT")
	//port := "3000"

	http.Handle("/", websocket.Handler(Echo))
	go CleanClients()
	fmt.Println("server start on port : " + port)

	log.Fatal(http.ListenAndServe(":"+port, nil))

}

func Echo(ws *websocket.Conn) {

	fmt.Println("--" + ws.RemoteAddr().String() + "--")

	if ValidateAddress(ws.RemoteAddr().String()) == false {
		err := websocket.Message.Send(ws, "Origin not valid")
		if err != nil {
			fmt.Println("Can't send Origin not valid")
		}
		ws.Close()

	} else {

		if Contains(ws) == false {
			myconn = append(myconn, ws)
			err := websocket.Message.Send(ws, "Welcome to Miguel Websocket Server")
			if err != nil {
				fmt.Println("Can't send welcome message")
			}
		}

		for {

			//fmt.Println(myconn)

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
				if conn != ws {
					err = websocket.Message.Send(conn, reply)
					if err != nil {
						fmt.Println("Can't send")
					}
				}
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

func ValidateAddress(a string) bool {

	for i := 0; i < len(addressAut); i++ {
		if addressAut[i] == a {
			return true
		}
	}
	return false
}

func CleanClients() {
	for range time.Tick(30 * time.Second) {
		fmt.Println(len(myconn))
		for _, n := range myconn {
			if n.IsServerConn() == false {
				ind := IndexOf(n)
				if ind != -1 {
					Remove(ind)
				}
				n.Close()
			}
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
