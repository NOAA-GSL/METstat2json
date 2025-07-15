package trial

import (
	"encoding/json"
	"strconv"
)

// ValidInt represents an int that tracks whether it has been explicitly set.
// The zero value of ValidInt is when Valid=false. ValidInt.IsZero is used as an easy
// way to detect that the ValidInt hasn't been explicitly set.
// Fields that are unset, can be omitted when marshalling to JSON by using the omitzero tag
type ValidInt struct {
	Value int
	Valid bool
}

// NewValidInt creates a ValidInt with a value.
func NewValidInt(value int) ValidInt {
	return ValidInt{Value: value, Valid: true}
}

// Set assigns a value and marks it as valid.
func (vi *ValidInt) Set(value int) {
	vi.Value = value
	vi.Valid = true
}

// Get returns the underlying value and validity.
func (vi ValidInt) Get() (int, bool) {
	return vi.Value, vi.Valid
}

// IsZero is used by omitzero, and returns true if the value is not valid.
func (vi ValidInt) IsZero() bool {
	return !vi.Valid
}

// Reset marks the value as invalid.
func (vi *ValidInt) Reset() {
	vi.Valid = false
	vi.Value = 0
}

// MarshalJSON implements json.Marshaler.
func (vi ValidInt) MarshalJSON() ([]byte, error) {
	if !vi.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(vi.Value)
}

// UnmarshalJSON implements json.Unmarshaler.
func (vi *ValidInt) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		vi.Valid = false
		vi.Value = 0
		return nil
	}

	var value int
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	vi.Value = value
	vi.Valid = true
	return nil
}

// String returns string representation.
func (vi ValidInt) String() string {
	if !vi.Valid {
		return ""
	}
	return strconv.Itoa(vi.Value)
}
