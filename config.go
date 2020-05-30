package progress

import "os"

// Config -
type Config struct {
	Prefix string // prefix before the glyphs
	Suffix string // suffix after the glyphs

	QueuedGlyph   string //
	QueuedMessage string // message to display while waiting to start

	ProgressGlyphs  []string // characters to animate
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
	queuedMessage := ""
	progressMessage := ""
	successMessage := ""
	failMessage := ""

	if len(args) > 0 {
		queuedMessage = args[0]
	}

	if len(args) > 1 {
		progressMessage = args[1]
	} else {
		progressMessage = queuedMessage
	}

	if len(args) > 2 {
		successMessage = args[2]
	} else {
		successMessage = queuedMessage
	}

	if len(args) > 3 {
		failMessage = args[3]
	} else {
		failMessage = queuedMessage
	}

	// https://en.wikipedia.org/wiki/List_of_Unicode_characters

	return &Config{
		Prefix: " ",
		Suffix: " ",

		QueuedGlyph:   string('\u2218'), // void bullet
		QueuedMessage: queuedMessage,

		ProgressGlyphs:  []string{string('\u25B6')}, // u25B6 - play
		ProgressMessage: progressMessage,

		SuccessGlyph:   string('\u2713'), // check mark
		SuccessMessage: successMessage,

		FailGlyph:   string('\u00D7'), // middle cross
		FailMessage: failMessage,

		Writer:     os.Stdout,
		HideCursor: true,
	}
}
