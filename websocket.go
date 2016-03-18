package main

import (
    "golang.org/x/net/websocket"
    "io"
    "net/http"
)

func EchoHandler(ws *websocket.Conn) {
    io.Copy(ws, ws)
}

func main() {
    http.Handle("/echo", websocket.Handler(EchoHandler))
    err := http.ListenAndServe(":" + os.Getenv("HTTP_PLATFORM_PORT"), nil)
    if err != nil {
        panic("ListenAndServe: " + err.Error())
    }
}
