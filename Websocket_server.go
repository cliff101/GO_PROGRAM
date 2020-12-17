package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {

	// Upgrader specifies parameters for upgrading an HTTP connection to a
	// WebSocket connection.
	upgrader := &websocket.Upgrader{}

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		//upgrader.CheckOrigin = func(r *http.Request) bool { fmt.Println(r); return true }  //警告: 跨域測試時才可解除註解
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("upgrade:", err)
			return
		}
		defer func() {
			log.Println("disconnect !!")
			c.Close()
		}()
		for {
			mtype, msg, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("receive: %s\n", msg)
			err = c.WriteMessage(mtype, msg)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	})
	log.Println("server start at :8899")
	//ListenAndServe starts an HTTP server with a given address and handler.
	//The handler is usually nil, which means to use DefaultServeMux. Handle and HandleFunc add handlers to DefaultServeMux:
	log.Fatal(http.ListenAndServe(":8899", nil))
}
