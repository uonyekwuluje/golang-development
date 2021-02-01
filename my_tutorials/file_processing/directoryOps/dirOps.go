package dops

import (
    "fmt"
    "io/ioutil"
     "log"
)

func ListRootDirectory() {
    files, err := ioutil.ReadDir("/")
    if err != nil {
        log.Fatal(err)
    }

    for _, f := range files {
            fmt.Println(f.Name())
    }
}
