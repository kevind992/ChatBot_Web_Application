//Author: Kevin Delassus
//G00270791
// This file contains the Go code used to run the server.
// It also handles the index.html which is kept in the static folder
 
package main

//Imports
import (

	"fmt"
	"net/http"
	"./eliza"
	
)
//functions handling the user input and eliza
func userInputHandler(w http.ResponseWriter, r *http.Request) {

	//Getting the user input and assigning it to input
	input := r.Header.Get("value")

	//sending the input to eliza to be processes and assigning the response to output
	output := eliza.ElizaStart(input)

	//Returning output to the user
 	fmt.Fprintf(w, "Eliza: %s\n", output)
}
//The main runs the server and parses the static folder which contains index.html
func main() {

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/user-input", userInputHandler)

	//Running of port 8080
	http.ListenAndServe(":8080", nil)
}
