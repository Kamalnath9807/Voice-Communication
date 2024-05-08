package main

import (
    "net/http"
    "github.com/gorilla/websocket"
    "fmt"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println("Error upgrading to WebSocket:", err)
        return
    }
    defer conn.Close()
}

func main() {
    http.HandleFunc("/ws", handleWebSocket)
    http.ListenAndServe(":8080", nil)
}
