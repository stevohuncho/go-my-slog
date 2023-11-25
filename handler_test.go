package gomyslog_test

import (
	"log/slog"
	"os"
	"testing"

	"github.com/stevohuncho/gomyslog"
)

func InitHandler() {
	handler := gomyslog.NewHandler(os.Stdout,
		gomyslog.SlogHandlerOpt(slog.HandlerOptions{
			Level: slog.LevelDebug,
		}),
	)
	slog.SetDefault(slog.New(handler))
}

func TestDebug(t *testing.T) {
	InitHandler()
	slog.Debug("executing database query")
}

func TestInfo(t *testing.T) {
	InitHandler()
	slog.Info("executing database query")
}

func TestWarn(t *testing.T) {
	InitHandler()
	slog.Warn("executing database query")
}

func TestError(t *testing.T) {
	slog.Error("failed to execute database query")
}

func TestPrefix(t *testing.T) {
	InitHandler()
	slog.Info("heres how you can use the prefix", gomyslog.Prefix("TEST"))
}

func TestGroup(t *testing.T) {
	InitHandler()
	slog.Info("Welcome to Go My Slog!", slog.String("use slog just like", "normal"), slog.Group("this package formats", slog.String("any", "string"), slog.String("you", "want"), slog.Int("package creators", 1)), slog.Bool("should you use it", true))
}

func TestSubGroups(t *testing.T) {

}

func TestLineWrap(t *testing.T) {

}

func TestHiding(t *testing.T) {

}

func TestStyleOverwritting(t *testing.T) {
	myCustomStyling := gomyslog.Styling{
		StylingColors: gomyslog.StylingColors{},
		StylingChars:  gomyslog.StylingChars{},
		ShowType:      false,
		ShowTime:      false,
		MaxValueLen:   100,
	}
	handler := gomyslog.NewHandler(os.Stdout,
		gomyslog.SlogHandlerOpt(slog.HandlerOptions{
			Level: slog.LevelDebug,
		}),
		gomyslog.StylingHandlerOpt(myCustomStyling),
	)
	slog.SetDefault(slog.New(handler))
}
