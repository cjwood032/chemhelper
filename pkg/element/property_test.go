package element

import "testing"

func TestConvertMassToStandard(t *testing.T) {	
	tests:= []struct{
		name		string
		value		float64
		unit		MassUnit
		prefix		Prefix
		expectedResult float64
		expectedError bool
	}{

		{
			name:           "1 kilogram to grams",
			value:          1,
			unit:           gram,
			prefix:         kilo,
			expectedResult: 1000,
			expectedError:  false,
		},
		{
			name:           "1 hectogram to grams",
			value:          1,
			unit:           gram,
			prefix:         hecto,
			expectedResult: 100,
			expectedError:  false,
		},
		{
			name:           "2 pounds to grams",
			value:          2,
			unit:           pound,
			prefix:         none,
			expectedResult: 907.184, // 2 * 453.592
			expectedError:  false,
		},
		{
			name:           "3 ounces to grams",
			value:          3,
			unit:           ounce,
			prefix:         none,
			expectedResult: 85.047, // 3 * 28.349
			expectedError:  false,
		},
		{
			name:           "100 milligrams to grams",
			value:          100,
			unit:           gram,
			prefix:         milli,
			expectedResult: 0.1,
			expectedError:  false,
		},
		{
			name:           "Invalid unit (unknown)",
			value:          2,
			unit:           unknown,
			prefix:         none,
			expectedResult: 2,
			expectedError:  true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := Mass{value: test.value, unit: test.unit, prefix: test.prefix}
			result := m.convertToStandard()
			if test.expectedError && result == test.expectedResult {
				t.Errorf("Expected error for %s but got %f", test.name, result)
			}
			if !test.expectedError && result != test.expectedResult {
				t.Errorf("Test %s failed: expected %f, got %f", test.name, test.expectedResult, result)
			}
		})
	}
}

func TestConvertVolume(t *testing.T) {
	volumeTests:= []struct{
		name	string
		prefix		Prefix
		value		float64
		expectedResult float64
		expectedError bool
	}{
		{prefix: none, value: 2, expectedResult: 2, expectedError:  false}, 
		{prefix: kilo, value: 8.25, expectedResult:  8250, expectedError:  false},
		{prefix: hecto, value: 9, expectedResult:  900, expectedError:  false},
		{prefix: deca, value: 10, expectedResult:  100, expectedError:  false},
		{prefix: deci, value: 1005, expectedResult:  100.5, expectedError:  false},
		{prefix: centi, value: 888, expectedResult:  8.88, expectedError:  false},
		{prefix: milli, value: 1618, expectedResult:  1.618, expectedError:  false},
		{prefix: micro, value: 1, expectedResult:  0.000001, expectedError:  false},
	}
	for _, test := range volumeTests {
		t.Run(test.name, func(t *testing.T) {
			v := Volume{value: test.value, unit: test.prefix}
			result := v.convertToStandard()
			if test.expectedError && result == test.expectedResult {
				t.Errorf("Expected error for %s but got %f", test.name, result)
			}
			if !test.expectedError && result != test.expectedResult {
				t.Errorf("Test %s failed: expected %f, got %f", test.name, test.expectedResult, result)
			}
		})
	}
}

