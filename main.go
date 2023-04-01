package main

import (
	"fmt"
	"os"

	"tabtoshi/chordpro"
)

func main() {
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

	chordpro.Print(song)
}
