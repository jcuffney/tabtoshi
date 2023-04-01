package chordpro

type Line struct {
	// the line of text representing the lyrics without annotations
	text string
	// mapping of the char index of the chord to the chord string
	chords map[int]string
}

type Section struct {
	Type  string
	Lines []Line
}

type Song struct {
	Metadata map[string]string
	Sections []Section
}
