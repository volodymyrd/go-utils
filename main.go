package main

import (
	"fmt"
	"github.com/volodymyrd/go-utils/utils/ip_utils"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/my-ip", myIpHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

// indexHandler responds to requests with our greeting.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Utils written in Go")
}

// myIpHandler responds to requests with IP address and location.
func myIpHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Got request %s", r.URL.Path)
	if r.URL.Path != "/my-ip" {
		http.NotFound(w, r)
		return
	}
	var ip = ip_utils.GetIPAddress(r)
	fmt.Fprint(w, "IP "+ip)
	fmt.Fprint(w, "\n")
	fmt.Fprint(w, "Location "+ip_utils.GetLocation(ip))
}
