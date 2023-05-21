package log

import (
	"errors"
	"io"
	"os"
	"reflect"
	"testing"
)

func TestLogger_Debug(t *testing.T) {
	type fields struct {
		out   *_Writer
		level *_Level
		args  []Arg
	}
	type args struct {
		msg  string
		args []Arg
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				out:   tt.fields.out,
				level: tt.fields.level,
				args:  tt.fields.args,
			}
			l.Debug(append([]Arg{Msg(tt.args.msg)}, tt.args.args...)...)
		})
	}
}

func TestLogger_Error(t *testing.T) {
	type fields struct {
		out   io.Writer
		level Level
		args  []Arg
	}
	type args struct {
		err  error
		args []Arg
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "with an error",
			fields: fields{
				out:   os.Stderr,
				level: LevelError,
				args:  []Arg{Prefix("logger test")},
			},
			args: args{
				err:  errors.New("err: hello world"),
				args: []Arg{Bool("cool", true)},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New(append(tt.fields.args, WithWriter(tt.fields.out), WithLevel(tt.fields.level))...)
			l.Error(tt.args.err, tt.args.args...)
		})
	}
}

func TestLogger_GetLevel(t *testing.T) {
	type fields struct {
		out   *_Writer
		level *_Level
		args  []Arg
	}
	tests := []struct {
		name   string
		fields fields
		want   Level
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				out:   tt.fields.out,
				level: tt.fields.level,
				args:  tt.fields.args,
			}
			if got := l.GetLevel(); got != tt.want {
				t.Errorf("GetLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogger_Info(t *testing.T) {
	type fields struct {
		out   io.Writer
		level Level
		args  []Arg
	}
	type args struct {
		msg  string
		args []Arg
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "print anything",
			fields: fields{
				out:   os.Stdout,
				level: LevelInfo,
				args:  nil,
			},
			args: args{
				msg:  "hello world",
				args: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New(append(tt.fields.args, WithWriter(tt.fields.out), WithLevel(tt.fields.level))...)
			l.Info(append([]Arg{Msg(tt.args.msg)}, tt.args.args...)...)
		})
	}
}

func TestLogger_Log(t *testing.T) {
	type fields struct {
		out   *_Writer
		level *_Level
		args  []Arg
	}
	type args struct {
		level Level
		args  []Arg
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				out:   tt.fields.out,
				level: tt.fields.level,
				args:  tt.fields.args,
			}
			l.Log(tt.args.level, tt.args.args...)
		})
	}
}

func TestLogger_SetLevel(t *testing.T) {
	type fields struct {
		out   *_Writer
		level *_Level
		args  []Arg
	}
	type args struct {
		level Level
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				out:   tt.fields.out,
				level: tt.fields.level,
				args:  tt.fields.args,
			}
			l.SetLevel(tt.args.level)
		})
	}
}

func TestLogger_Warn(t *testing.T) {
	type fields struct {
		out   *_Writer
		level *_Level
		args  []Arg
	}
	type args struct {
		msg  string
		args []Arg
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				out:   tt.fields.out,
				level: tt.fields.level,
				args:  tt.fields.args,
			}
			l.Warn(append([]Arg{Msg(tt.args.msg)}, tt.args.args...)...)
		})
	}
}

func TestLogger_With(t *testing.T) {
	type fields struct {
		out   *_Writer
		level *_Level
		args  []Arg
	}
	type args struct {
		args []Arg
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Logger
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				out:   tt.fields.out,
				level: tt.fields.level,
				args:  tt.fields.args,
			}
			if got := l.With(tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("With() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		opt []Arg
	}
	tests := []struct {
		name string
		args args
		want *Logger
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.opt...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
