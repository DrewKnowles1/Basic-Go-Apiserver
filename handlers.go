package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

//define a struct that contains all the details we want to keep
//for our book record
type book struct {
	ID       string `json: "id"`
	Title    string `json: "title"`
	Author   string `json: "author"`
	Quantity int    `json: "quantity"`
}

//Create a slice of type book (slice is just a handy array, that can be
// resizedas far as we are concerned here)
var books = []book{
	{ID: "1234", Title: "Lotr", Author: "tolken", Quantity: 5},
	{ID: "1235", Title: "hobbit", Author: "tolken", Quantity: 10},
	{ID: "1236", Title: "history of time", Author: "hawking", Quantity: 7},
}

// getBooks, can be called on the application object we defined in routing.go
func (a *application) getBooks(w http.ResponseWriter, r *http.Request) {
	//This endpoint should only respond to GET methods
	if r.Method != http.MethodGet {
		message := fmt.Sprintf("unsupported method %v", r.Method)
		http.Error(w, message, http.StatusMethodNotAllowed)
		//exit function if the method (i.e r.Method, wasnt "GET")
		return
	}
	//Marshall the books slice we created above, to variable 'b'
	b, err := json.Marshal(books)
	if err != nil {
		fmt.Println(err)
	}
	//Store the sting value of 'b' in var 'n'
	n := string(b)
	//printing value of n, to w, basically our response
	fmt.Fprintf(w, "%s", n)

}

//Post request, to add new books
func (a *application) postBooks(w http.ResponseWriter, r *http.Request) {
	//Again we this time check the method type the user sends is a POST
	if r.Method != http.MethodPost {
		message := fmt.Sprintf("unsupported Method %v", r.Method)
		http.Error(w, message, http.StatusMethodNotAllowed)
		//Print error and exit if not
		return
	}
	//read in the body of the message the user has POSTED to us
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "body cannot be read", http.StatusBadRequest)
		return
	}
	var newBook book
	//populate the var of "newBook", with the json unmarshalled value of body
	err = json.Unmarshal(body, &newBook)
	if err != nil {
		http.Error(w, "malformed json", http.StatusBadRequest)
		return
	}
	//this needs validation, not super happy with how ive tackled this
	// but at least makes sure that the Author/ID/Title are present before adding
	//to our slice of books

	// if newBook.Author == "" || newBook.ID == "" || newBook.Title == "" {
	// 	http.Error(w, "Malformed json", http.StatusBadRequest)
	// 	return
	// }

	_, err = verifyBook(newBook)
	if err != nil {
		http.Error(w, "Malformed json", http.StatusBadRequest)
		return
	}

	entry := book{newBook.ID, newBook.Title, newBook.Author, newBook.Quantity}
	books = append(books, entry)
	w.WriteHeader(http.StatusAccepted)

}
