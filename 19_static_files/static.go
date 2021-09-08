package main

import (
	"net/http"
	"fmt"
)

func main() {
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/",fs)
	
	fmt.Println("Server ready in port 3000")
	http.ListenAndServe(":3000",nil)

} 
