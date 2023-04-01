package chordpro

import (
	"fmt"
	"io"
	"os"
)

func ReadTextFile(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", fmt.Errorf("failed to open file %s: %w", filepath, err)
	}
	defer file.Close()

	fileContent, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s: %w", filepath, err)
	}

	return string(fileContent), nil
}

func Print(song Song) {
	for idx, section := range song.Sections {
		fmt.Println(section.Type)
		for _, line := range section.Lines {
			chordsLine := getChordsLine(line.chords)
			if chordsLine != "" {
				fmt.Println(chordsLine)
			}
			fmt.Println(line.text)
		}
		if idx != len(song.Sections)-1 {
			fmt.Println("")
		}
	}
}
