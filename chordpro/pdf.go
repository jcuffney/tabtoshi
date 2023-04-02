package chordpro

import (
	"github.com/signintech/gopdf"
)

const LineHeight = 14
const margin = 50 // can i get this from the config?
var config = gopdf.Config{
	PageSize: *gopdf.PageSizeA4,
}

func addLine(pdf *gopdf.GoPdf, line string) {
	pdf.Cell(nil, line+"\n")
	pdf.Br(LineHeight)
	if pdf.GetY() > config.PageSize.H-margin {
		pdf.AddPage()
	}
}

func WritePDF(filepath string, song Song) error {
	pdf := gopdf.GoPdf{}
	pdf.Start(config)
	pdf.AddPage()

	// Set the font family and size
	pdf.AddTTFFont("Arial", "./chordpro/assets/arial.ttf")
	pdf.SetFont("Arial", "", 14)

	pdf.CellWithOption(nil, song.Metadata["title"]+"\n", gopdf.CellOption{Align: gopdf.Center, Border: 1})
	pdf.Br(5 * LineHeight)

	for idx, section := range song.Sections {
		addLine(&pdf, section.Type)
		for _, line := range section.Lines {
			chordsLine := getChordsLine(line.Chords)
			if chordsLine != "" {
				addLine(&pdf, chordsLine)
			}
			addLine(&pdf, line.Text)
		}
		if idx != len(song.Sections)-1 {
			pdf.Br(2 * LineHeight)
		}

	}

	// Write some text to the page
	// pdf.Cell(nil, "Hello, World!")

	// pdf.AddHeader(func() {
	// 	pdf.SetY(5)
	// 	pdf.Cell(nil, "header")
	// })
	// pdf.AddFooter(func() {
	// 	pdf.SetY(825)
	// 	pdf.Cell(nil, "footer")
	// })

	// Save the document to a file
	err := pdf.WritePdf(filepath)
	if err != nil {
		panic(err)
	}
	return nil
}
