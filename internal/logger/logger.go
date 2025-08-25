package logger

import (
	"os"

	"github.com/rs/zerolog"
)


func New(serviseName string)Logger{
	logger := zerolog.New(zerolog.ConsoleWriter{
        Out: os.Stderr, 
        TimeFormat: "2006-01-02 15:04:05",
    }).
    Level(zerolog.TraceLevel).
    With().
	Str("servise", serviseName).
    Timestamp().
	CallerWithSkipFrameCount(3).
    Logger()


	return &zerologger{
		logger: &logger,
		serviseName: serviseName,
	}
}

func (l *zerologger) Trace(msg string, fields ...Field) {
	e := l.logger.Trace()
	for _, f := range fields {
		e = e.Interface(f.Key, f.Value)
	}
	e.Caller(1).Str("servise: ", l.serviseName).Msg(msg)
}

func (l *zerologger) Debug(msg string, fields ...Field) {
	e := l.logger.Debug()
	for _, f := range fields {
		e = e.Interface(f.Key, f.Value)
	}
	e.Msg(msg)
}

func (l *zerologger) Info(msg string, fields ...Field) {
	e := l.logger.Info()
	for _, f := range fields {
		e = e.Interface(f.Key, f.Value)
	}
	e.Msg(msg)
}

func (l *zerologger) Warn(msg string, fields ...Field) {
	e := l.logger.Warn()
	for _, f := range fields {
		e = e.Interface(f.Key, f.Value)
	}
	e.Msg(msg)
}

func (l *zerologger) Error(msg string, fields ...Field) {
	e := l.logger.Error()
	for _, f := range fields {
		e = e.Interface(f.Key, f.Value)
	}
	e.Msg(msg)
}