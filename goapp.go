#Simple app which listens on 8484 and shows hostname of server from which it is being serived.
package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	h, _ := os.Hostname()
	fmt.Fprintf(w, "Hi there, I'm served from %s!", h)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8484", nil)
}
