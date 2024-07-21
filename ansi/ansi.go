package ansi

import (
	"bytes"
	"fmt"
)

/*** esc codes ***/

const (
	ESC_oct = '\033'
	ESC_hex = '\x1b'
	ESC_utf = '\u001b'
	CSI     = '['
)

/*** color/graphics mode ***/

const (
	RESET = "0"

	BOLD       = "1"
	RESET_BOLD = "22"

	DIM       = "2"
	RESET_DIM = "22"

	ITALICS       = "3"
	RESET_ITALICS = "23"

	UNDERLINE       = "4"
	RESET_UNDERLINE = "24"

	INVERSE       = "7"
	RESET_INVERSE = "27"

	HIDDEN       = "8"
	RESET_HIDDEN = "28"

	STRIKETHROUGH       = "9"
	RESET_STRIKETHROUGH = "29"
)

/*** color codes ***/

const (
	// foreground colors
	FG_DEFAULT = "39"
	FG_BLACK   = "30"
	FG_RED     = "31"
	FG_GREEN   = "32"
	FG_YELLOW  = "33"
	FG_BLUE    = "34"
	FG_MAGENTA = "35"
	FG_CYAN    = "36"
	FG_WHITE   = "37"

	// background colors

	BG_DEFAULT = "49"
	BG_BLACK   = "40"
	BG_RED     = "41"
	BG_GREEN   = "42"
	BG_YELLOW  = "43"
	BG_BLUE    = "44"
	BG_MAGENTA = "45"
	BG_CYAN    = "46"
	BG_WHITE   = "47"
)

// Aprint is an analog for fmt.Print. The passed *Style object
// defines how the code should format msg when printed.
func Aprint(style *Style, msg string) {
	fmt.Print(style.Code())
	fmt.Print(msg)

	// reset styles
	ResetStyles()
}

// Aprintf is an analog for fmt.Printf. The passed *Style object
// defines how the code should format msg when printed. Any additional
// args will be used to fill format specifiers in msg.
func Aprintf(style *Style, msg string, args ...any) {
	fmt.Print(style.Code())

	// ensures formatting doesn't spill into new line
	if lastChar(msg) == '\n' {
		fmt.Printf(msg[:len(msg)-1], args...)
		ResetStyles()
		fmt.Print("\n")
	} else {
		fmt.Printf(msg, args...)
		ResetStyles()
	}
}

// Aprintf is an analog for fmt.Println. The passed *Style object
// defines how the code should format msg whne printed. Prints an
// addition newline after msg
func Aprintln(style *Style, msg string) {
	fmt.Print(style.Code())
	fmt.Print(msg)

	ResetStyles()
	fmt.Print("\x1b[0J") // erase from cursor to end of screen so highlighting stops after text ends
	fmt.Print("\n")
}

// Asprint is an analog for fmt.Sprint. The passed *Style object
// defines how msg should be formatted when printed. returns printable
// value as a string.
func Asprint(style *Style, msg string) string {
	var out bytes.Buffer

	out.WriteString(style.Code())
	out.WriteString(msg)
	out.WriteString(ResetStyleString())

	return out.String()
}

// Asprintf is an analog for fmt.Sprintf. The passed *Style object
// defines how msg should be formatted when printed. Any additional
// args will be used to fill format specifiers in msg. returns printable
// value as a string.
func Asprintf(style *Style, msg string, args ...any) string {
	var out bytes.Buffer

	out.WriteString(style.Code())

	// ensure formatting doesn't spill into new line
	if lastChar(msg) == '\n' {
		out.WriteString(fmt.Sprintf(msg[:len(msg)-1], args...))
		out.WriteString(ResetStyleString())
		out.WriteByte('\n')
	} else {
		out.WriteString(fmt.Sprintf(msg, args...))
		out.WriteString(ResetStyleString())
	}

	return out.String()
}

func ResetStyles() {
	fmt.Print("\x1b[22;23;24;25;26;28;29;0m")
}

func ResetStyleString() string {
	return fmt.Sprintf("\x1b[22;23;24;25;26;28;29;0m")
}

/*** Style Sequences ***/

type Style struct {
	codes []string // a list of the commands
	args  string   // ';' delimited string of commands
}

func NewStyle(args ...string) *Style {
	c := []string{}
	if len(args) == 0 {
		c = append(c, "0")
	}

	for _, arg := range args {
		c = append(c, arg)
	}

	seq := &Style{codes: c}
	seq.initArgs()

	return seq
}

func (s *Style) initArgs() {
	var out bytes.Buffer

	for i, c := range s.codes {
		out.WriteString(c)
		if i != len(s.codes)-1 {
			out.WriteRune(';')
		}
	}

	s.args = out.String()
}

func (s *Style) Args() string {
	if len(s.args) == 0 {
		s.initArgs()
	}

	return s.args
}

func (s *Style) Code() string {
	return fmt.Sprintf("%c%c%sm", ESC_hex, CSI, s.Args())
}

/*** helpers ***/

func lastChar(s string) byte {
	if len(s) < 1 {
		return 0
	}

	return s[len(s)-1]

}
