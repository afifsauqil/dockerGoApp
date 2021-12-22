package main

import (
    "log"
    "net/http"
    "fmt"
)

func about(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "About me\nName: Afif Sauqil Arifin\nFrom: Surabaya\nAge: 18 years old")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
        	for _, h := range headers {
            	fmt.Fprintf(w, "%v: %v\n", name, h)
        	}
    	}
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/about", about)
	http.HandleFunc("/headers", headers)

	log.Println("listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
