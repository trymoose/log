package number

import (
	"bytes"
	"errors"
	"fmt"
	"golang.org/x/exp/constraints"
)

// MarshalText converts a number to bytes.
func MarshalText[N constraints.Integer](n N) []byte { return []byte(fmt.Sprintf("%d", n)) }

// UnmarshalText unmarshals ASCII number bytes to a given numeric type.
// Behavior for when the number represented by b is > max N is undefined.
func UnmarshalText[N constraints.Integer](b []byte) (N, error) {
	var n N
	parsed, err := fmt.Fscanf(bytes.NewReader(b), "%d", &n)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal number: %w", err)
	} else if parsed != 1 {
		return 0, errors.New("did not parse correct number of items")
	}
	return n, nil
}
