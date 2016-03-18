package main

import (
  "fmt"
  "net/http"
  "os"
  "runtime"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, yokohama.go " + runtime.Version())
}

func main() {
  http.HandleFunc("/", viewHandler)
  http.ListenAndServe(":" + os.Getenv("HTTP_PLATFORM_PORT"), nil)
}
