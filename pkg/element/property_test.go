package element

import (
	"fmt"
	"testing"
)
var testCompounds = []struct {
	compound     Compound  // this mass is the atomic mass
	expectedError bool
	massForMoles Mass // this mass is for molar calculations
	expectedMoles float64
}{
	{
		compound: Compound{
			Symbol: "H2O",
			Elements: []ElementMoles{
				{Element: Element{Symbol: "H", Name: "Hydrogen", AtomicNumber: 1, AtomicWeight: 1.008}, Moles: 2},
				{Element: Element{Symbol: "O", Name: "Oxygen", AtomicNumber: 8, AtomicWeight: 15.999}, Moles: 1},
			},
			Mass: Mass{value: 18.015}},
			massForMoles: Mass{value: 18.015, unit: gram, prefix: none},
			expectedMoles: 1,
			expectedError: false,

	},
	{
		compound: Compound{
			Symbol: "NaCl",
			Elements: []ElementMoles{
				{Element: Element{Symbol: "Na", Name: "Sodium", AtomicNumber: 11, AtomicWeight: 22.990}, Moles: 1},
				{Element: Element{Symbol: "Cl", Name: "Chlorine", AtomicNumber: 17, AtomicWeight: 35.45}, Moles: 1},
			},
			Mass: Mass{value: 58.44}},
			massForMoles: Mass{value: 58.44,prefix: kilo, unit: gram},
			expectedMoles: 1000,
			expectedError: false,
	},
	{
		compound: Compound{Symbol:  "HOH", // This is an alternative way to symbolize water
			Elements: []ElementMoles{
			{Element: Element{Symbol: "H", Name: "Hydrogen", AtomicNumber: 1, AtomicWeight: 1.008}, Moles: 2},
			{Element: Element{Symbol: "O", Name: "Oxygen", AtomicNumber: 8, AtomicWeight: 15.999}, Moles: 1},
		},
		Mass: Mass{value: 18.015}},
		massForMoles: Mass{value: 50, prefix: none, unit: pound},
		expectedMoles: 1258.9287,
		expectedError: false,
	},
	{
		compound: Compound{
			Symbol:  "C6H12O6",
			Elements: []ElementMoles{
				{Element: Element{Symbol: "C", Name: "Carbon", AtomicNumber: 6, AtomicWeight: 12.011}, Moles: 6},
				{Element: Element{Symbol: "H", Name: "Hydrogen", AtomicNumber: 1, AtomicWeight: 1.008}, Moles: 12},
				{Element: Element{Symbol: "O", Name: "Oxygen", AtomicNumber: 8, AtomicWeight: 15.999}, Moles: 6},
			},
		Mass: Mass{value: 180.156}},
		massForMoles: Mass{value: 18.0156,prefix: deca, unit: gram},
		expectedMoles: 1,
		expectedError: false,
	},
	{
		compound: Compound{
			Symbol: "XYZ",
			Elements: nil},
		massForMoles: Mass{value: 5,prefix: kilo, unit: gram},
		expectedMoles: 0,
		expectedError: true,
	},
	{
		compound: Compound{Symbol: "H2O1X",
		Elements: nil},
		massForMoles: Mass{value: 15,prefix: kilo, unit: gram},
		expectedMoles: 0,
		expectedError: true,
	},
	{
		compound: Compound{Symbol: "",
		Elements: nil},
		massForMoles: Mass{value: 1,prefix: kilo, unit: gram},
		expectedMoles: 0,
		expectedError: true,
	},
}
var testElementMoles = []struct {
	element     ElementMoles  
	expectedError bool
	massForMoles Mass // this mass is for molar calculations
	expectedMoles float64
}{
	{element: ElementMoles{
		Element: Element{AtomicNumber: 1, Symbol: "H", Name: "Hydrogen", AtomicWeight: 1.008}},
		massForMoles: Mass{value: 10.08, unit: gram, prefix: none},
		expectedMoles: 10,
		expectedError: false,
	},
	{element: ElementMoles{
		Element: Element{AtomicNumber: 6, Symbol: "C", Name: "Carbon", AtomicWeight: 12.011}},
		massForMoles: Mass{value: 7.2066, unit: gram, prefix: kilo},
		expectedMoles: 600,
		expectedError: false,
	},
	{element: ElementMoles{
		Element: Element{AtomicNumber: 8, Symbol: "O", Name: "Oxygen", AtomicWeight: 15.999}},
		massForMoles: Mass{value: 0.006354604, unit: ounce, prefix: none},
		expectedMoles: 0.0113,
		expectedError: false,
	},
	{element: ElementMoles{
		Element: Element{AtomicNumber: 211, Symbol: "XX", Name: "Baddium", AtomicWeight: 85501}},
		massForMoles: Mass{value: 1, unit: pound, prefix: none},
		expectedMoles: .0053,
		expectedError: false,
	},
	{element: ElementMoles{
		Element: Element{AtomicNumber: 17, Symbol: "Cl", Name: "Chlorine", AtomicWeight: 35.45}},
		massForMoles: Mass{value: 3.545, unit: gram, prefix: milli},
		expectedMoles: 0.0001,
		expectedError: false,
	},
}
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
			unit:           unknownMass,
			prefix:         none,
			expectedResult: 2,
			expectedError:  true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := Mass{value: test.value, unit: test.unit, prefix: test.prefix}
			result, err := convertToStandardValue(m)
			if !test.expectedError && err !=nil  {
				t.Errorf("Unexpected error for %s: %s", test.name, err)
			}
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
			result, err := convertToStandardValue(v)
			if !test.expectedError && err !=nil  {
				t.Errorf("Unexpected error for %s: %s", test.name, err)
			}
			if test.expectedError && result == test.expectedResult {
				t.Errorf("Expected error for %s but got %f", test.name, result)
			}
			if !test.expectedError && result != test.expectedResult {
				t.Errorf("Test %s failed: expected %f, got %f", test.name, test.expectedResult, result)
			}
		})
	}
}

func TestMolarMass(t *testing.T) {
	for _, test := range testCompounds {
		t.Run(fmt.Sprintf("Testing Compound:%s", test.compound.Symbol), func(t *testing.T) {
			expected := test.compound.Mass.value
			err := test.compound.getMolarMass()
			if !test.expectedError && err!=nil {
				t.Errorf("unexpected error")
			}
			actual := test.compound.MolarMass
			if (actual != expected){
				t.Errorf("Expected %v mass, but got %v", expected, actual)
			}
			if (test.expectedError && err == nil) {
				t.Errorf("Expected error but got none")
			}
		})
	}
}
func TestMolesByMass(t *testing.T) {
	for _, test := range testCompounds {
		t.Run(fmt.Sprintf("Testing Compound:%s", test.compound.Symbol), func(t *testing.T) {
			expected := test.expectedMoles
			err := test.compound.getMolesFromMass(test.massForMoles)
			if !test.expectedError && err!=nil {
				t.Errorf("unexpected error")
			}
			actual := SetToSigFigs(test.compound.Moles)
			if (actual != expected){
				t.Errorf("Expected %v moles, but got %v", expected, actual)
			}
			if (test.expectedError && err == nil) {
				t.Errorf("Expected error but got none")
			}
		})
	}
}
func TestMoles(t *testing.T) {
	for _, test := range testCompounds {
		t.Run(fmt.Sprintf("Testing Compound:%s", test.compound.Symbol), func(t *testing.T) {
			expected := test.expectedMoles
			test.compound.MolarMass = test.compound.Mass.value
			standardMass,_ := test.massForMoles.convertToStandard()
			err := test.compound.getMoles(standardMass)
			if !test.expectedError && err!=nil {
				t.Errorf("unexpected error")
			}
			actual := SetToSigFigs(test.compound.Moles)
			if (actual != expected){
				t.Errorf("Expected %v moles, but got %v", expected, actual)
			}
			if (test.expectedError && err == nil) {
				t.Errorf("Expected error but got none")
			}
		})
	}
}
func TestMolesOfElements(t *testing.T) {
	for _, test := range testElementMoles {
		t.Run(fmt.Sprintf("Testing Element:%s", test.element.Element.Symbol), func(t *testing.T) {
			expected := test.expectedMoles
			
			
			err := test.element.getMoles(test.massForMoles)
			if !test.expectedError && err!=nil {
				t.Errorf("unexpected error")
			}
			actual := SetToSigFigs(test.element.Moles)
			if (actual != expected){
				t.Errorf("Expected %v moles, but got %v", expected, actual)
			}
			if (test.expectedError && err == nil) {
				t.Errorf("Expected error but got none")
			}
		})
	}
}

func TestNewMass(t *testing.T){
	var gram MassUnit = gram
	var pound MassUnit = pound
	var kilo Prefix = kilo
	
	tests := []struct {
		name		string
		value float64
		unit *MassUnit
		prefix *Prefix
		expectedResult float64
		expectedError bool
	}{
		{
			name:           "1 kilogram to grams",
			value:          1,
			unit:           &gram,
			prefix:         &kilo,
			expectedResult: 1000,
			expectedError:  false,
		},
		{
			name:           ".5 kilogram with no unit to grams",
			value:          .5,
			unit:           nil,
			prefix:         &kilo,
			expectedResult: 500,
			expectedError:  false,
		},
		{
			name:           "1 pounds to grams no prefix",
			value:          1,
			unit:           &pound,
			prefix:         nil,
			expectedResult: 453.592,
			expectedError:  false,
		},
		{
			name:           "10 grams only passing value",
			value:          10,
			unit:           nil,
			prefix:         nil,
			expectedResult: 10,
			expectedError:  false,
		},
		{
			name:           "Invalid value",
			value:          0,
			unit:           &gram,
			prefix:         &kilo,
			expectedResult: 1000,
			expectedError:  true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var m Mass
			var e error
			//if neither are null
			// 2 cases if one is null
			//if both are null
			if test.unit != nil && test.prefix !=nil {
				m,e =  NewMass(test.value, *test.unit ,*test.prefix)
			} else if test.unit == nil && test.prefix !=nil {
				m,e =  NewMass(test.value, *test.prefix)
			} else if test.unit != nil && test.prefix ==nil {
				m,e =  NewMass(test.value, *test.unit)
			} else if test.unit == nil && test.prefix ==nil {
				m,e =  NewMass(test.value)
			} else {
				t.Error("Unexpected testing condition")
			}
			

			result, err := convertToStandardValue(m)
			if !test.expectedError && (err !=nil || e !=nil)  {
				t.Errorf("Unexpected error for %s: %s", test.name, err)
			}
			if test.expectedError && result == test.expectedResult {
				t.Errorf("Expected error for %s but got %f", test.name, result)
			}
			if !test.expectedError && result != test.expectedResult {
				t.Errorf("Test %s failed: expected %f, got %f", test.name, test.expectedResult, result)
			}
		})
	}
}