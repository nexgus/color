package color

import (
	"fmt"
	"strconv"
	"strings"
)

// Attribute defines a single SGR Code
type Attribute int

const escape = "\x1b"

// Base attributes
const (
	Normal Attribute = iota // Reset
	Bold                    // increased intensity
	Faint                   // decreased intensity, or dim
	Italic
	Underline
	BlinkSlow
	BlinkRapid
	ReverseVideo // invert
	Concealed    // hide
	CrossedOut   // strike
)

const (
	NormalBold Attribute = iota + 22
	NormalItalic
	NormalUnderline
	NormalBlinking
	_
	NormalReversed
	NormalConcealed
	NormalCrossedOut
)

// Foreground text colors
const (
	FgBlack Attribute = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

// Foreground Hi-Intensity text colors
const (
	FgHiBlack Attribute = iota + 90
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

// Background text colors
const (
	BgBlack Attribute = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)

// Background Hi-Intensity text colors
const (
	BgHiBlack Attribute = iota + 100
	BgHiRed
	BgHiGreen
	BgHiYellow
	BgHiBlue
	BgHiMagenta
	BgHiCyan
	BgHiWhite
)

func Set(attrs ...Attribute) string {
	format := make([]string, len(attrs))
	for idx, attr := range attrs {
		format[idx] = strconv.Itoa(int(attr))
	}
	sequence := strings.Join(format, ";")
	return fmt.Sprintf("%s[%sm", escape, sequence)
}

func Reset() string {
	return fmt.Sprintf("%s[%dm", escape, Normal)
}
