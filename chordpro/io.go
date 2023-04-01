package chordpro

import (
	"encoding/json"
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

func WriteJsonToFile(filepath string, song Song) error {
	json := ToJson(song)
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(json)
	if err != nil {
		return err
	}
	return nil
}

func Print(song Song) {
	for idx, section := range song.Sections {
		fmt.Println(section.Type)
		for _, line := range section.Lines {
			chordsLine := getChordsLine(line.Chords)
			if chordsLine != "" {
				fmt.Println(chordsLine)
			}
			fmt.Println(line.Text)
		}
		if idx != len(song.Sections)-1 {
			fmt.Println("")
		}
	}
}

func ToJson(song Song) []byte {
	json, err := json.MarshalIndent(song, "", "  ")
	if err != nil {
		err := fmt.Errorf("failed to print json %w", err)
		fmt.Println(err)
		os.Exit(1)
	}
	return json
}

func PrintJson(song Song) {
	jsonString := string(ToJson(song))
	fmt.Println(jsonString)
}
