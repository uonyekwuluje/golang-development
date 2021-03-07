package main
/*
This code is slash driven text bot for slack
*/


import (
    "fmt"
    "log"
    "net/http"
)

func main() {

    http.HandleFunc("/", interactionHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}


func interactionHandler(w http.ResponseWriter, r *http.Request) {
     slackcommand := r.FormValue("command")

     if slackcommand == "/interaction" {
	fmt.Fprint(w,"What's up buddy!?!?!")
     } else if slackcommand == "/build" {
        fmt.Fprint(w,"Getting Ready to Build")
     } else if slackcommand == "/deploy" {
	fmt.Fprint(w,"Getting Ready to Deploy")
     } else {
        fmt.Fprint(w,"I do not understand your command.")
    }
}
