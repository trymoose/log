package log

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"runtime"
	"strconv"
	"strings"
	"time"
	_ "unsafe"
)

type (
	// Arg is something that can be printed in a log message or used to control its output.
	Arg interface{ _IsArg() }

	_ControlArg interface {
		Arg
		_IsControlArg()
	}
	_PrintableArg interface {
		Arg
		_Key() string
		_Value() string
	}
)

type _blank struct{}

func (*_blank) _IsArg()        {}
func (*_blank) _IsControlArg() {}

var _Skip = &_blank{}

// Skip is an [Arg] to not print anything.
func Skip() Arg { return _Skip }

type _KeyValue [2]string

// KeyValue is a raw key=value string pair.
func KeyValue(key, value string) Arg { return _KeyValue{key, value} }
func (_KeyValue) _IsArg()            {}
func (kv _KeyValue) _Key() string    { return kv[0] }
func (kv _KeyValue) _Value() string  { return kv[1] }

type _Msg string

const MsgKey = "msg"

// Msg is a trimmed and quoted message with the key [MsgKey].
// Only the most recent given Msg will be used.
func Msg(format string, i ...any) Arg {
	if len(i) > 0 {
		format = fmt.Sprintf(format, i...)
	}
	return _Msg(strconv.Quote(strings.TrimSpace(format)))
}

func (_Msg) _IsArg()          {}
func (_Msg) _Key() string     { return MsgKey }
func (m _Msg) _Value() string { return string(m) }

const ErrorKey = "error"

// Err creates an [Arg] with the key [ErrorKey].
// If err == nil [Skip] is returned.
func Err(err error) Arg {
	if err != nil {
		return KeyValue(ErrorKey, strconv.Quote(strings.TrimSpace(err.Error())))
	}
	return Skip()
}

type _AddSkip int

// AddSkip increments the skip value of a call to a log function.
// The log functions only call skip with the key [DebugCallerKey] when the [Level] [LevelDebug] is valid.
func AddSkip(skip int) Arg          { return _AddSkip(skip) }
func (_AddSkip) _IsArg()            {}
func (_AddSkip) _IsControlArg()     {}
func (s _AddSkip) Add(skip int) int { return int(s) + skip }

const DebugCallerKey = "loc"

// Caller creates a caller [Arg] with the given skip.
// skip == 0 represents the caller of [Caller].
func Caller(key string, skip int) Arg {
	_, fn, ln, ok := runtime.Caller(skip + 1)
	if ok {
		return KeyValue(key, fmt.Sprintf("{ %s:%d }", fn, ln))
	}
	return Skip()
}

// String creates a key=value pair with value being a string quoted with [strconv.Quote].
func String(key, value string) Arg { return KeyValue(key, fmt.Sprintf("%q", value)) }

// Int is a key=value pair with a given integer value.
func Int[N constraints.Integer](key string, value N) Arg {
	return KeyValue(key, fmt.Sprintf("%d", value))
}

// Float is a key=value pair with a given float value.
func Float[N constraints.Float](key string, value N) Arg {
	return KeyValue(key, fmt.Sprintf("%f", value))
}

// Bool is a key=value pair with a given boolean value.
func Bool(key string, value bool) Arg { return KeyValue(key, fmt.Sprintf("%t", value)) }

// Duration is a key=value pair with a value as the unquoted output of [time.Duration.String].
func Duration(key string, v time.Duration) Arg { return KeyValue(key, v.String()) }

// Time is a key=value pair with a value as the unquoted output of [time.Time.Format] with the format of [time.RFC3339].
func Time(key string, v time.Time) Arg { return KeyValue(key, v.Format(time.RFC3339)) }

type _Prefix string

const PrefixKey = "prefix"

func Prefix(prefix string) Arg {
	prefix = strings.TrimSpace(prefix)
	if prefix == "" {
		return Skip()
	}
	prefix = strings.ReplaceAll(prefix, " ", "_")
	return _Prefix(strings.ToUpper(prefix))
}

func (_Prefix) _IsArg()          {}
func (_Prefix) _Key() string     { return PrefixKey }
func (p _Prefix) _Value() string { return string(p) }

type _Producer func() Arg

// Producer creates an [Arg] at log time.
func Producer(fn func() Arg) Arg { return _Producer(fn) }
func (_Producer) _IsArg()        {}
func (_Producer) _IsControl()    {}
func (p _Producer) New() Arg {
	if p != nil {
		return p()
	}
	return Skip()
}

type _IfDebug _Producer

// IfDebug is for a func that produces an [Arg] when [LevelDebug] is valid.
// producer is only called if [LevelDebug] is valid.
func IfDebug(producer _Producer) Arg { return _IfDebug(producer) }
func (_IfDebug) _IsArg()             {}
func (_IfDebug) _IsControl()         {}
func (id _IfDebug) Arg() Arg {
	if id != nil {
		return id()
	}
	return Skip()
}
