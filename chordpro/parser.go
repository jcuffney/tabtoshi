package chordpro

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func parseLine(line string) (Line, error) {
	var chords map[int]string = make(map[int]string)
	// parse out any chords
	re := regexp.MustCompile(`\[.*?\]`)
	matches := re.FindAllString(line, -1)
	indexes := re.FindAllStringIndex(line, -1) // returns an array of starting and endpoint index [0,3]
	offset := 0
	for idx, match := range matches {
		chordIndexes := indexes[idx]
		start := chordIndexes[0]
		end := chordIndexes[1]
		delta := end - start
		chords[start+offset] = match
		offset -= delta
	}
	line = re.ReplaceAllString(line, "")
	return Line{text: line, chords: chords}, nil
}

func parseMetadata(data string) (string, map[string]string, error) {
	var metadata map[string]string = make(map[string]string)
	re := regexp.MustCompile("{.*:.*}")
	matches := re.FindAllString(data, -1)
	for _, match := range matches {
		arr := strings.Split(match, ":")
		metadata[arr[0]] = arr[1]
	}
	// remove metadata from data string
	data = re.ReplaceAllString(data, "")
	return data, metadata, nil
}

func getSection(line string) string {
	re := regexp.MustCompile(`\[\[.*\]\]`)
	match := re.FindString(line)
	return match
}

func getChordsLine(chords map[int]string) string {
	if len(chords) == 0 {
		return ""
	}
	bytes := bytes.Repeat([]byte{0x20}, 100)
	for key := range chords {
		chord := []byte(chords[key])
		for i := 0; i < len(chord); i++ {
			bytes[key+i] = chord[i]
		}
	}
	return string(bytes)
}

func Parse(data string) (Song, error) {
	song := Song{}
	var err error

	// parse out metadata
	data, metadata, err := parseMetadata(data)
	if err != nil {
		err := fmt.Errorf("failed to parse metadata %w", err)
		fmt.Println(err)
		os.Exit(1)
	}
	song.Metadata = metadata

	sections := []Section{}
	section := Section{ // start with a empty section
		Type:  "",
		Lines: []Line{},
	}

	// check for metadata
	lines := strings.Split(data, "\n")
	// parse each line
	for _, line := range lines {
		// if the line is empty - continue
		line := strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// check if the line is a section - note this needs to go before the line is parsed.
		sectionName := getSection(line)
		if sectionName != "" {
			// before replacing the current section - check if it should be added
			if len(section.Lines) > 0 {
				sections = append(sections, section)
			}
			// replace existing section with the new one
			section = Section{
				Type:  sectionName,
				Lines: []Line{},
			}
			continue
		}

		// since it's not a section - it's a line
		l, err := parseLine(line)
		if err != nil {
			err := fmt.Errorf("failed to parse line %w", err)
			fmt.Println(err)
			os.Exit(1)
		}
		section.Lines = append(section.Lines, l)
	}
	// add the last section if it has at least 1 line
	if len(section.Lines) > 0 {
		sections = append(sections, section)
	}

	song.Sections = sections

	return song, nil
}
