package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	PORT := ":8001"
	m := http.NewServeMux()
	srv := &http.Server{
		Addr:         PORT,
		Handler:      m,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}
	m.HandleFunc("/time", timeHandler)
	m.HandleFunc("/", myHandler)
	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}
func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)
}
func timeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC1123)
	Body := "The current time is:"
	fmt.Fprintf(w, "<h1 align=\"center\">%s</h1>", Body)
	fmt.Fprintf(w, "<h2 align=\"center\">%s</h2>n", t)
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served time for :%s\n", r.Host)
}
