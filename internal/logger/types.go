package logger

import "github.com/rs/zerolog"

type Logger interface {
	Trace(msg string, fields ...Field)
	Debug(msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)
}

type Field struct {
	Key   string
	Value any
}

type zerologger struct {
	logger      *zerolog.Logger
	serviseName string
}
