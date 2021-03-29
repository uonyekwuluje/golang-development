package main

import (
   "fmt"
   "net/http"
   "html/template"
   "log"
   "time"
)

type PageVariables struct {
   Date string
   Time string
}


func helloWorld(w http.ResponseWriter, r *http.Request){
    now := time.Now() // find the time right now
    HomePageVars := PageVariables{ //store the date and time in a struct
      Date: now.Format("02-01-2006"),
      Time: now.Format("15:04:05"),
    }

    t, err := template.ParseFiles("home.html") //parse the html file homepage.html
    if err != nil { // if there is an error
	  log.Print("template parsing error: ", err) // log it
    }

    err = t.Execute(w, HomePageVars) //execute the template and pass it the HomePageVars struct to fill in the gaps
    if err != nil { // if there is an error
	  log.Print("template executing error: ", err) //log it
    }
}

func main() {
    http.HandleFunc("/", helloWorld)

    fmt.Println("[INFO] : Server listening on PORT 8080")
    http.ListenAndServe(":8080", nil)
}
