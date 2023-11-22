package log

import (
	"context"
	"log/slog"
)

func (h *Handler) Handle(ctx context.Context, r slog.Record) error {
	var log string

	level := h.level[r.Level]
	// render time section
	if h.showTime {
		log += h.time.Bg().InvertGround().Render().Sprint(h.leftCap)
		log += h.time.Bg().Render().Add(h.time.Fg().ToAttrs()...).Sprint(r.Time.Format(" 2006/01/02 15:05:05.000 "))
		log += h.time.Bg().InvertGround().Render().Add(level.Bg().ToAttrs()...).Sprint(h.rightArrow)
	} else {
		log += level.Bg().InvertGround().Render().Sprint(h.leftCap)
	}
	// render level section
	log += level.Bg().Render().Add(level.Fg().ToAttrs()...).Sprintf(" %s ", r.Level.String())
	log += level.Bg().InvertGround().Render().Add(h.msg.Bg().ToAttrs()...).Sprint(h.rightArrow)
	// render message section
	log += h.msg.Bg().Render().Add(h.msg.Fg().ToAttrs()...).Sprintf(" %s ", r.Message)
	log += h.msg.Bg().InvertGround().Render().Sprint(h.rightArrow)
	// parse fields
	fields := h.ParseFields(r)
	if fields.HasPrefix() {
		prefix := fields.GetPrefix()
		_ = prefix
	}
	// if fields exist add header
	if len(fields.attrs) > 0 {
		log = h.msg.Fg().Render().Sprint(h.topCurve+h.dash) + log
		// render fields
		log += fields.Render("")
	}
	// print log
	h.l.Println(log)
	return nil
}

func Prefix(prefix string) slog.Attr {
	return slog.String("_696969_PREFIX", prefix)
}
