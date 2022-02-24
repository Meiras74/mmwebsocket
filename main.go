package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

// T receive and send Json
type T struct {
	Name   string
	Msg    string
	Action int
}

var allClients map[string]clients

type clients struct {
	Conn *websocket.Conn
	Name string
}

// NewClient cria novo cliente
func NewClient(ws *websocket.Conn, login string) {
	clients := new(clients)

	clients.Conn = ws
	clients.Name = login

	allClients[clients.Name] = *clients

	fmt.Println(allClients)
}

// Echo echo from srver
func Echo(ws *websocket.Conn) {
	//fmt.Println(ws.Config())

	for {

		//var reply string
		var data T

		//err := websocket.Message.Receive(ws, &reply)
		err := websocket.JSON.Receive(ws, &data)
		if err != nil {
			fmt.Println("Error receive : " + err.Error())
			break
		}

		if data.Action == 1 {
			if _, exists := allClients[data.Name]; !exists {
				NewClient(ws, data.Name)
				data.Msg = data.Name + " as login"
				for _, key := range allClients {
					if key.Conn != ws {
						err = websocket.JSON.Send(key.Conn, data)
						if err != nil {
							fmt.Println("Can't send")
							//break
						}
					}
				}
			} else {
				data.Msg = " error : " + data.Name + " already login"

				//err := websocket.Message.Send(ws, msg);
				err = websocket.JSON.Send(ws, data)
				if err != nil {
					fmt.Println("Can't send")
				}
				//break
			}

		} else if data.Action == 3 {
			for _, key := range allClients {
				if key.Conn != ws {
					data.Msg = data.Name + " Logout "
					err = websocket.JSON.Send(key.Conn, data)
					if err != nil {
						fmt.Println("Can't send")
						//break
					}
				}
			}
			for _, key := range allClients {
				if key.Conn == ws {
					delete(allClients, key.Name)
					data.Msg = data.Name + " Logout "
					err = key.Conn.Close()
					if err != nil {
						fmt.Println("Can't close")
						//break
					}
				}
			}
			fmt.Println(allClients)
		} else {
			//msg := "Received:  " + reply
			for _, key := range allClients {
				if key.Conn != ws {
					data.Msg = data.Name + " send: " + data.Msg
					err = websocket.JSON.Send(key.Conn, data)
					if err != nil {
						fmt.Println("Can't send")
						//break
					}
				}
			}
		}

	}
}

func main() {
	var port string

	flag.StringVar(&port, "port", "8181", "port for server")
	flag.Parse()

	fmt.Println(port)

	fmt.Println("Start Server")

	allClients = make(map[string]clients)

	http.Handle("/", websocket.Handler(Echo))

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
