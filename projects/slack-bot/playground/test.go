package main

import (
     "fmt"
     "net/http"
)

func main() {
     http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {

     //Read the Request Parameter "command"
     command := r.FormValue("command")

     //Ideally do other checks for tokens/username/etc

     if command == "/categories" {

        fmt.Fprint(w,"Travel\nMapping\nSports\nFinancial\nMusic\nFood\nWeather")
     } else {
         fmt.Fprint(w,"I do not understand your command.")
     }
}
