package gomyslog

import (
	"context"
	"log/slog"
)

func (h *Handler) Handle(ctx context.Context, r slog.Record) error {
	var log string

	level := h.level[r.Level]
	// render time section
	if h.ShowTime {
		log += h.time.Bg().InvertGround().Render().Sprint(h.leftCap)
		log += h.time.Render().Sprint(r.Time.Format(" 2006/01/02 15:05:05.000 "))
	} else {
		log += level.Bg().InvertGround().Render().Sprint(h.leftCap)
	}
	// parse fields
	fields := h.ParseFields(r)
	if fields.HasPrefix() {
		prefix := fields.GetPrefix()
		// if prefix render
		if len(prefix) > 0 {
			log += h.time.Bg().InvertGround().Render().Add(h.prefix.Bg().ToAttrs()...).Sprint(h.rightArrow)
			log += h.prefix.Render().Sprintf(" %s ", prefix)
			log += h.prefix.Bg().InvertGround().Render().Add(level.Bg().ToAttrs()...).Sprint(h.rightArrow)
		} else {
			log += h.time.Bg().InvertGround().Render().Add(level.Bg().ToAttrs()...).Sprint(h.rightArrow)
		}
	}
	// render level section
	log += level.Render().Sprintf(" %s ", r.Level.String())
	log += level.Bg().InvertGround().Render().Add(h.msg.Bg().ToAttrs()...).Sprint(h.rightArrow)
	// render message section
	log += h.msg.Render().Sprintf(" %s ", r.Message)
	log += h.msg.Bg().InvertGround().Render().Sprint(h.rightArrow)
	// if fields exist add header
	if len(fields.attrs) > 0 {
		log = h.tree.Render().Sprint(h.topCurve+h.dash) + log
		// render fields
		log += fields.Render()
	}
	// print log
	h.l.Println(log)
	return nil
}
