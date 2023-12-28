package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/joeychilson/resume/resume"
)

func main() {
	ctx := context.Background()

	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <resume.json>")
	}

	resumeFileName := os.Args[1]

	log.Println("Generating resume.html...")

	log.Println("Reading resume.json...")
	resumeJSON, err := os.Open(resumeFileName)
	if err != nil {
		log.Fatalf("failed to open resume.json: %v", err)
	}
	defer resumeJSON.Close()

	log.Println("Parsing resume.json...")
	resumeStruct := resume.Resume{}
	err = json.NewDecoder(resumeJSON).Decode(&resumeStruct)
	if err != nil {
		log.Fatalf("failed to parse resume.json: %v", err)
	}

	log.Println("Creating resume.html...")
	resumeHTML, err := os.Create("resume.html")
	if err != nil {
		log.Fatalf("failed to create resume.html: %v", err)
	}
	defer resumeHTML.Close()

	err = resume.Page(resumeStruct).Render(ctx, resumeHTML)
	if err != nil {
		log.Fatalf("failed to write page: %v", err)
	}

	log.Println("Successfully generated resume.html!")
}
