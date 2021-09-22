package main

import (
  "fmt"
  "log"
  "net/http"
  "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
  response := os.Getenv("RESPONSE_BODY")
  fmt.Fprintln(w, response)
}

func main() {
  http.HandleFunc("/", handler)
  log.Fatal(http.ListenAndServe(":9093", nil))
}
