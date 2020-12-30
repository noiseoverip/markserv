package main

import (
  "flag"
  "log"
  "net/http"
)

func main() {

  dir := flag.String("dir", "", "directory to serve")
  addr := flag.String("addr", ":3000", "IP addr to bind and port to listen on")
  baseURL := flag.String("baseURL", "/", "base URL urlPath")

  flag.Parse()

  fs := http.FileServer(http.Dir(*dir))

  http.Handle(*baseURL, http.StripPrefix(*baseURL, fs))

  log.Printf("Listening on:%s path:%s", *addr, *baseURL)
  err := http.ListenAndServe(*addr, nil)
  if err != nil {
    log.Fatal(err)
  }
}
