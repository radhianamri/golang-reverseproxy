package main

import (
	"fmt"
	"log"
	"net/http"
)

const indexHTML = `
<!DOCTYPE html>
<html>
	<head>
		<title>Hello World</title>
		<script src="/static/app.js"></script>
		<link rel="stylesheet" href="/static/app.css"">
	</head>
	<body>
	Hello, gopher!<br>
	<img src="https://blog.golang.org/go-brand/logos.jpg" height="100">
	</body>
</html>
`

func main() {
	http.Handle("/static/",
		http.FileServer(http.Dir(".")))

	log.Println("Server started at :443")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		if pusher, ok := w.(http.Pusher); ok {
			httpPushFile("/static/app.js", pusher)
			httpPushFile("/static/app.css", pusher)
		}

		fmt.Fprintf(w, indexHTML)

	})

	err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
	if err != nil {
		panic(err)
	}

}

func httpPushFile(fileName string, pusher http.Pusher) {
	if err := pusher.Push(fileName, nil); err != nil {
		log.Printf("Failed to push: %v", err)
	}
}
