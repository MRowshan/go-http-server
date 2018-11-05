# go-http-server

The app.go file can be run using: \n
	`go run app.go`
or
	```
	go build app.go
	./app (linux)
	./app.go (windows)
	```

###Running the app.go file hosts a webpage (/static/index.html) on port 9090


##A brief look at the code is shown below:

A function which fills a http webpage with the following plain text
```
func write_text(wrt http.ResponseWriter, r *http.Request) {
        wrt.Write([]byte("This is a GO Line HTTP Server\n"))
        wrt.Write([]byte("Hello World"))
}
```

Method which serves files in the static directory and assigns them to the url /
`http.Handle("/", http.FileServer(http.Dir("static")))`
Method which serves the function "write_text" and assigns output to the url "/text"
`http.HandleFunc("/text", write_text)`
Serve webpages on port 9090
```
if err := http.ListenAndServe(":9090", nil); err != nil {
panic(err)
}
```
