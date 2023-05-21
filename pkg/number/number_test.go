package number

import (
	"errors"
	"golang.org/x/exp/constraints"
	"testing"
)

func TestUnmarshalText(t *testing.T) {
	_TestUnmarshalText[int32](t, "positive int", "123", 123, false, nil)
	_TestUnmarshalText[int32](t, "positive int with sign", "+123", 123, false, nil)
	_TestUnmarshalText[int32](t, "negative int", "-123", -123, false, nil)
	_TestUnmarshalText[uint8](t, "overflow int", "256", 0, true, nil)

	_TestUnmarshalText[uint32](t, "uint", "321", 321, false, nil)
	_TestUnmarshalText[uint32](t, "uint with sign", "+123", 0, true, nil)
	_TestUnmarshalText[uint32](t, "negative uint", "-123", 0, true, nil)
	_TestUnmarshalText[uint8](t, "overflow uint", "300", 0, true, nil)
}

func _TestUnmarshalText[N constraints.Integer](t *testing.T, name string, args string, want N, anyErr bool, wantErr error) {
	t.Helper()
	t.Run(name, func(t *testing.T) {
		got, err := UnmarshalText[N]([]byte(args))
		if err != nil {
			if anyErr {
				return
			} else if !errors.Is(err, wantErr) {
				t.Errorf("func=UnmarshalText error=%v wantAnyErr=%t wantErr=%v", err, anyErr, wantErr)
			}
			t.Errorf("func=UnmarshalText error=%v wantAnyErr=%t wantErr=%v", err, anyErr, wantErr)
		} else if got != want {
			t.Errorf("func=UnmarshalText got=%d, want=%d", got, want)
		}
	})
}
