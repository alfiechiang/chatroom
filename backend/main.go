package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

func main() {

	http.HandleFunc("/ws", func(rw http.ResponseWriter, r *http.Request) {

		conn, err := websocket.Accept(rw, r,  &websocket.AcceptOptions{InsecureSkipVerify: true})
		if err != nil {
			log.Println(err)
			return
		}

		defer conn.Close(websocket.StatusInternalError, "Internal Error")
		ctx, cancel := context.WithTimeout(r.Context(), time.Second*30)
		defer cancel()

		var v interface{}
		err = wsjson.Read(ctx, conn, &v)
		if err != nil {
			log.Println(err)
			return
		}
		conn.Close(websocket.StatusNormalClosure, "")

	})
	log.Fatal(http.ListenAndServe(":8000", nil))

}
