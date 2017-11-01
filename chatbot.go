//Author: Kevin Delassus
 
package main

import (
	//"fmt"
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	
	http.ListenAndServe(":8080", nil)
}