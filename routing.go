package main

import (
	"fmt"
	"net/http"
)

type application struct {
}

func startAPI() {
	fmt.Println("Starting webserver...")

	//create new mux
	mux := http.NewServeMux()
	app := application{}
	//route traffic that hits /getbooks to the getBooks method
	mux.HandleFunc("/getbooks", app.getBooks)
	//same as above
	mux.HandleFunc("/postbooks", app.postBooks)
	//Advertise webserver on localhost:8080, so localhost:8080/getbooks
	//will trigger getBooks() method
	http.ListenAndServe(":8080", mux)
}
