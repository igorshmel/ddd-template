package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
)

// Logger is interface of logger
type Logger interface {
	Info(format string, v ...interface{})
	Warn(format string, v ...interface{})
	Error(format string, v ...interface{})
	Fatal(format string, v ...interface{})
	Panic(format string, v ...interface{})
	Debug(format string, v ...interface{})
	WithPrefix(prefix string) Logger
	WithRequestID(requestID string) Logger
	WithMethod(method string) Logger
}

// Log is struct for logger
type Log struct {
	zlog      *zerolog.Logger
	requestID string
	method    string
}

// New is constructor for the logger
func New(isDebug ...bool) Log {
	logLevel := zerolog.InfoLevel

	var debug bool
	if isDebug != nil {
		if len(isDebug) != 0 {
			debug = isDebug[0]
		}
	}

	if debug {
		logLevel = zerolog.DebugLevel
	}

	writer := zerolog.ConsoleWriter{
		Out:        os.Stderr,
		NoColor:    true,
		TimeFormat: "2006/01/02 - 15:04:05.000000",
		PartsOrder: []string{
			zerolog.LevelFieldName,
			"request_id",
			zerolog.TimestampFieldName,
			"method",
			zerolog.MessageFieldName,
		},
		FieldsExclude: []string{
			"request_id",
			"method",
		},
		FormatLevel: func(i interface{}) string {
			if l, ok := i.(string); ok {
				switch l {
				case zerolog.LevelTraceValue:
					return "[TRC]"
				case zerolog.LevelDebugValue:
					return "[DBG]"
				case zerolog.LevelInfoValue:
					return "[INF]"
				case zerolog.LevelWarnValue:
					return "[WRN]"
				case zerolog.LevelErrorValue:
					return "[ERR]"
				case zerolog.LevelFatalValue:
					return "[FTL]"
				case zerolog.LevelPanicValue:
					return "[PNC]"
				default:
					return "[???]"
				}
			} else {
				return "[???]"
			}
		},
		FormatMessage: func(i interface{}) string {
			return fmt.Sprintf("| %s", i)
		},
	}

	zerolog.SetGlobalLevel(logLevel)
	logger := zerolog.New(writer).With().Timestamp().Logger()

	return Log{zlog: &logger}
}

// WithPrefix setting a prefix for log message
func (l Log) WithPrefix(prefix string) Logger {
	l.requestID = prefix
	return l
}

// WithRequestID setting a request id for log message
func (l Log) WithRequestID(reqID string) Logger {
	l.requestID = reqID
	return l
}

// WithMethod setting a method for log message
func (l Log) WithMethod(method string) Logger {
	l.method = method
	return l
}

// Info logs a message with info level
func (l Log) Info(format string, v ...interface{}) {
	l.prepareMessage(l.zlog.Info(), format, v...)
}

// Warn logs a message with warn level
func (l Log) Warn(format string, v ...interface{}) {
	l.prepareMessage(l.zlog.Warn(), format, v...)
}

// Warn logs a message with warn level
func (l Log) Error(format string, v ...interface{}) {
	l.prepareMessage(l.zlog.Error(), format, v...)
}

// Fatal logs a message with fatal level
func (l Log) Fatal(format string, v ...interface{}) {
	l.prepareMessage(l.zlog.Fatal(), format, v...)
}

// Panic logs a message with panic level
func (l Log) Panic(format string, v ...interface{}) {
	l.prepareMessage(l.zlog.Panic(), format, v...)
}

// Debug logs a message with debug level
func (l Log) Debug(format string, v ...interface{}) {
	l.prepareMessage(l.zlog.Debug(), format, v...)
}

// prepareMessage prepares the message for log
func (l Log) prepareMessage(event *zerolog.Event, format string, v ...interface{}) {
	if len(l.method) != 0 {
		event.Str("method", fmt.Sprintf("| %s", l.method))
	} else {
		event.Str("method", "")
	}

	if len(l.requestID) != 0 {
		event.Str("request_id", fmt.Sprintf("[%s]", l.requestID))
	} else {
		event.Str("request_id", "")
	}

	if v != nil {
		event.Msgf(format, v...)
		return
	}

	event.Msg(format)
}
