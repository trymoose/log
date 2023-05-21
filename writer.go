package log

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

// WithWriter sets the writer of [Logger].
// The writer is shared with all children until [Logger.With] is called with [WithNewWriter].
// Each write to the writer is locked until the message completes, errors are ignored.
// Default value is [os.Stdout].
// If w is nil, [os.Stdout] will be used.
// Change the value of an existing logger with [WithNewWriter].
func WithWriter(w io.Writer) Arg { return &_WithWriter{w: w} }

// _Writer is a synced writer.
type _Writer struct {
	w io.Writer
	l sync.Mutex
}

// _NewWriter returns an initialized writer.
// If w == nil, [os.Stdout] is used.
func _NewWriter(w io.Writer) *_Writer {
	if w == nil {
		w = os.Stdout
	}
	return &_Writer{w: w}
}

// WriteString writes args to the output.
// The writer is locked before the message is written.
// Errors are ignored.
// Only args implementing [_PrintableArg] are printed.
// [_PrintableArg._Key] and [_PrintableArg._Value] are cleaned with [strings.TrimSpace].
// If the key is empty it is skipped.
// If there are multiple empty values, only the first one will be printed.
func (w *_Writer) WriteString(args []Arg) {
	out := make([]string, 0, len(args))
	hadBlank := map[string]bool{}
	for _, a := range args {
		if a, ok := a.(_PrintableArg); ok {
			key := strings.TrimSpace(a._Key())
			value := strings.TrimSpace(a._Value())

			if key == "" {
				continue
			} else if value == "" {
				if hadBlank[key] {
					continue
				}
				hadBlank[key] = true
			}

			out = append(out, fmt.Sprintf("%s=%s", key, value))
		}
	}

	w.l.Lock()
	defer w.l.Unlock()
	_, _ = fmt.Fprintf(w.w, "%s\n", strings.Join(out, " "))
}

func (*_Writer) _IsControlArg() {}
func (*_Writer) _IsArg()        {}

type _WithWriter struct{ w io.Writer }

func (*_WithWriter) _IsControlArg() {}
func (*_WithWriter) _IsArg()        {}
