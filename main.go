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

	if len(os.Args) < 3 {
		log.Fatal("Usage: go run main.go <resume.json> <output.html>")
	}

	resumeFileName := os.Args[1]
	outputFileName := os.Args[2]

	log.Println("Generating " + outputFileName + "...")

	log.Println("Reading " + resumeFileName + "...")
	resumeJSON, err := os.Open(resumeFileName)
	if err != nil {
		log.Fatalf("failed to open %s: %v", resumeFileName, err)
	}
	defer resumeJSON.Close()

	log.Println("Parsing " + resumeFileName + "...")
	resumeStruct := resume.Resume{}
	err = json.NewDecoder(resumeJSON).Decode(&resumeStruct)
	if err != nil {
		log.Fatalf("failed to parse %s: %v", resumeFileName, err)
	}

	log.Println("Creating " + outputFileName + "...")
	resumeHTML, err := os.Create(outputFileName)
	if err != nil {
		log.Fatalf("failed to create %s: %v", outputFileName, err)
	}
	defer resumeHTML.Close()

	err = resume.Page(resumeStruct).Render(ctx, resumeHTML)
	if err != nil {
		log.Fatalf("failed to write page: %v", err)
	}

	log.Println("Successfully generated " + outputFileName + "!")
}
