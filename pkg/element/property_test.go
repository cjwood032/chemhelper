package element

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
)

// need volume moles
// need to test get moles with combination of masses/volumes
var preciseHOHMoles, _ =decimal.NewFromString("1258.9286705523175132") // floats are not precise enough
var testCompounds = []struct {
	compound     Compound  // this mass is the atomic mass
	expectedError bool
	massForMoles Mass // this mass is for molar calculations
	expectedMoles decimal.Decimal
}{
	{
		compound: Compound{
			Symbol: "H2O",
			Elements: []ElementMoles{
				{Element: Element{Symbol: "H", Name: "Hydrogen", AtomicNumber: 1, AtomicWeight: decimal.NewFromFloat(1.008)}, Moles: decimal.NewFromFloat(2)},
				{Element: Element{Symbol: "O", Name: "Oxygen", AtomicNumber: 8, AtomicWeight: decimal.NewFromFloat(15.999)}, Moles: decimal.NewFromFloat(1)},
			},
			Mass: Mass{value: decimal.NewFromFloat(18.015)}},
			massForMoles: Mass{value: decimal.NewFromFloat(18.015), unit: gram, prefix: none},
			expectedMoles: decimal.NewFromFloat(1),
			expectedError: false,

	},
	{
		compound: Compound{
			Symbol: "NaCl",
			Elements: []ElementMoles{
				{Element: Element{Symbol: "Na", Name: "Sodium", AtomicNumber: 11, AtomicWeight: decimal.NewFromFloat(22.990)}, Moles: decimal.NewFromFloat(1)},
				{Element: Element{Symbol: "Cl", Name: "Chlorine", AtomicNumber: 17, AtomicWeight: decimal.NewFromFloat(35.45)}, Moles: decimal.NewFromFloat(1)},
			},
			Mass: Mass{value: decimal.NewFromFloat(58.44)}},
			massForMoles: Mass{value: decimal.NewFromFloat(58.44), prefix: kilo, unit: gram},
			expectedMoles: decimal.NewFromFloat(1000),
			expectedError: false,
	},
	{
		compound: Compound{Symbol:  "HOH", // This is an alternative way to symbolize water
			Elements: []ElementMoles{
			{Element: Element{Symbol: "H", Name: "Hydrogen", AtomicNumber: 1, AtomicWeight: decimal.NewFromFloat(1.008)}, Moles: decimal.NewFromFloat(2)},
			{Element: Element{Symbol: "O", Name: "Oxygen", AtomicNumber: 8, AtomicWeight: decimal.NewFromFloat(15.999)}, Moles: decimal.NewFromFloat(1)},
		},
		Mass: Mass{value: decimal.NewFromFloat(18.015)}},
		massForMoles: Mass{value: decimal.NewFromFloat(50), prefix: none, unit: pound},
		expectedMoles: preciseHOHMoles,
		expectedError: false,
	},
	{
		compound: Compound{
			Symbol:  "C6H12O6",
			Elements: []ElementMoles{
				{Element: Element{Symbol: "C", Name: "Carbon", AtomicNumber: 6, AtomicWeight: decimal.NewFromFloat(12.011)}, Moles: decimal.NewFromFloat(6)},
				{Element: Element{Symbol: "H", Name: "Hydrogen", AtomicNumber: 1, AtomicWeight: decimal.NewFromFloat(1.008)}, Moles: decimal.NewFromFloat(12)},
				{Element: Element{Symbol: "O", Name: "Oxygen", AtomicNumber: 8, AtomicWeight: decimal.NewFromFloat(15.999)}, Moles: decimal.NewFromFloat(6)},
			},
		Mass: Mass{value: decimal.NewFromFloat(180.156)}},
		massForMoles: Mass{value: decimal.NewFromFloat(18.0156),prefix: deca, unit: gram},
		expectedMoles: decimal.NewFromFloat(1),
		expectedError: false,
	},
	{
		compound: Compound{
			Symbol: "XYZ",
			Elements: nil},
		massForMoles: Mass{value: decimal.NewFromFloat(5),prefix: kilo, unit: gram},
		expectedMoles: decimal.Zero,
		expectedError: true,
	},
	{
		compound: Compound{Symbol: "H2O1X",
		Elements: nil},
		massForMoles: Mass{value: decimal.NewFromFloat(15), prefix: kilo, unit: gram},
		expectedMoles: decimal.Zero,
		expectedError: true,
	},
	{
		compound: Compound{Symbol: "",
		Elements: nil},
		massForMoles: Mass{value: decimal.NewFromFloat(1), prefix: kilo, unit: gram},
		expectedMoles: decimal.Zero,
		expectedError: true,
	},
	{
		compound: Compound{
			Symbol: "H2O",
			Elements: []ElementMoles{
				{Element: Element{Symbol: "H", Name: "Hydrogen", AtomicNumber: 1, AtomicWeight: decimal.NewFromFloat(1.008)}, Moles: decimal.NewFromFloat(2)},
				{Element: Element{Symbol: "O", Name: "Oxygen", AtomicNumber: 8, AtomicWeight: decimal.NewFromFloat(15.999)}, Moles: decimal.NewFromFloat(1)},
			},
			Mass: Mass{value: decimal.NewFromFloat(18.015)}},
			massForMoles: Mass{value: decimal.NewFromFloat(18.015), unit: gram, prefix: micro},
			expectedMoles: decimal.NewFromFloat(.000001),
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
			name:           "1 micrograms to grams",
			value:          1,
			unit:           gram,
			prefix:         micro,
			expectedResult: .000001,
			expectedError:  false,
		},
		{
			name:           "1000 pounds to grams",
			value:          1000,
			unit:           pound,
			prefix:         none,
			expectedResult: 453592,
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
			m := Mass{value: decimal.NewFromFloat(test.value), unit: test.unit, prefix: test.prefix}
			result, err := convertToStandardValue(m)
			if !test.expectedError && err !=nil  {
				t.Errorf("Unexpected error for %s: %s", test.name, err)
			}
			if test.expectedError && result.Equal(decimal.NewFromFloat(test.expectedResult)) {
				t.Errorf("Expected error for %s but got %v", test.name, result)
			}
			if !test.expectedError && !result.Equal(decimal.NewFromFloat(test.expectedResult)) {
				t.Errorf("Test %s failed: expected %v, got %v", test.name, test.expectedResult, result)
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
			v := Volume{value: decimal.NewFromFloat(test.value), unit: test.prefix}
			result, err := convertToStandardValue(v)
			if !test.expectedError && err !=nil  {
				t.Errorf("Unexpected error for %s: %s", test.name, err)
			}
			if test.expectedError && result.Equal(decimal.NewFromFloat(test.expectedResult)) {
				t.Errorf("Expected error for %s but got %v", test.name, result)
			}
			if !test.expectedError && !result.Equal(decimal.NewFromFloat(test.expectedResult)) {
				t.Errorf("Test %s failed: expected %v, got %v", test.name, test.expectedResult, result)
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
			if !actual.Equal(expected){
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
			actual := test.compound.Moles
			if (!actual.Equal(expected)){
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
			actual := test.compound.Moles
			if (!actual.Equal(expected)){
				t.Errorf("Expected %v moles, but got %v", expected, actual)
			}
			if (test.expectedError && err == nil) {
				t.Errorf("Expected error but got none")
			}
		})
	}
}

func TestMolesOfElements(t *testing.T) {
	var preciseOMoles, _ = decimal.NewFromString("17.7192324520282518") //too precise for floats
	var testElementMoles = []struct {
		element     ElementMoles  
		expectedError bool
		massForMoles Mass // this mass is for molar calculations
		expectedMoles decimal.Decimal
	}{
		{element: ElementMoles{
			Element: Element{AtomicNumber: 1, Symbol: "H", Name: "Hydrogen", AtomicWeight: decimal.NewFromFloat(1.008)}},
			massForMoles: Mass{value: decimal.NewFromFloat(10.08), unit: gram, prefix: none},
			expectedMoles: decimal.NewFromFloat(10),
			expectedError: false,
		},
		{element: ElementMoles{
			Element: Element{AtomicNumber: 6, Symbol: "C", Name: "Carbon", AtomicWeight: decimal.NewFromFloat(12.011)}},
			massForMoles: Mass{value: decimal.NewFromFloat(7.2066), unit: gram, prefix: kilo},
			expectedMoles: decimal.NewFromFloat(600),
			expectedError: false,
		},
		{element: ElementMoles{
			Element: Element{AtomicNumber: 8, Symbol: "O", Name: "Oxygen", AtomicWeight: decimal.NewFromFloat(15.999)}},
			massForMoles: Mass{value: decimal.NewFromFloat(10), unit: ounce, prefix: none},
			expectedMoles: preciseOMoles,
			expectedError: false,
		},
		{element: ElementMoles{
			Element: Element{AtomicNumber: 211, Symbol: "XX", Name: "Baddium", AtomicWeight: decimal.NewFromFloat(453.592)}},
			massForMoles: Mass{value: decimal.NewFromFloat(100), unit: pound, prefix: none},
			expectedMoles: decimal.NewFromFloat(100),
			expectedError: false,
		},
		{element: ElementMoles{
			Element: Element{AtomicNumber: 17, Symbol: "Cl", Name: "Chlorine", AtomicWeight: decimal.NewFromFloat(35.45)}},
			massForMoles: Mass{value: decimal.NewFromFloat(3.545), unit: gram, prefix: milli},
			expectedMoles: decimal.NewFromFloat(0.0001),
			expectedError: false,
		},
	}
	for _, test := range testElementMoles {
		t.Run(fmt.Sprintf("Testing Element:%s", test.element.Element.Symbol), func(t *testing.T) {
			expected := test.expectedMoles
			err := test.element.getMoles(test.massForMoles)
			if !test.expectedError && err!=nil {
				t.Errorf("unexpected error")
			}
			actual := test.element.Moles//SetToSigFigs(test.element.Moles, 4)
			if (!actual.Equal(expected)){
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
			value := decimal.NewFromFloat(test.value)
			if test.unit != nil && test.prefix !=nil {
				m,e =  NewMass(value, *test.unit ,*test.prefix)
			} else if test.unit == nil && test.prefix !=nil {
				m,e =  NewMass(value, *test.prefix)
			} else if test.unit != nil && test.prefix ==nil {
				m,e =  NewMass(value, *test.unit)
			} else if test.unit == nil && test.prefix ==nil {
				m,e =  NewMass(value)
			} else {
				t.Error("Unexpected testing condition")
			}

			result, err := convertToStandardValue(m)
			if !test.expectedError && (err !=nil || e !=nil)  {
				t.Errorf("Unexpected error for %s: %s", test.name, err)
			}
			if test.expectedError && (err == nil && e == nil){
				t.Errorf("Expected error for %s but got %v", test.name, result)
			}
			if !test.expectedError && !result.Equal(decimal.NewFromFloat(test.expectedResult)) {
				t.Errorf("Test %s failed: expected %v, got %v", test.name, test.expectedResult, result)
			}
		})
	}
}
