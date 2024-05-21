package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello %s, welcome to the server.", r.URL.Path[1:])
}

func main(){
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)

	
	//run this in terminal to start the server
	//go to localhost:9000/Brendan in a browser
}