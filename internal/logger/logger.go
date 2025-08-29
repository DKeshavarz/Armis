package logger

import (
	"os"

	"github.com/rs/zerolog"
)

const (
	PADSTR = " "
	MESSAGE_SIZE = 40
)

func New(serviseName string)Logger{
	writer := zerolog.ConsoleWriter{
        Out: os.Stderr, 
        TimeFormat: "2006-01-02 15:04:05",
    }
	// writer.FormatCaller = func(i interface{}) string {
	// 	if i == nil {
	// 		return padRight("-", " ", 20)
	// 	}
	// 	return padRight(i.(string), " ", 20)
	// }
	logger := zerolog.New(writer).
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
	e.Msg(padRight(msg, PADSTR, MESSAGE_SIZE))
}

func (l *zerologger) Debug(msg string, fields ...Field) {
	e := l.logger.Debug()
	for _, f := range fields {
		e = e.Interface(f.Key, f.Value)
	}
	e.Msg(padRight(msg, PADSTR, MESSAGE_SIZE))
}

func (l *zerologger) Info(msg string, fields ...Field) {
	e := l.logger.Info()
	// for _, f := range fields {
	// 	e = e.Interface(f.Key, f.Value)
	// }
	for i := len(fields) - 1 ; i >= 0 ; i-- {
		e = e.Interface(fields[i].Key, fields[i].Value)
	}
	e.Msg(padRight(msg, PADSTR, MESSAGE_SIZE))
}

func (l *zerologger) Warn(msg string, fields ...Field) {
	e := l.logger.Warn()
	for _, f := range fields {
		e = e.Interface(f.Key, f.Value)
	}
	e.Msg(padRight(msg, PADSTR, MESSAGE_SIZE))
}

func (l *zerologger) Error(msg string, fields ...Field) {
	e := l.logger.Error()
	for _, f := range fields {
		e = e.Interface(f.Key, f.Value)
	}
	e.Msg(padRight(msg, PADSTR, MESSAGE_SIZE))
}

// ******************** helpers *******************************

func padRight(str, pad string, length int) string {
	for len(str) < length {
		str += pad
	}
	if len(str) > length {
		str = str[:length]
	}
	return str
}