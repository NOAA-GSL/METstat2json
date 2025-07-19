package validtypes

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
	if !vi.Valid { // Only called when the omitzero tag isn't used
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

// UnmarshalText sets ValidInt from text values (i.e. stat file input)
// It treats "NA" and empty strings as a null/invalid value.
func (vi *ValidInt) UnmarshalText(text []byte) error {
	s := string(text)
	if s == "" || s == "NA" {
		vi.Reset()
		return nil
	}

	val, err := strconv.Atoi(s)
	if err != nil {
		vi.Reset()
		return err
	}

	vi.Set(val)
	return nil
}

// String returns string representation.
func (vi ValidInt) String() string {
	if !vi.Valid {
		return ""
	}
	return strconv.Itoa(vi.Value)
}

// ValidFloat represents a float64 that tracks whether it has been explicitly set.
// The zero value of ValidFloat is when Valid=false. ValidFloat.IsZero is used as an easy
// way to detect that the ValidFloat hasn't been explicitly set.
// Fields that are unset, can be omitted when marshalling to JSON by using the omitzero tag
type ValidFloat struct {
	Value float64
	Valid bool
}

// NewValidFloat creates a ValidFloat with a value.
func NewValidFloat(value float64) ValidFloat {
	return ValidFloat{Value: value, Valid: true}
}

// Set assigns a value and marks it as valid.
func (vf *ValidFloat) Set(value float64) {
	vf.Value = value
	vf.Valid = true
}

// Get returns the underlying value and validity.
func (vf ValidFloat) Get() (float64, bool) {
	return vf.Value, vf.Valid
}

// IsZero is used by omitzero, and returns true if the value is not valid.
func (vf ValidFloat) IsZero() bool {
	return !vf.Valid
}

// Reset marks the value as invalid.
func (vf *ValidFloat) Reset() {
	vf.Valid = false
	vf.Value = 0.0
}

// MarshalJSON implements json.Marshaler.
func (vf ValidFloat) MarshalJSON() ([]byte, error) {
	if !vf.Valid { // Only called when the omitzero tag isn't used
		return []byte("null"), nil
	}

	return json.Marshal(vf.Value)
}

// UnmarshalJSON implements json.Unmarshaler.
func (vf *ValidFloat) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		vf.Valid = false
		vf.Value = 0.0
		return nil
	}

	var value float64
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	vf.Value = value
	vf.Valid = true
	return nil
}

// UnmarshalText sets ValidFloat from text values (i.e. stat file input)
// It treats "NA" and empty strings as a null/invalid value.
func (vf *ValidFloat) UnmarshalText(text []byte) error {
	s := string(text)
	if s == "" || s == "NA" {
		vf.Reset()
		return nil
	}

	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		vf.Reset()
		return err
	}

	vf.Set(val)
	return nil
}

// String returns string representation.
func (vf ValidFloat) String() string {
	if !vf.Valid {
		return ""
	}
	return strconv.FormatFloat(vf.Value, 'f', -1, 64)
}

// ValidString represents a string that tracks whether it has been explicitly set.
// The zero value of ValidString is when Valid=false. ValidString.IsZero is used as an easy
// way to detect that the ValidString hasn't been explicitly set.
// It treats "NA" as an invalid/unset value.
// Fields that are unset can be omitted when marshalling to JSON by using the omitzero tag.
type ValidString struct {
	Value string
	Valid bool
}

// NewValidString creates a ValidString with a value.
// It treats "NA" as an invalid value.
func NewValidString(value string) ValidString {
	if value == "NA" {
		return ValidString{}
	}
	return ValidString{Value: value, Valid: true}
}

// Set assigns a value and marks it as valid.
// It treats "NA" as a signal to reset the value to its invalid state.
func (vs *ValidString) Set(value string) {
	if value == "NA" {
		vs.Reset()
		return
	}
	vs.Value = value
	vs.Valid = true
}

// Get returns the underlying value and validity.
func (vs ValidString) Get() (string, bool) {
	return vs.Value, vs.Valid
}

// IsZero is used by omitzero, and returns true if the value is not valid.
func (vs ValidString) IsZero() bool {
	return !vs.Valid
}

// Reset marks the value as invalid.
func (vs *ValidString) Reset() {
	vs.Valid = false
	vs.Value = ""
}

// MarshalJSON implements json.Marshaler.
func (vs ValidString) MarshalJSON() ([]byte, error) {
	if !vs.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(vs.Value)
}

// UnmarshalJSON implements json.Unmarshaler.
// It treats a JSON string "NA" as a null/invalid value.
func (vs *ValidString) UnmarshalJSON(data []byte) error {
	s := string(data)
	if s == "null" || s == `"NA"` {
		vs.Reset()
		return nil
	}

	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	vs.Value = value
	vs.Valid = true
	return nil
}

// UnmarshalText sets ValidString from text values (i.e. stat file input)
// It treats "NA" and empty strings as a null/invalid value.
func (vs *ValidString) UnmarshalText(text []byte) error {
	s := string(text)
	if s == "" || s == "NA" {
		vs.Reset()
		return nil
	}

	vs.Set(s)
	return nil
}

// String returns string representation.
func (vs ValidString) String() string {
	if !vs.Valid {
		return ""
	}
	return vs.Value
}
