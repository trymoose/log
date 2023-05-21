package log

import (
	"sync"
)

// Level is the level to print log messages at.
type Level uint8

const LevelKey = "level"

const (
	// LevelAll prints at all levels.
	LevelAll Level = iota
	// LevelDebug should be only used in dev.
	// Adds file and line number of log call to output.
	LevelDebug
	// LevelInfo verbose logging.
	LevelInfo
	// LevelWarn only warning and errors are printed.
	LevelWarn
	// LevelError only errors are printed.
	LevelError
	// LevelNone nothing is printed.
	LevelNone
)

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelError:
		return "ERROR"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	default:
		return "UNKNOWN"
	}
}

// WithLevel creates a new shared level for this logger and child loggers until [WithLevel] is used again with [Logger.With].
func WithLevel(level Level) Arg { return _WithNewLevel(level) }

type _WithNewLevel Level

func (_WithNewLevel) _IsControlArg() {}
func (_WithNewLevel) _IsArg()        {}

type _Level struct {
	level Level
	l     sync.Mutex
}

func _NewLevel(level Level) *_Level { return &_Level{level: level} }

func (*_Level) _IsControlArg() {}
func (*_Level) _IsArg()        {}

func (l *_Level) Set(level Level) {
	l.l.Lock()
	defer l.l.Unlock()
	l.level = level
}

func (l *_Level) Get() Level {
	l.l.Lock()
	defer l.l.Unlock()
	return l.level
}

func (l *_Level) CanPrint(test Level) bool {
	l.l.Lock()
	defer l.l.Unlock()
	logLevel := l.level
	return test >= logLevel
}
