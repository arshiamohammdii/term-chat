package chat

const (
	// Reset resets the color
	Reset = "\033[0m"

	// Bold makes the following text bold
	Bold = "\033[1m"

	// Dim dims the following text
	Dim = "\033[2m"

	// Italic makes the following text italic
	Italic = "\033[3m"

	// Underline underlines the following text
	Underline = "\033[4m"

	// Blink blinks the following text
	Blink = "\033[5m"

	// Invert inverts the following text
	Invert = "\033[7m"

	// Newline
	Newline = "\r\n"

	// BEL
	Bel = "\007"
)

type style string

type Style interface {
	String() string
	Render(s string) string
}

func (c style) String() string {
	return string(c)
}

func (c style) Render(s string) string {
	return c.String() + s + Reset
}

type Theme struct {
	System style
	Text   style
	User   style
}

func InitTheme() *Theme {
	return &Theme{
		System: Dim,
		User:   Italic,
		Text:   Bold,
	}
}
