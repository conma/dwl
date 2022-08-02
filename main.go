package main

import (
  "dwl/entity"
  "dwl/handler"
  "flag"
  "log"
  "net/http"
)

var (
  fs = flag.NewFlagSet("dwl", flag.ExitOnError)
  httpAddr = fs.String("http-addr", ":8080", "HTTP server address")
)

func main() {
  db, err := handler.Connect()

  if err != nil {
    log.Fatalln(err)
  }

  db.AutoMigrate(&entity.Link{})

  server := &http.Server {
    Handler: handler.NewHandler(db),
    Addr: *httpAddr,
  }

  if err := server.ListenAndServe(); err != nil {
    log.Fatalln(err)
  }
}