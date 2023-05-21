package log

import (
	"sync/atomic"
)

var _G atomic.Pointer[Logger]

func init() { Reset() }

func Add(args ...Arg)                   { _G.Store(_G.Load().With(args...)) }
func With(args ...Arg) *Logger          { return _G.Load().With(args...) }
func Reset()                            { _G.Store(New()) }
func GetLevel() Level                   { return _G.Load().GetLevel() }
func SetLevel(level Level)              { _G.Load().SetLevel(level) }
func Log(level Level, args ...Arg)      { _G.Load().Log(level, append(args, AddSkip(1))...) }
func Info(args ...Arg)                  { _G.Load().Info(append(args, AddSkip(1))...) }
func Debug(args ...Arg)                 { _G.Load().Debug(append(args, AddSkip(1))...) }
func Warn(args ...Arg)                  { _G.Load().Warn(append(args, AddSkip(1))...) }
func Error(err error, args ...Arg) bool { return _G.Load().Error(err, append(args, AddSkip(1))...) }
