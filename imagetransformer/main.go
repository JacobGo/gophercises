package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./frontend/build")))
	http.HandleFunc("/upload", processImage)
	http.ListenAndServe(":8000", nil)
}

var totalImages int
func processImage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Wrong method.", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(15 << 20) // 10 MB max file size
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing form: %s",err), http.StatusUnprocessableEntity)
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving image: %s",err), http.StatusUnprocessableEntity)
		return
	}
	defer file.Close()

	n := r.FormValue("n")
	if n == "" {
		http.Error(w, "Error: missing n flag", http.StatusUnprocessableEntity)
		return
	}

	mode := r.FormValue("mode")
	if mode == "" {
		http.Error(w, "Error: missing mode flag", http.StatusUnprocessableEntity)
		return
	}

	cmd := exec.Command("primitive",
		"-i", "-",
			"-o", "-",
			"-n", n,
			"-m", mode)
	cmd.Stdin = file
	cmd.Stdout = w

	err = cmd.Start(); if err != nil {
		http.Error(w, fmt.Sprintf("Error from primitive: %s", err), http.StatusInternalServerError)
		return
	}
	cmd.Wait()

	fmt.Printf("Generated %d images.\n", totalImages + 1)
	totalImages++
}
