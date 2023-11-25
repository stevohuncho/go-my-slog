package gomyslog

import (
	"log/slog"

	"github.com/fatih/color"
)

type Styling struct {
	StylingColors
	StylingChars
	ShowType    bool
	ShowTime    bool
	MaxValueLen int
}

type Colors []Color

type Color color.Attribute

func (c Color) InvertGround() Color {
	if c >= 30 && c <= 37 {
		return c + 10
	}
	if c >= 40 && c <= 47 {
		return c - 10
	}
	if c >= 90 && c <= 97 {
		return c + 10
	}
	if c >= 100 && c <= 107 {
		return c - 10
	}
	return 6969
}

func (c Color) IsBg() bool {
	return (c >= 40 && c <= 47) || (c >= 100 && c <= 107) || (c >= 0 && c <= 9)
}

func (c Color) IsFg() bool {
	return (c >= 30 && c <= 37) || (c >= 90 && c <= 97) || (c >= 0 && c <= 9)
}

func (c Colors) Bg() Colors {
	colors := Colors{}
	for _, color := range c {
		if color.IsBg() {
			colors = append(colors, color)
		}
	}
	return colors
}

func (c Colors) Fg() Colors {
	colors := Colors{}
	for _, color := range c {
		if color.IsFg() {
			colors = append(colors, color)
		}
	}
	return colors
}

func (c Colors) InvertGround() Colors {
	colors := make(Colors, len(c))
	for i, color := range c {
		colors[i] = color.InvertGround()
	}
	return colors
}

func (c Colors) ToAttrs() []color.Attribute {
	attrs := []color.Attribute{}
	for _, _color := range c {
		attrs = append(attrs, color.Attribute(_color))
	}
	return attrs
}

func (c Colors) Render() *color.Color {
	return color.New(c.ToAttrs()...)
}

type LevelColorMap map[slog.Level]Colors

type StylingColors struct {
	level  LevelColorMap
	time   Colors
	prefix Colors
	msg    Colors
	typing Colors
	key    Colors
	value  Colors
	tree   Colors
	ktov   Colors
}

type StylingChars struct {
	leftCap      string
	rightCap     string
	rightArrow   string
	topCurve     string
	bottomCurve  string
	t            string
	sideT        string
	dash         string
	l            string
	topCorner    string
	bottomCorner string
	dottedDash   string
	dottedL      string
}

func DefaultStyling() Styling {
	return Styling{
		StylingColors: StylingColors{
			level: LevelColorMap{
				slog.LevelDebug: Colors{Color(color.BgMagenta), Color(color.FgWhite)},
				slog.LevelInfo:  Colors{Color(color.BgBlue), Color(color.FgWhite)},
				slog.LevelWarn:  Colors{Color(color.BgYellow), Color(color.FgWhite)},
				slog.LevelError: Colors{Color(color.BgRed), Color(color.FgWhite)},
			},
			time:   Colors{Color(color.BgHiBlack), Color(color.FgBlack)},
			prefix: Colors{Color(color.FgWhite), Color(color.BgBlack), Color(color.Bold)},
			msg:    Colors{Color(color.BgHiBlue), Color(color.FgWhite)},
			typing: Colors{Color(color.FgGreen)},
			key:    Colors{Color(color.FgBlue), Color(color.Bold)},
			value:  Colors{Color(color.Italic), Color(color.Bold)},
			tree:   nil,
			ktov:   nil,
		},
		StylingChars: StylingChars{
			leftCap:      "",
			rightCap:     "",
			rightArrow:   "",
			topCurve:     "╭",
			bottomCurve:  "╰",
			t:            "┬",
			sideT:        "├",
			dash:         "─",
			l:            "│",
			topCorner:    "┐",
			bottomCorner: "└",
			dottedDash:   "╌",
			dottedL:      "╎",
		},
		ShowType:    true,
		ShowTime:    true,
		MaxValueLen: 50,
	}
}
