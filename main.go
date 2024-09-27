package main

import (
    "fmt"
    "log"
    "net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        if err := r.ParseForm(); err != nil {
            http.Error(w, "error in parsing form", http.StatusBadRequest)
            return
        }

        name := r.FormValue("name")
        address := r.FormValue("address")
        if name == "" || address == "" {
            http.Error(w, "no null values allowed", http.StatusBadRequest)
            return
        }
        fmt.Fprintf(w, "Name: %s\nAddress: %s", name, address)
    } 
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    fileServer := http.FileServer(http.Dir("./static"))

    http.Handle("/", fileServer)
    http.HandleFunc("/form", formHandler)
    http.HandleFunc("/hello", helloHandler)

    fmt.Println("Starting server at port 8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
