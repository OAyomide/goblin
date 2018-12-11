package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	//we want to import mux
	r := mux.NewRouter()

	//decalre out port
	port := getPort()
	//we want to call the function to set all these the things we are setting --get started button payload, etc
	setGetStartedPayload("GET_STARTED")

	//declare the VERB OUR route takes
	r.HandleFunc("/", indexHandler).Methods("GET")
	r.HandleFunc("/webhook", webhookGetHandler).Methods("GET")
	r.HandleFunc("/webhook", webhookPostHandler).Methods("POST")
	fmt.Printf("Server up and running. Running on PORT: %s\n", port)
	//pass to our good ol http
	err := http.ListenAndServe(port, r)

	//if there is an error starting our server
	if err != nil {
		log.Fatal("Error listening and server:", err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	fmt.Fprint(w, "Got my server up and running in Go. Yay!!")
}

//a function to get our port incase we deploy to heroku
func getPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		/**
		*TODO: get the port declared in the yml config.
		 */
		port = ":3500"
		fmt.Printf("PORT NOT DEFINED. USING THE PORT %s as the running port\n", port)
	}
	return port
}
