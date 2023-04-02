package main

import (
	"fmt"
	"log"
	"os"

	chordpro "github.com/jcuffney/tabtoshi/chordpro"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	filepath := "./data/that-funny-feeling__bo-burnham.cbpro"
	data, err := chordpro.ReadTextFile(filepath)
	if err != nil {
		err = fmt.Errorf("failed to read file %s", filepath)
		fmt.Println(err)
		os.Exit(1)
	}

	song, err := chordpro.Parse(data)
	if err != nil {
		err = fmt.Errorf("failed to parse file %w", err)
		fmt.Println(err)
		os.Exit(1)
	}

	// chordpro.Print(song)

	// chordpro.PrintJson(song)

	// chordpro.WriteJsonToFile("./data/that-funny-feeling__bo-burnham.json", song)

	chordpro.WritePDF("./data/that-funny-feeling__bo-burnham.pdf", song)
}
