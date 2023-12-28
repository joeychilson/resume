package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/joeychilson/resume/resume"
	"github.com/joeychilson/resume/server/page"
)

func main() {
	http.Handle("/example.resume.json", http.FileServer(http.Dir(".")))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		page.Page().Render(r.Context(), w)
	})

	http.HandleFunc("/resume", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method is not supported.", http.StatusNotFound)
			return
		}

		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			http.Error(w, "File is too large.", http.StatusBadRequest)
			return
		}

		file, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Error retrieving the file.", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		resumeStruct := resume.Resume{}
		err = json.NewDecoder(file).Decode(&resumeStruct)
		if err != nil {
			http.Error(w, "Error decoding JSON.", http.StatusInternalServerError)
			return
		}

		err = resume.Page(resumeStruct).Render(r.Context(), w)
		if err != nil {
			http.Error(w, "Error rendering resume.", http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("/example", func(w http.ResponseWriter, r *http.Request) {
		exampleResume, err := http.Get("http://localhost:8080/example.resume.json")
		if err != nil {
			http.Error(w, "Error retrieving example resume.", http.StatusInternalServerError)
			return
		}

		resumeStruct := resume.Resume{}
		err = json.NewDecoder(exampleResume.Body).Decode(&resumeStruct)
		if err != nil {
			http.Error(w, "Error decoding JSON.", http.StatusInternalServerError)
			return
		}

		err = resume.Page(resumeStruct).Render(r.Context(), w)
		if err != nil {
			http.Error(w, "Error rendering resume.", http.StatusInternalServerError)
			return
		}
	})

	log.Println("Listening on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}
