package main

import (
  "github.com/go-martini/martini"
  "os"
)

func main() {
  m := martini.Classic()
  m.Get("/", func() string {
    return "Hello Martini!"
  })
  m.RunOnAddr(":" + os.Getenv("HTTP_PLATFORM_PORT"))
}
