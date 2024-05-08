// handlers/websocket.go
package handlers

import (
    "github.com/gorilla/websocket"
    "net/http"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
        return
    }
    defer conn.Close()

    for {    
        _, msg, err := conn.ReadMessage()
        if err != nil {
            break
        }
        // message logics are here for example we can implement twilio here
        err = conn.WriteMessage(websocket.TextMessage, msg)
        if err != nil {
			fmt.Println("Error writing message:", err)
            break
        }
    }
}
