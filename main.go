package main

import (
	"net/http"

	"github.com/gmac220/project-0/navigation"
)

func main() {
	//navigation.Welcome()
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/navigation/formsuccess.html", navigation.Formsubmit)
	http.ListenAndServe(":9000", nil)
}
