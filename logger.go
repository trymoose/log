package log

import (
	"os"
	"time"
)

var (
	_DefaultWriter       = os.Stdout
	_DefaultLevel  Level = LevelInfo
)

// Option represent a top level option for the [Logger].
// Options can only be used with [New].
type Option func(*Logger)

// Logger can be used to print a structured log to an output.
// All logging is done to a single output.
type Logger struct {
	out   *_Writer
	level *_Level
	args  []Arg
}

// New creates a new logger with the following defaults:
//   - output: [os.Stdout]
//   - level: [LevelInfo]
func New(opt ...Arg) *Logger {
	var l Logger
	l.args = _HandleArgs(append([]Arg{
		WithLevel(_DefaultLevel),
		WithWriter(_DefaultWriter),
	},
		opt...,
	),
		_ArgHandlerSkip[_WithNewLevel](func(level _WithNewLevel) { l.level = _NewLevel(Level(level)) }),
		_ArgHandlerSkip[*_Level](func(level *_Level) { l.level = level }),
		_ArgHandlerSkip[*_WithWriter](func(w *_WithWriter) { l.out = _NewWriter(w.w) }),
		_ArgHandlerSkip[*_Writer](func(w *_Writer) { l.out = w }),
	)
	return &l
}

// With creates a sublogger with the given args embedded into each message of the child.
// This can be used to embed the url of a request for example.
func (l *Logger) With(args ...Arg) *Logger {
	return New(append(append(l.args,
		l.out,
		l.level,
	), args...)...)
}

func (l *Logger) SetLevel(level Level) { l.level.Set(level) }
func (l *Logger) GetLevel() Level      { return l.level.Get() }

// Log does the actual logging at the given level.
func (l *Logger) Log(level Level, args ...Arg) {
	now := time.Now()
	debug := false
	if !l.level.CanPrint(level) {
		return
	} else if l.level.CanPrint(LevelDebug) {
		debug = true
	}

	skip := 1
	out := l.out
	msg := Skip()
	prefix := make([]Arg, 0, len(l.args))

	args = _HandleArgs(append(l.args, args...),
		_ArgHandlerSkip[*_WithWriter](func(w *_WithWriter) { out = _NewWriter(w.w) }),
		_ArgHandlerSkip[_AddSkip](func(delta _AddSkip) { skip = delta.Add(skip) }),
		_ArgHandlerSkip[_Msg](func(m _Msg) { msg = m }),
		_ArgHandlerSkip[_Prefix](func(p _Prefix) { prefix = append(prefix, p) }),
		_ArgHandler[_Producer](func(p _Producer) Arg { return p() }),
		_ArgHandler[_IfDebug](func(id _IfDebug) Arg {
			if debug {
				return id()
			}
			return Skip()
		}),
	)

	debugArg := Skip()
	if debug {
		debugArg = Caller(DebugCallerKey, skip)
	}

	prefix = append([]Arg{debugArg, Time("time", now), KeyValue(LevelKey, level.String())}, append(prefix, msg)...)
	out.WriteString(append(prefix, args...))
}

// Debug logs a message at [LevelDebug].
func (l *Logger) Debug(args ...Arg) {
	l.Log(LevelDebug, append(args, AddSkip(1))...)
}

// Info logs a message at [LevelInfo].
func (l *Logger) Info(args ...Arg) {
	l.Log(LevelInfo, append(args, AddSkip(1))...)
}

// Warn logs a message at [LevelWarn].
func (l *Logger) Warn(args ...Arg) {
	l.Log(LevelWarn, append(args, AddSkip(1))...)
}

// Error logs a message at [LevelError].
func (l *Logger) Error(err error, args ...Arg) bool {
	if err != nil {
		l.Log(LevelError, append([]Arg{Err(err), AddSkip(1)}, args...)...)
		return true
	}
	return false
}
