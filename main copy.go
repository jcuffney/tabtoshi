// package main

// import (
// 	"bytes"
// 	"fmt"
// 	"io"
// 	"os"
// 	"regexp"
// 	"strings"
// )

// type Line struct {
// 	text string
// 	// mapping of the index of the chord to the chord
// 	chords map[int]string
// }

// type Section struct {
// 	Type  string
// 	Lines []Line
// }

// type Song struct {
// 	Metadata map[string]string
// 	Sections []Section
// }

// func readTextFile(filepath string) (string, error) {
// 	file, err := os.Open(filepath)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to open file %s: %w", filepath, err)
// 	}
// 	defer file.Close()

// 	fileContent, err := io.ReadAll(file)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to read file %s: %w", filepath, err)
// 	}

// 	return string(fileContent), nil
// }

// func parseLine(line string) (Line, error) {
// 	var chords map[int]string = make(map[int]string)
// 	// parse out any chords
// 	re := regexp.MustCompile(`\[.*?\]`)
// 	matches := re.FindAllString(line, -1)
// 	indexes := re.FindAllStringIndex(line, -1) // returns an array of starting and endpoint index [0,3]
// 	offset := 0
// 	for idx, match := range matches {
// 		chordIndexes := indexes[idx]
// 		start := chordIndexes[0]
// 		end := chordIndexes[1]
// 		delta := end - start
// 		chords[start+offset] = match
// 		offset -= delta
// 	}
// 	line = re.ReplaceAllString(line, "")
// 	return Line{text: line, chords: chords}, nil
// }

// func parseMetadata(data string) (string, map[string]string, error) {
// 	var metadata map[string]string = make(map[string]string)
// 	re := regexp.MustCompile("{.*:.*}")
// 	matches := re.FindAllString(data, -1)
// 	for _, match := range matches {
// 		arr := strings.Split(match, ":")
// 		metadata[arr[0]] = arr[1]
// 	}
// 	// remove metadata from data string
// 	data = re.ReplaceAllString(data, "")
// 	return data, metadata, nil
// }

// func getSection(line string) string {
// 	re := regexp.MustCompile(`\[\[.*\]\]`)
// 	match := re.FindString(line)
// 	return match
// }

// func getChordsLine(chords map[int]string) string {
// 	if len(chords) == 0 {
// 		return ""
// 	}
// 	bytes := bytes.Repeat([]byte{0x20}, 100)
// 	for key := range chords {
// 		chord := []byte(chords[key])
// 		for i := 0; i < len(chord); i++ {
// 			bytes[key+i] = chord[i]
// 		}
// 	}
// 	return string(bytes)
// }

// func parse(data string) (Song, error) {
// 	song := Song{}
// 	var err error

// 	// parse out metadata
// 	data, metadata, err := parseMetadata(data)
// 	if err != nil {
// 		err := fmt.Errorf("failed to parse metadata %w", err)
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// 	song.Metadata = metadata

// 	sections := []Section{}
// 	section := Section{ // start with a empty section
// 		Type:  "",
// 		Lines: []Line{},
// 	}

// 	// check for metadata
// 	lines := strings.Split(data, "\n")
// 	// parse each line
// 	for _, line := range lines {
// 		// if the line is empty - continue
// 		line := strings.TrimSpace(line)
// 		if line == "" {
// 			continue
// 		}

// 		// check if the line is a section - note this needs to go before the line is parsed.
// 		sectionName := getSection(line)
// 		if sectionName != "" {
// 			// before replacing the current section - check if it should be added
// 			if len(section.Lines) > 0 {
// 				sections = append(sections, section)
// 			}
// 			// replace existing section with the new one
// 			section = Section{
// 				Type:  sectionName,
// 				Lines: []Line{},
// 			}
// 			continue
// 		}

// 		// since it's not a section - it's a line
// 		l, err := parseLine(line)
// 		if err != nil {
// 			err := fmt.Errorf("failed to parse line %w", err)
// 			fmt.Println(err)
// 			os.Exit(1)
// 		}
// 		section.Lines = append(section.Lines, l)
// 	}
// 	// add the last section if it has at least 1 line
// 	if len(section.Lines) > 0 {
// 		sections = append(sections, section)
// 	}

// 	song.Sections = sections

// 	return song, nil
// }

// func main() {
// 	filepath := "song.cbpro"
// 	data, err := readTextFile(filepath)
// 	if err != nil {
// 		err = fmt.Errorf("failed to read file %s", filepath)
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	song, err := parse(data)
// 	if err != nil {
// 		err = fmt.Errorf("failed to parse file %w", err)
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	print(song)
// }
