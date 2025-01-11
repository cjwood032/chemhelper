package test

import (
	"testing"

	"github.com/cjwood032/chemhelper/pkg/models"
)
func TestNewPeriodicTableSize(t *testing.T){
	pd := *models.NewPeriodicTable()
	expected := 118
	if len(pd.Elements) != expected {
		t.Errorf("Expected table length of %v, but got %v", expected, len(pd.Elements))
	}
}

//We test the main organic elements since they are the most used, and most important
func TestNewPeriodicTableElements(t *testing.T){
	pd := *models.NewPeriodicTable()
	expectedHydrogenValues :=  newTestElement(1, "H", "Hydrogen", 1.008, 2.20, 120.0, 1, 1)
	actualHydrogenValues := pd.Elements[0]
	expectedCarbonValues := newTestElement(6, "C", "Carbon", 12.011, 2.55, 170.0, 14, 2)
	actualCarbonValues := pd.Elements[5]
	expectedNitrogenValues := newTestElement(7, "N", "Nitrogen", 14.007, 3.04, 155.0, 15, 2)
	actualNitrogenValues := pd.Elements[6]
	expectedOxygenValues := newTestElement(8, "O", "Oxygen", 15.999, 3.44, 152.0, 16, 2)
	actualOxygenValues := pd.Elements[7]
	if  actualHydrogenValues != expectedHydrogenValues {
		ThrowElementError(t, expectedHydrogenValues, actualHydrogenValues)
	}
	if  expectedCarbonValues != actualCarbonValues {
		ThrowElementError(t, expectedCarbonValues, actualCarbonValues)
	}
	if  expectedNitrogenValues != actualNitrogenValues {
		ThrowElementError(t, expectedNitrogenValues, actualNitrogenValues)
	}
	if  expectedOxygenValues != actualOxygenValues {
		ThrowElementError(t, expectedOxygenValues, actualOxygenValues)
	}
}
func newTestElement(number int, symbol string, name string, weight float64, en float64, radius float64, group int, period int, ) models.Element{
	return models.Element{AtomicNumber: number, Symbol: symbol, Name: name, AtomicWeight: weight, Electronegativity: en, VanDerWaalsRadius: radius, Group: group, Period: period}
}
func ThrowElementError(t *testing.T, expected models.Element, actual models.Element,) {
	t.Errorf("Expected %s to equal %v, but got %v", expected.Name, expected, actual)
}