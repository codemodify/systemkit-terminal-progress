package progress

import "os"

// Config -
type Config struct {
	Prefix          string   // prefix before the animated characters
	ProgressGlyphs  []string // characters to animate
	Suffix          string   // suffix after the animated characters
	ProgressMessage string   // message to display while the operation is going

	SuccessGlyph   string //
	SuccessMessage string // message to display on success

	FailGlyph   string //
	FailMessage string // message to display on failure

	Writer     *os.File //
	HideCursor bool     //
}
