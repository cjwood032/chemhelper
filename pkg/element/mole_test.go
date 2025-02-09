package element

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
)

// need volume moles
// need to test get moles with combination of masses/volumes

func TestMolesByMass(t *testing.T) {
	for _, test := range TestCompounds {
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

func TestMolesByVolume(t *testing.T) {
	for _, test := range TestCompounds {
		t.Run(fmt.Sprintf("Testing Compound:%s", test.compound.Symbol), func(t *testing.T) {
			expected := test.expectedMoles
			actualMoles, err := test.compound.Volume.getMoles(test.molarity)
			if !test.expectedError && err!=nil {
				t.Errorf("unexpected error")
			}
			if (!actualMoles.Equal(expected)){
				t.Errorf("Expected %v moles, but got %v", expected, actualMoles)
			}
			if (test.expectedError && err == nil) {
				t.Errorf("Expected error but got none")
			}
		})
	}
}

func TestMoles(t *testing.T) {
	for _, test := range TestCompounds {
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
