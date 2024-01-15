package main

import (
	"fmt"
	"net/http"
	"strings"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	if player == "Bob" {
		fmt.Fprint(w, "20")
		return
	}
	if player == "Bryce" {
		fmt.Fprint(w, "10")
		return
	}
}
