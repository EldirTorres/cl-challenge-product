package controller

import (
	"fmt"
	"net/http"
)

func HealtCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The app is up")
}
