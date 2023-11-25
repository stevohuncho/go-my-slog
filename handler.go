package gomyslog

import (
	"io"
	"log"
	"log/slog"
	"os"
)

type Handler struct {
	slog.Handler
	l *log.Logger
	Styling
	*HandlerOpts
}

type HandlerOpts struct {
	slogOpts slog.HandlerOptions
	outFile  *os.File
	styling  Styling
}

type HandlerOptFunc func(*HandlerOpts)

func DefaultHandlerOpts() *HandlerOpts {
	return &HandlerOpts{
		slogOpts: slog.HandlerOptions{},
		outFile:  nil,
		styling:  DefaultStyling(),
	}
}

func SlogHandlerOpt(slogOpts slog.HandlerOptions) HandlerOptFunc {
	return func(opts *HandlerOpts) {
		opts.slogOpts = slogOpts
	}
}

func OutFileHandlerOpt(outfile *os.File) HandlerOptFunc {
	return func(opts *HandlerOpts) {
		opts.outFile = outfile
	}
}

func StylingHandlerOpt(styling Styling) HandlerOptFunc {
	return func(opts *HandlerOpts) {
		opts.styling = styling
	}
}

func NewHandler(out io.Writer, opts ...HandlerOptFunc) (h *Handler) {
	h = &Handler{}
	hOpts := DefaultHandlerOpts()
	for _, opt := range opts {
		opt(hOpts)
	}
	h.HandlerOpts = hOpts
	h.Handler = slog.NewJSONHandler(out, &hOpts.slogOpts)
	h.Styling = hOpts.styling
	h.l = log.New(out, "", 0)
	return h
}
