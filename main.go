package main

import (
  "log"
  "net/http"
)

var d = new(DataBase)

func main(){
  d.Init("indigo")
  http.HandleFunc("/", handleRequests)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
