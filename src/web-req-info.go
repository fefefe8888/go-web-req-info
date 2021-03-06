package main

import (
	"fmt"
	"log"
	"strings"
	"net/http"
)

// HelloServer responds to requests with the given URL path.
func HelloServer(w http.ResponseWriter, r *http.Request) {

	var user_agent string = r.Header.Get("User-Agent")
	var cookies string = r.Header.Get("Cookie")
	var requester_ip string = r.RemoteAddr
	var xff string = r.Header.Get("X-Forwarded-For")
	var xfh string = r.Header.Get("X-Forwarded-Host")
	var query string = r.URL.RawQuery

	fmt.Fprintf(w, "You requested to:\n%s\n\n", r.Method)

	if len(xfh) > 0 {
		fmt.Fprintf(w, "Your original request is for the host:\n%s\n\n", xfh)
	}

	fmt.Fprintf(w, "Your requested URL is:\n%s%s\n\n", r.Host, r.URL.Path)

	if len(query) > 0 {
		fmt.Fprintf(w, "Query in your request:\n%s\n\n", query)
	}

	fmt.Fprintf(w, "You request is from this IP address:\n%s\n\n", requester_ip)

	if len(xff) > 0 {
		fmt.Fprintf(w, "Your request is going through the following IP addresses:\n")
		ips := strings.Split(xff, ", ")
		for _, ip := range ips {
			fmt.Fprintf(w, "%s\n", ip)
		}
		fmt.Fprintf(w, "\n")
	}

	if len(cookies) > 0 {
		fmt.Fprintf(w, "Your Cookies are:\n%s\n\n", cookies)
	}

	if len(user_agent) > 0 {
		fmt.Fprintf(w, "Your User-Agent is:\n%s\n", user_agent)
	}

	log.Printf("Received request from %s for path: %s from %s", requester_ip, r.URL.Path, user_agent)
}

func main() {

	var addr string = ":80"
	handler := http.HandlerFunc(HelloServer)

	log.Printf("Starting webserver on %s", addr)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("Could not listen on port %s %v", addr, err)
	}
}
