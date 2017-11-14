//Author: Kevin Delassus
 
package main

import (

	"fmt"
	"net/http"
	
)
type Message struct {
    message string
}

func userinputhandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w,"User : %s \n", r.URL.Query().Get("value"))
	fmt.Fprintf(w, "Chatbot : Hello, %s!, How are you today?", r.URL.Query().Get("value")) //.Path[1:])



}
func main() {

	fmt.Println(runEliza())

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/user-input", userinputhandler)

	http.ListenAndServe(":8080", nil)
}