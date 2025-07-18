package trial

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestIsZero confirms that IsZero behaves as expected.
func TestIsZero(t *testing.T) {
	t.Parallel()
	t.Run("ValidInt", func(t *testing.T) {
		t.Parallel()
		testCases := []struct {
			name     string
			action   func(vi *ValidInt)
			expected bool
		}{
			{
				name:     "Unset ValidInt should be zero value",
				action:   func(vi *ValidInt) {},
				expected: true,
			},
			{
				name:     "Set to 0 should NOT be zero value",
				action:   func(vi *ValidInt) { vi.Set(0) },
				expected: false,
			},
			{
				name:     "Set to non-zero should NOT be zero value",
				action:   func(vi *ValidInt) { vi.Set(42) },
				expected: false,
			},
			{
				name: "Reset should be zero value",
				action: func(vi *ValidInt) {
					vi.Set(100) // Set to a value first
					vi.Reset()
				},
				expected: true,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()
				var vi ValidInt
				tc.action(&vi)
				assert.Equal(t, tc.expected, vi.IsZero())
			})
		}
	})

	t.Run("ValidFloat", func(t *testing.T) {
		t.Parallel()
		testCases := []struct {
			name     string
			action   func(vf *ValidFloat)
			expected bool
		}{
			{
				name:     "Unset should be zero value",
				action:   func(vf *ValidFloat) {},
				expected: true,
			},
			{
				name:     "Set to 0.0 should NOT be zero value",
				action:   func(vf *ValidFloat) { vf.Set(0.0) },
				expected: false,
			},
			{
				name:     "Set to non-zero should NOT be zero value",
				action:   func(vf *ValidFloat) { vf.Set(42.5) },
				expected: false,
			},
			{
				name: "Reset should be zero value",
				action: func(vf *ValidFloat) {
					vf.Set(100.5)
					vf.Reset()
				},
				expected: true,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()
				var vf ValidFloat
				tc.action(&vf)
				assert.Equal(t, tc.expected, vf.IsZero())
			})
		}
	})

	t.Run("ValidString", func(t *testing.T) {
		t.Parallel()
		testCases := []struct {
			name     string
			action   func(vs *ValidString)
			expected bool
		}{
			{
				name:     "Unset should be zero value",
				action:   func(vs *ValidString) {},
				expected: true,
			},
			{
				name:     "Set to empty string should NOT be zero value",
				action:   func(vs *ValidString) { vs.Set("") },
				expected: false,
			},
			{
				name:     "Set to NA should be zero value",
				action:   func(vs *ValidString) { vs.Set("NA") },
				expected: true,
			},
			{
				name:     "Set to non-empty string should NOT be zero value",
				action:   func(vs *ValidString) { vs.Set("hello") },
				expected: false,
			},
			{
				name: "Reset should be zero value",
				action: func(vs *ValidString) {
					vs.Set("world")
					vs.Reset()
				},
				expected: true,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()
				var vs ValidString
				tc.action(&vs)
				assert.Equal(t, tc.expected, vs.IsZero())
			})
		}
	})
}

// TestMarshalJSON tests marshaling our validtypes with different JSON tags
func TestMarshalJSON(t *testing.T) {
	t.Run("ValidInt", func(t *testing.T) {
		type OmitzeroData struct {
			Field1 ValidInt `json:"FIELD1,omitzero"`
			Field2 ValidInt `json:"FIELD2,omitzero"`
			Field3 ValidInt `json:"FIELD3,omitzero"`
			Field4 ValidInt `json:"FIELD4,omitzero"`
		}
		type OmitemptyData struct {
			Field1 ValidInt `json:"FIELD1,omitempty"`
			Field2 ValidInt `json:"FIELD2,omitempty"`
			Field3 ValidInt `json:"FIELD3,omitempty"`
			Field4 ValidInt `json:"FIELD4,omitempty"`
		}
		testCases := []struct {
			name         string
			setupData    func() interface{}
			expectedJSON string
		}{
			{
				name: "omitzero tag",
				setupData: func() interface{} {
					sd := OmitzeroData{}
					sd.Field1.Set(1)
					sd.Field2.Set(0)
					sd.Field4.Set(5)
					// Field3 remains unset
					return sd
				},
				expectedJSON: `{"FIELD1":1,"FIELD2":0,"FIELD4":5}`,
			},
			{
				name: "omitempty tag",
				setupData: func() interface{} {
					sd := OmitemptyData{}
					sd.Field1.Set(1)
					sd.Field2.Set(0)
					sd.Field4.Set(5)
					// Field3 remains unset
					return sd
				},
				expectedJSON: `{"FIELD1":1,"FIELD2":0,"FIELD3":null,"FIELD4":5}`,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				data := tc.setupData()
				jsonData, err := json.Marshal(data)
				require.NoError(t, err, "Failed to marshal JSON")

				jsonStr := string(jsonData)
				t.Logf("Generated JSON: %s", jsonStr)
				assert.JSONEq(t, tc.expectedJSON, jsonStr)
			})
		}
	})

	t.Run("ValidFloat", func(t *testing.T) {
		type OmitzeroData struct {
			Field1 ValidFloat `json:"FIELD1,omitzero"`
			Field2 ValidFloat `json:"FIELD2,omitzero"`
			Field3 ValidFloat `json:"FIELD3,omitzero"`
			Field4 ValidFloat `json:"FIELD4,omitzero"`
		}
		type OmitemptyData struct {
			Field1 ValidFloat `json:"FIELD1,omitempty"`
			Field2 ValidFloat `json:"FIELD2,omitempty"`
			Field3 ValidFloat `json:"FIELD3,omitempty"`
			Field4 ValidFloat `json:"FIELD4,omitempty"`
		}
		testCases := []struct {
			name         string
			setupData    func() interface{}
			expectedJSON string
		}{
			{
				name: "omitzero tag",
				setupData: func() interface{} {
					sd := OmitzeroData{}
					sd.Field1.Set(1)
					sd.Field2.Set(0)
					sd.Field4.Set(5.2)
					// Field3 remains unset
					return sd
				},
				expectedJSON: `{"FIELD1":1,"FIELD2":0,"FIELD4":5.2}`,
			},
			{
				name: "omitempty tag",
				setupData: func() interface{} {
					sd := OmitemptyData{}
					sd.Field1.Set(1)
					sd.Field2.Set(0)
					sd.Field4.Set(5.2)
					// Field3 remains unset
					return sd
				},
				expectedJSON: `{"FIELD1":1,"FIELD2":0,"FIELD3":null,"FIELD4":5.2}`,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				data := tc.setupData()
				jsonData, err := json.Marshal(data)
				require.NoError(t, err, "Failed to marshal JSON")

				jsonStr := string(jsonData)
				t.Logf("Generated JSON: %s", jsonStr)
				assert.JSONEq(t, tc.expectedJSON, jsonStr)
			})
		}
	})

	t.Run("ValidString", func(t *testing.T) {
		type OmitzeroData struct {
			Field1 ValidString `json:"FIELD1,omitzero"`
			Field2 ValidString `json:"FIELD2,omitzero"`
			Field3 ValidString `json:"FIELD3,omitzero"`
			Field4 ValidString `json:"FIELD4,omitzero"`
		}
		type OmitemptyData struct {
			Field1 ValidString `json:"FIELD1,omitempty"`
			Field2 ValidString `json:"FIELD2,omitempty"`
			Field3 ValidString `json:"FIELD3,omitempty"`
			Field4 ValidString `json:"FIELD4,omitempty"`
		}
		testCases := []struct {
			name         string
			setupData    func() interface{}
			expectedJSON string
		}{
			{
				name: "omitzero tag",
				setupData: func() interface{} {
					sd := OmitzeroData{}
					sd.Field1.Set("hello")
					sd.Field2.Set("")   // empty string is valid
					sd.Field4.Set("NA") // this will be unset
					// Field3 remains unset
					return sd
				},
				expectedJSON: `{"FIELD1":"hello","FIELD2":""}`,
			},
			{
				name: "omitempty tag",
				setupData: func() interface{} {
					sd := OmitemptyData{}
					sd.Field1.Set("hello")
					sd.Field2.Set("")
					sd.Field4.Set("NA")
					// Field3 remains unset
					return sd
				},
				expectedJSON: `{"FIELD1":"hello","FIELD2":"","FIELD3":null,"FIELD4":null}`,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				data := tc.setupData()
				jsonData, err := json.Marshal(data)
				require.NoError(t, err, "Failed to marshal JSON")

				jsonStr := string(jsonData)
				t.Logf("Generated JSON: %s", jsonStr)
				assert.JSONEq(t, tc.expectedJSON, jsonStr)
			})
		}
	})
}

// TestUnmarshalJSON tests unmarshaling JSON into appropriate validtype fields
func TestUnmarshalJSON(t *testing.T) {
	t.Run("ValidInt", func(t *testing.T) {
		type SomeData struct {
			Field1 ValidInt `json:"FIELD1"`
			Field2 ValidInt `json:"FIELD2"`
			Field3 ValidInt `json:"FIELD3"`
			Field4 ValidInt `json:"FIELD4"`
		}

		testCases := []struct {
			name      string
			jsonData  string
			expected  SomeData
			expectErr bool
		}{
			{
				name:     "valid JSON",
				jsonData: `{"FIELD1":1,"FIELD2":0,"FIELD3":4,"FIELD4":5}`,
				expected: SomeData{
					Field1: NewValidInt(1),
					Field2: NewValidInt(0),
					Field3: NewValidInt(4),
					Field4: NewValidInt(5),
				},
				expectErr: false,
			},
			{
				name:     "JSON with missing fields",
				jsonData: `{"FIELD1":1,"FIELD4":5}`,
				expected: SomeData{
					Field1: NewValidInt(1),
					Field2: ValidInt{}, // Stays invalid (omitted)
					Field3: ValidInt{}, // Stays invalid (omitted)
					Field4: NewValidInt(5),
				},
				expectErr: false,
			},
			{
				name:     "JSON with null fields",
				jsonData: `{"FIELD1":null,"FIELD2":null,"FIELD3":null,"FIELD4":null}`,
				expected: SomeData{
					Field1: ValidInt{},
					Field2: ValidInt{}, // Stays invalid (omitted)
					Field3: ValidInt{}, // Stays invalid (omitted)
					Field4: ValidInt{},
				},
				expectErr: false,
			},
			{
				name:      "invalid JSON type for field (string instead of number)",
				jsonData:  `{"FIELD1":"not-a-number","FIELD2":0}`,
				expected:  SomeData{}, // Not checked on error
				expectErr: true,
			},
			{
				name:      "invalid JSON syntax",
				jsonData:  `{"FIELD1":1,"FIELD2":0,}`, // trailing comma is invalid
				expected:  SomeData{},                 // Not checked on error
				expectErr: true,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				actual := SomeData{}
				err := json.Unmarshal([]byte(tc.jsonData), &actual)

				if tc.expectErr {
					t.Logf("Reported Error: %s", err)
					assert.Error(t, err, "Expected error when unmarshaling invalid JSON")
				} else {
					require.NoError(t, err, "Failed to unmarshal valid JSON")
					assert.Equal(t, tc.expected, actual)
				}
			})
		}
	})

	t.Run("ValidFloat", func(t *testing.T) {
		type SomeData struct {
			Field1 ValidFloat `json:"FIELD1"`
			Field2 ValidFloat `json:"FIELD2"`
			Field3 ValidFloat `json:"FIELD3"`
			Field4 ValidFloat `json:"FIELD4"`
		}

		testCases := []struct {
			name      string
			jsonData  string
			expected  SomeData
			expectErr bool
		}{
			{
				name:     "valid JSON",
				jsonData: `{"FIELD1":1,"FIELD2":0,"FIELD3":4,"FIELD4":5.2}`,
				expected: SomeData{
					Field1: NewValidFloat(1),
					Field2: NewValidFloat(0),
					Field3: NewValidFloat(4),
					Field4: NewValidFloat(5.2),
				},
				expectErr: false,
			},
			{
				name:     "JSON with missing fields",
				jsonData: `{"FIELD1":1,"FIELD4":5.2}`,
				expected: SomeData{
					Field1: NewValidFloat(1),
					Field2: ValidFloat{}, // Stays invalid (omitted)
					Field3: ValidFloat{}, // Stays invalid (omitted)
					Field4: NewValidFloat(5.2),
				},
				expectErr: false,
			},
			{
				name:     "JSON with null fields",
				jsonData: `{"FIELD1":null,"FIELD2":null,"FIELD3":null,"FIELD4":null}`,
				expected: SomeData{
					Field1: ValidFloat{},
					Field2: ValidFloat{},
					Field3: ValidFloat{},
					Field4: ValidFloat{},
				},
				expectErr: false,
			},
			{
				name:      "invalid JSON type for field (string instead of number)",
				jsonData:  `{"FIELD1":"not-a-number","FIELD2":0}`,
				expected:  SomeData{}, // Not checked on error
				expectErr: true,
			},
			{
				name:      "invalid JSON syntax",
				jsonData:  `{"FIELD1":1,"FIELD2":0,}`, // trailing comma is invalid
				expected:  SomeData{},                 // Not checked on error
				expectErr: true,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				actual := SomeData{}
				err := json.Unmarshal([]byte(tc.jsonData), &actual)

				if tc.expectErr {
					t.Logf("Reported Error: %s", err)
					assert.Error(t, err, "Expected error when unmarshaling invalid JSON")
				} else {
					require.NoError(t, err, "Failed to unmarshal valid JSON")
					assert.Equal(t, tc.expected, actual)
				}
			})
		}
	})

	t.Run("ValidString", func(t *testing.T) {
		type SomeData struct {
			Field1 ValidString `json:"FIELD1"`
			Field2 ValidString `json:"FIELD2"`
			Field3 ValidString `json:"FIELD3"`
			Field4 ValidString `json:"FIELD4"`
		}

		testCases := []struct {
			name      string
			jsonData  string
			expected  SomeData
			expectErr bool
		}{
			{
				name:     "valid JSON with strings",
				jsonData: `{"FIELD1":"hello","FIELD2":"","FIELD3":"world","FIELD4":"NA"}`,
				expected: SomeData{
					Field1: NewValidString("hello"),
					Field2: NewValidString(""),
					Field3: NewValidString("world"),
					Field4: ValidString{}, // NA becomes invalid
				},
				expectErr: false,
			},
			{
				name:     "JSON with missing fields",
				jsonData: `{"FIELD1":"hello"}`,
				expected: SomeData{
					Field1: NewValidString("hello"),
					Field2: ValidString{},
					Field3: ValidString{},
					Field4: ValidString{},
				},
				expectErr: false,
			},
			{
				name:     "JSON with null fields",
				jsonData: `{"FIELD1":null,"FIELD2":null}`,
				expected: SomeData{
					Field1: ValidString{},
					Field2: ValidString{},
				},
				expectErr: false,
			},
			{
				name:      "invalid JSON type for field (number instead of string)",
				jsonData:  `{"FIELD1":123}`,
				expected:  SomeData{},
				expectErr: true,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				actual := SomeData{}
				err := json.Unmarshal([]byte(tc.jsonData), &actual)

				if tc.expectErr {
					t.Logf("Reported Error: %s", err)
					assert.Error(t, err, "Expected error when unmarshaling invalid JSON")
				} else {
					require.NoError(t, err, "Failed to unmarshal valid JSON")
					assert.Equal(t, tc.expected, actual)
				}
			})
		}
	})
}

// TestString tests the String() method of the validtypes
func TestString(t *testing.T) {
	t.Run("ValidInt", func(t *testing.T) {
		resetInt := NewValidInt(100)
		resetInt.Reset()

		testCases := []struct {
			name           string
			input          ValidInt
			expectedString string
		}{
			{
				name:           "unset ValidInt returns empty string",
				input:          ValidInt{}, // zero value
				expectedString: "",
			},
			{
				name:           "set to zero returns '0'",
				input:          NewValidInt(0),
				expectedString: "0",
			},
			{
				name:           "positive value",
				input:          NewValidInt(42),
				expectedString: "42",
			},
			{
				name:           "negative value",
				input:          NewValidInt(-123),
				expectedString: "-123",
			},
			{
				name:           "reset returns empty string",
				input:          resetInt,
				expectedString: "",
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				result := tc.input.String()
				assert.Equal(t, tc.expectedString, result, "String() returned unexpected value")
			})
		}
	})

	t.Run("ValidFloat", func(t *testing.T) {
		resetFloat := NewValidFloat(100.2)
		resetFloat.Reset()

		testCases := []struct {
			name           string
			input          ValidFloat
			expectedString string
		}{
			{
				name:           "unset ValidFloat returns empty string",
				input:          ValidFloat{}, // zero value
				expectedString: "",
			},
			{
				name:           "set to zero returns '0'",
				input:          NewValidFloat(0),
				expectedString: "0",
			},
			{
				name:           "positive value",
				input:          NewValidFloat(42),
				expectedString: "42",
			},
			{
				name:           "negative value",
				input:          NewValidFloat(-123),
				expectedString: "-123",
			},
			{
				name:           "Fractional value",
				input:          NewValidFloat(123.456789),
				expectedString: "123.456789",
			},
			{
				name:           "reset returns empty string",
				input:          resetFloat,
				expectedString: "",
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				result := tc.input.String()
				assert.Equal(t, tc.expectedString, result, "String() returned unexpected value")
			})
		}
	})

	t.Run("ValidString", func(t *testing.T) {
		resetString := NewValidString("something")
		resetString.Reset()

		testCases := []struct {
			name           string
			input          ValidString
			expectedString string
		}{
			{
				name:           "unset returns empty string",
				input:          ValidString{},
				expectedString: "",
			},
			{
				name:           "set to empty string returns empty string",
				input:          NewValidString(""),
				expectedString: "",
			},
			{
				name:           "set to NA returns empty string",
				input:          NewValidString("NA"),
				expectedString: "",
			},
			{
				name:           "set to a value returns the value",
				input:          NewValidString("hello"),
				expectedString: "hello",
			},
			{
				name:           "reset returns empty string",
				input:          resetString,
				expectedString: "",
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				result := tc.input.String()
				assert.Equal(t, tc.expectedString, result, "String() returned unexpected value")
			})
		}
	})
}

// TestNewValidType tests the NewValid* constructors
func TestNewValidType(t *testing.T) {
	t.Run("ValidInt", func(t *testing.T) {
		testCases := []struct {
			name          string
			value         int
			expectedValue int
			expectedValid bool
		}{
			{
				name:          "zero value",
				value:         0,
				expectedValue: 0,
				expectedValid: true,
			},
			{
				name:          "positive value",
				value:         42,
				expectedValue: 42,
				expectedValid: true,
			},
			{
				name:          "negative value",
				value:         -100,
				expectedValue: -100,
				expectedValid: true,
			},
			{
				name:          "max int",
				value:         2147483647, // Assuming 32-bit int
				expectedValue: 2147483647,
				expectedValid: true,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				vi := NewValidInt(tc.value)

				// Check struct fields
				assert.Equal(t, tc.expectedValue, vi.Value, "Value field mismatch")
				assert.Equal(t, tc.expectedValid, vi.Valid, "Valid field mismatch")

				// Verify using Get method as well
				value, valid := vi.Get()
				assert.Equal(t, tc.expectedValue, value, "Value from Get() mismatch")
				assert.Equal(t, tc.expectedValid, valid, "Valid from Get() mismatch")

				// IsZero should always be false for NewValidInt
				assert.False(t, vi.IsZero(), "IsZero() should be false")
			})
		}
	})

	t.Run("ValidFloat", func(t *testing.T) {
		testCases := []struct {
			name          string
			value         float64
			expectedValue float64
			expectedValid bool
		}{
			{
				name:          "zero value",
				value:         0,
				expectedValue: 0,
				expectedValid: true,
			},
			{
				name:          "positive value",
				value:         42,
				expectedValue: 42,
				expectedValid: true,
			},
			{
				name:          "negative value",
				value:         -100,
				expectedValue: -100,
				expectedValid: true,
			},
			{
				name:          "fractional value",
				value:         123.4567890,
				expectedValue: 123.456789,
				expectedValid: true,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				vi := NewValidFloat(tc.value)
				delta := 0.0000001

				// Check struct fields
				assert.InDelta(t, tc.expectedValue, vi.Value, delta, "Value field mismatch")
				assert.Equal(t, tc.expectedValid, vi.Valid, "Valid field mismatch")

				// Verify using Get method as well
				value, valid := vi.Get()
				assert.InDelta(t, tc.expectedValue, value, delta, "Value from Get() mismatch")
				assert.Equal(t, tc.expectedValid, valid, "Valid from Get() mismatch")

				// IsZero should always be false for NewValidFloat
				assert.False(t, vi.IsZero(), "IsZero() should be false")
			})
		}
	})

	t.Run("ValidString", func(t *testing.T) {
		testCases := []struct {
			name          string
			value         string
			expectedValue string
			expectedValid bool
		}{
			{
				name:          "non-empty string",
				value:         "hello",
				expectedValue: "hello",
				expectedValid: true,
			},
			{
				name:          "empty string",
				value:         "",
				expectedValue: "",
				expectedValid: true,
			},
			{
				name:          "NA string",
				value:         "NA",
				expectedValue: "", // Value should be empty for invalid
				expectedValid: false,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				vs := NewValidString(tc.value)

				assert.Equal(t, tc.expectedValue, vs.Value, "Value field mismatch")
				assert.Equal(t, tc.expectedValid, vs.Valid, "Valid field mismatch")

				value, valid := vs.Get()
				assert.Equal(t, tc.expectedValue, value, "Value from Get() mismatch")
				assert.Equal(t, tc.expectedValid, valid, "Valid from Get() mismatch")

				assert.Equal(t, !tc.expectedValid, vs.IsZero(), "IsZero() mismatch")
			})
		}
	})
}
