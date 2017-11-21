//Author: Kevin Delassus
 
package main

import (

	"fmt"
	"net/http"
	"./eliza"
	
)



func userInputHandler(w http.ResponseWriter, r *http.Request) {

	//input := r.URL.Query().Get("value")
	input := r.Header.Get("value")

	fmt.Println("The user input is: ",input)

	output := eliza.ElizaStart(input)

	//Returning input and output to the user
	//fmt.Fprintf(w, "User: %s\n", input)
 	fmt.Fprintf(w, "Eliza: %s\n", output)

}
func main() {

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/user-input", userInputHandler)

	http.ListenAndServe(":8080", nil)
}
