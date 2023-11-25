package gomyslog

import (
	"fmt"
	"log/slog"
)

type Fields struct {
	attrs []slog.Attr
	Styling
}

type Padding struct {
	space int
	char  string
}

func (f Fields) Render() string {
	log := ""
	for i, attr := range f.attrs {
		log += "\n"
		// add its personal prefix
		if i+1 == len(f.attrs) {
			log += f.tree.Render().Sprint(f.bottomCurve)
		} else {
			log += f.tree.Render().Sprint(f.sideT)
		}
		log += f.tree.Render().Sprint(f.dash + f.dash)
		// render fields
		switch v := attr.Value.Any().(type) {
		case []slog.Attr:
			typeLen := 0
			if f.ShowType {
				log += f.typing.Render().Sprint("(group)")
				typeLen = 7
			}
			log += f.key.Render().Sprint(attr.Key)
			if len(v) == 1 {
				log += f.tree.Render().Sprint(f.dash)
			} else {
				log += f.tree.Render().Sprint(f.t)
			}
			log += f.NewFields(v).RenderGroup("", []Padding{{space: 3 + typeLen + len(attr.Key), char: f.l}}...)
		default:
			// prevent wierd text wrapping
			lines := []string{""}
			linesIdx := 0
			for i, c := range fmt.Sprint(v) {
				if c == '\n' {
					linesIdx += 1
					lines = append(lines, "")
				} else if (i+1)%f.MaxValueLen == 0 {
					linesIdx += 1
					lines = append(lines, "")
					lines[linesIdx] += f.value.Render().Sprint(string(c))
				} else {
					lines[linesIdx] += f.value.Render().Sprint(string(c))
				}
			}
			typeLen := 0
			if f.ShowType {
				typeString := fmt.Sprintf("%T", v)
				log += f.typing.Render().Sprintf("(%T)", typeString)
				typeLen = len(typeString) + 2
			}
			log += f.key.Render().Sprint(attr.Key)
			if len(lines) == 1 {
				log += f.ktov.Render().Sprint(f.dottedDash)
				log += f.value.Render().Sprint(lines[0])
			} else {
				for i, line := range lines {
					if i != 0 {
						log += "\n"
						for _, padding := range []Padding{{space: 3 + typeLen + len(attr.Key), char: f.l}} {
							log += f.tree.Render().Sprintf("%-*s", padding.space, f.l)
						}
					}
					if i == 0 {
						log += f.ktov.Render().Sprint(f.topCorner)
					} else if i+1 == len(lines) {
						log += f.ktov.Render().Sprint(f.bottomCorner)
					} else {
						log += f.ktov.Render().Sprint(f.dottedL)
					}
					log += f.value.Render().Sprint(line)
				}
			}
		}
	}
	return log
}

func (f Fields) RenderGroup(log string, paddings ...Padding) string {
	for i, attr := range f.attrs {
		if i != 0 {
			log += "\n"
		}
		for _, padding := range paddings {
			if i == 0 {
				padding.space = 0
				padding.char = ""
			}
			log += f.tree.Render().Sprintf("%-*s", padding.space, padding.char)
		}
		if i != 0 {
			if i+1 == len(f.attrs) {
				log += f.tree.Render().Sprint(f.bottomCurve)
			} else {
				log += f.tree.Render().Sprint(f.sideT)
			}
		}
		switch v := attr.Value.Any().(type) {
		case []slog.Attr:
			typeLen := 0
			if f.ShowType {
				log += f.typing.Render().Sprint("(group)")
				typeLen += 7
			}
			log += f.key.Render().Sprint(attr.Key)
			if len(v) == 1 {
				log += f.tree.Render().Sprint(f.dash)
			} else {
				log += f.tree.Render().Sprint(f.t)
			}
			var char string
			if i+1 == len(f.attrs) {
				char = ""
			} else {
				char = f.l
			}
			log += f.NewFields(v).RenderGroup("", append(paddings, Padding{space: 1 + typeLen + len(attr.Key), char: char})...)
		default:
			lines := []string{""}
			linesIdx := 0
			for i, c := range fmt.Sprint(v) {
				if (i+1)%f.MaxValueLen == 0 || c == '\n' {
					linesIdx += 1
					lines = append(lines, "")
					lines[linesIdx] += f.value.Render().Sprint(string(c))
				} else {
					lines[linesIdx] += f.value.Render().Sprint(string(c))
				}
			}
			typeLen := 0
			if f.ShowType {
				typeString := fmt.Sprintf("%T", v)
				log += f.typing.Render().Sprintf("(%T)", typeString)
				typeLen = len(typeString) + 2
			}
			log += f.key.Render().Sprint(attr.Key)
			if len(lines) == 1 {
				log += f.ktov.Render().Sprint(f.dottedDash)
				log += f.value.Render().Sprint(lines[0])
			} else {
				for i, line := range lines {
					if i != 0 {
						log += "\n"
						for _, padding := range append(paddings, Padding{space: 1 + typeLen + len(attr.Key), char: f.l}) {
							log += f.tree.Render().Sprintf("%-*s", padding.space, f.l)
						}
					}
					if i == 0 {
						log += f.ktov.Render().Sprint(f.topCorner)
					} else if i+1 == len(lines) {
						log += f.ktov.Render().Sprint(f.bottomCorner)
					} else {
						log += f.ktov.Render().Sprint(f.dottedL)
					}
					log += f.value.Render().Sprint(line)
				}
			}
		}
	}
	return log
}

func (h *Handler) ParseFields(r slog.Record) Fields {
	fields := Fields{
		Styling: h.Styling,
	}
	r.Attrs(func(a slog.Attr) bool {
		fields.attrs = append(fields.attrs, a)
		return true
	})
	return fields
}

func (f Fields) NewFields(attrs []slog.Attr) Fields {
	return Fields{
		attrs:   attrs,
		Styling: f.Styling,
	}
}

func Prefix(prefix string) slog.Attr {
	return slog.String("_696969_PREFIX", prefix)
}

func (f Fields) HasPrefix() (has bool) {
	for _, attr := range f.attrs {
		if attr.Key == "_696969_PREFIX" {
			return true
		}
	}
	return false
}

func (f *Fields) GetPrefix() string {
	var prefix string
	newAttrs := []slog.Attr{}
	for _, attr := range (*f).attrs {
		if attr.Key == "_696969_PREFIX" {
			prefix = fmt.Sprintf("%v", attr.Value)
		} else {
			newAttrs = append(newAttrs, attr)
		}
	}
	(*f).attrs = newAttrs
	return prefix
}
