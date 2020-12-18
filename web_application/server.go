package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
)

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })

    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hi")
    })

    http.HandleFunc("/test", HelloServer)

    log.Fatal(http.ListenAndServe(":8080", nil))
}


func HelloServer(w http.ResponseWriter, r *http.Request) {
   fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
