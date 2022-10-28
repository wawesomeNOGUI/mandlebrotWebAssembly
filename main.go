package main

import (
	"net/http"
)

func main () {
	fileServer := http.FileServer(http.Dir("./public"))
	http.Handle("/", fileServer)

	err := http.ListenAndServe(":80", nil) //Http server blocks
	if err != nil {
		panic(err)
	}
}	