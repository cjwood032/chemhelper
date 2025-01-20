package element

import (
	"fmt"
	"testing"
)

// Simplified Periodic Table with only a few elements for testing.
func NewTestPeriodicTable() *PeriodicTable {
	return &PeriodicTable{
		Elements: []Element{
			{AtomicNumber: 1, Symbol: "H", Name: "Hydrogen", AtomicWeight: 1.008},
			{AtomicNumber: 6, Symbol: "C", Name: "Carbon", AtomicWeight: 12.011},
			{AtomicNumber: 8, Symbol: "O", Name: "Oxygen", AtomicWeight: 15.999},
			{AtomicNumber: 11, Symbol: "Na", Name: "Sodium", AtomicWeight: 22.990},
			{AtomicNumber: 17, Symbol: "Cl", Name: "Chlorine", AtomicWeight: 35.45},
		},
	}
}
var testCompounds = []struct { // Mass is the atomic mass
	compound     Compound
	expectedError bool
}{
	{
		compound: Compound{Symbol: "H2O",
		Elements: []ElementMoles{
			{Element: Element{Symbol: "H", Name: "Hydrogen", AtomicNumber: 1, AtomicWeight: 1.008}, Moles: 2},
			{Element: Element{Symbol: "O", Name: "Oxygen", AtomicNumber: 8, AtomicWeight: 15.999}, Moles: 1},
		},
		Mass: Mass{value: 18.015}},
		expectedError: false,
	},
	{
		compound: Compound{Symbol: "NaCl",
		Elements: []ElementMoles{
			{Element: Element{Symbol: "Na", Name: "Sodium", AtomicNumber: 11, AtomicWeight: 22.990}, Moles: 1},
			{Element: Element{Symbol: "Cl", Name: "Chlorine", AtomicNumber: 17, AtomicWeight: 35.45}, Moles: 1},
		},
		Mass: Mass{value: 58.44}},
		expectedError: false,
	},
	{
		compound: Compound{Symbol:  "HOH", // This is an alternative way to symbolize water
		Elements: []ElementMoles{
			{Element: Element{Symbol: "H", Name: "Hydrogen", AtomicNumber: 1, AtomicWeight: 1.008}, Moles: 2},
			{Element: Element{Symbol: "O", Name: "Oxygen", AtomicNumber: 8, AtomicWeight: 15.999}, Moles: 1},
		},
		Mass: Mass{value: 18.015}},
		expectedError: false,
	},
	{
		compound: Compound{Symbol:  "C6H12O6",
		Elements: []ElementMoles{
			{Element: Element{Symbol: "C", Name: "Carbon", AtomicNumber: 6, AtomicWeight: 12.011}, Moles: 6},
			{Element: Element{Symbol: "H", Name: "Hydrogen", AtomicNumber: 1, AtomicWeight: 1.008}, Moles: 12},
			{Element: Element{Symbol: "O", Name: "Oxygen", AtomicNumber: 8, AtomicWeight: 15.999}, Moles: 6},
		},
		Mass: Mass{value: 180.156}},
		expectedError: false,
	},
	{
		compound: Compound{Symbol: "XYZ",
		Elements: nil},
		expectedError: true,
	},
	{
		compound: Compound{Symbol: "H2O1X",
		Elements: nil},
		expectedError: true,
	},
	{
		compound: Compound{Symbol: "",
		Elements: nil},
		expectedError: true,
	},
}


func TestParseCompound(t *testing.T) {
	pt := NewTestPeriodicTable()

	for _, test := range testCompounds {
		t.Run(fmt.Sprintf("Testing Compound:%s", test.compound.Symbol), func(t *testing.T) {
			result, err := ParseCompoundElements(test.compound.Symbol, pt)
			if (err != nil) != test.expectedError {
				t.Errorf("Expected error: %v, but got: %v", test.expectedError, err)
			}

			// If no error is expected, compare the result to the expected value
			if !test.expectedError {
				if len(result) != len(test.compound.Elements) {
					t.Errorf("Expected %d elements, but got %d", len(test.compound.Elements), len(result))
				}

				// Sort both the result and the expected elements to ensure order-independence
				sortElementMoles(result)
				sortElementMoles(test.compound.Elements)

				for i, elem := range result {
					if elem.Element.Symbol != test.compound.Elements[i].Element.Symbol || elem.Moles != test.compound.Elements[i].Moles {
						t.Errorf("Expected ElementMoles {%s, %.2f}, but got {%s, %.2f}",
							test.compound.Elements[i].Element.Symbol, test.compound.Elements[i].Moles,
							elem.Element.Symbol, elem.Moles)
					}
				}
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
