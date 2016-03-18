package main

import (
    "golang.org/x/net/websocket"
    "io"
    "net/http"
)

func echoHandler(ws *websocket.Conn) {
    io.Copy(ws, ws)
}

func main() {

    http.Handle("/", websocket.Handler(echoHandler))
    http.ListenAndServe(":" + os.Getenv("HTTP_PLATFORM_PORT"), nil)
}
