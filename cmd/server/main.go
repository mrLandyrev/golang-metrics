package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/update", handleUpdate)

	http.ListenAndServe(":8000", router)
}

func handleUpdate(w http.ResponseWriter, h *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}
