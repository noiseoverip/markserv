package main

import (
  "flag"
  "log"
  "net/http"
)

func main() {

  dir := flag.String("dir", "", "directory to serve")
  addr := flag.String("addr", ":3000", "IP addr to bind and port to liten on")

  fs := http.FileServer(http.Dir(*dir))
  http.Handle("/", fs)

  log.Printf("Listening on %s", addr)
  err := http.ListenAndServe(*addr, nil)
  if err != nil {
    log.Fatal(err)
  }
}
