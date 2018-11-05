package main

import (
        "net/http"
)

func write_text(wrt http.ResponseWriter, r *http.Request) {
        wrt.Write([]byte("This is a GO Line HTTP Server\n"))
        wrt.Write([]byte("Hello World"))
}

func main() {
		http.Handle("/", http.FileServer(http.Dir("static")))
        http.HandleFunc("/text", write_text)
        if err := http.ListenAndServe(":9090", nil); err != nil {
                panic(err)
        }
}
