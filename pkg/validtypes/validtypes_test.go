package trial

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestValidIntIsZero confirms that IsZero behaves as expected.
func TestValidInt_IsZero(t *testing.T) {
	testCases := []struct {
		name     string
		action   func(vi *ValidInt)
		expected bool
	}{
		{
			name:     "Unset ValidInt should be zero",
			action:   func(vi *ValidInt) {},
			expected: true,
		},
		{
			name:     "Set to 0 should NOT be zero",
			action:   func(vi *ValidInt) { vi.Set(0) },
			expected: false,
		},
		{
			name:     "Set to non-zero should NOT be zero",
			action:   func(vi *ValidInt) { vi.Set(42) },
			expected: false,
		},
		{
			name: "Reset should be zero",
			action: func(vi *ValidInt) {
				vi.Set(100) // Set to a value first
				vi.Reset()
			},
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var vi ValidInt
			tc.action(&vi)
			assert.Equal(t, tc.expected, vi.IsZero())
		})
	}
}

// TestValidIntMarshal tests marshaling ValidInt with different JSON tags
func TestValidInt_MarshalJSON(t *testing.T) {
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
}

// TestValidIntUnmarshal tests unmarshaling JSON into ValidInt fields
func TestValidInt_UnmarshalJSON(t *testing.T) {
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
}

// TestValidIntString tests the String() method of ValidInt
func TestValidInt_String(t *testing.T) {
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
}

// TestNewValidInt tests the NewValidInt constructor
func Test_NewValidInt(t *testing.T) {
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

			// Check Value field
			assert.Equal(t, tc.expectedValue, vi.Value, "Value field mismatch")

			// Check Valid field
			assert.Equal(t, tc.expectedValid, vi.Valid, "Valid field mismatch")

			// Verify using Get method as well
			value, valid := vi.Get()
			assert.Equal(t, tc.expectedValue, value, "Value from Get() mismatch")
			assert.Equal(t, tc.expectedValid, valid, "Valid from Get() mismatch")

			// IsZero should always be false for NewValidInt
			assert.False(t, vi.IsZero(), "IsZero() should be false")
		})
	}
}
