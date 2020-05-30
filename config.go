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

// NewDefaultConfig -
func NewDefaultConfig(args ...string) *Config {
	progressMessage := ""
	successMessage := ""
	failMessage := ""

	if len(args) > 0 {
		progressMessage = args[0]
	}

	if len(args) > 1 {
		successMessage = args[1]
	} else {
		successMessage = progressMessage
	}

	if len(args) > 2 {
		failMessage = args[2]
	} else {
		failMessage = progressMessage
	}

	return &Config{
		Prefix:          " ",
		ProgressGlyphs:  []string{string('\u25B6')}, // u00BB - double arrow, u25B6 - play
		Suffix:          " ",
		ProgressMessage: progressMessage,
		SuccessGlyph:    string('\u2713'), // check mark
		SuccessMessage:  successMessage,
		FailGlyph:       string('\u00D7'), // middle cross
		FailMessage:     failMessage,
		Writer:          os.Stdout,
		HideCursor:      true,
	}
}
