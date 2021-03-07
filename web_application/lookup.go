package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", testBranch)
    log.Fatal(http.ListenAndServe(":8080", nil))
}


func testBranch(w http.ResponseWriter, r *http.Request) {
     //Read the Request Parameter "command"
     command := r.FormValue("command")

     //Ideally do other checks for tokens/username/etc
     if command == "/categories" {
        fmt.Fprint(w,"Travel\nMapping\nSports\nFinancial\nMusic\nFood\nWeather\nNigeria")
     } else {
         fmt.Fprint(w,"I do not understand your command.")
     }
}
