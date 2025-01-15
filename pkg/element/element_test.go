package element

import (
	"testing"
)
 var elementSymbols = []string{
	"H", "He", "Li", "Be", "B", "C", "N", "O", "F", "Ne", 
	"Na", "Mg", "Al", "Si", "P", "S", "Cl", "Ar", "K", "Ca", 
	"Sc", "Ti", "V", "Cr", "Mn", "Fe", "Co", "Ni", "Cu", "Zn", 
	"Ga", "Ge", "As", "Se", "Br", "Kr", "Rb", "Sr", "Y", "Zr", 
	"Nb", "Mo", "Tc", "Ru", "Rh", "Pd", "Ag", "Cd", "In", "Sn", 
	"Sb", "I", "Te", "Xe", "Cs", "Ba", "La", "Ce", "Pr", "Nd", 
	"Pm", "Sm", "Eu", "Gd", "Tb", "Dy", "Ho", "Er", "Tm", "Yb", 
	"Lu", "Hf", "Ta", "W", "Re", "Os", "Ir", "Pt", "Au", "Hg", 
	"Tl", "Pb", "Bi", "Po", "At", "Rn", "Fr", "Ra", "Ac", "Th", 
	"Pa", "U", "Np", "Pu", "Am", "Cm", "Bk", "Cf", "Es", "Fm", 
	"Md", "No", "Lr", "Rf", "Db", "Sg", "Bh", "Hs", "Mt", "Ds", 
	"Rg", "Cn", "Fl", "Mc", "Lv", "Ts", "Og",
}

func TestNewPeriodicTableSize(t *testing.T){
	pd := NewPeriodicTable()
	expected := 118
	if len(pd.Elements) != expected {
		t.Errorf("Expected table length of %v, but got %v", expected, len(pd.Elements))
	}
}

//We test the main organic elements since they are the most used, and most important
func TestNewPeriodicTableElements(t *testing.T){
	pd := NewPeriodicTable()
	expectedElements := []struct {
		index   int
		element Element
	}{
		{0, newTestElement(1, "H", "Hydrogen", 1.008, 2.20, 120.0, 1, 1)},
		{5, newTestElement(6, "C", "Carbon", 12.011, 2.55, 170.0, 14, 2)},
		{6, newTestElement(7, "N", "Nitrogen", 14.007, 3.04, 155.0, 15, 2)},
		{7, newTestElement(8, "O", "Oxygen", 15.999, 3.44, 152.0, 16, 2)},
	}

	// Iterate over expected elements and compare
	for _, e := range expectedElements {
		actual := pd.Elements[e.index]
		checkElement(t, e.element, actual)
	}

}
func TestFindElementBySymbol(t *testing.T){
	pd := NewPeriodicTable()
	expected := newTestElement(46, "Pd", "Palladium", 106.42, 2.20, 163.0, 10, 5)
	actual, found := pd.FindElementBySymbol(expected.Symbol)
	if actual == nil || !found {
		t.Errorf("Expected to find %s in the periodic table ", expected.Name)
	}
	checkElement(t, expected, *actual)
	expected = newTestElement(27, "Co", "Cobalt", 58.933194, 1.88, 152.0, 9, 4)
	actual, found = pd.FindElementBySymbol(expected.Symbol)
	if actual == nil || !found {
		t.Errorf("Expected to find %s in the periodic table ", expected.Name)
	}
	checkElement(t, expected, *actual)
	for _, expectedSymbol := range elementSymbols {
		_, found = pd.FindElementBySymbol(expectedSymbol)
		if !found {
			t.Errorf("Expected to find %s in the periodic table ", expectedSymbol)
		}
	}
	unexpectedSymbol := "Xx"
	_, found = pd.FindElementBySymbol(unexpectedSymbol)
		if found {
			t.Errorf("Expected to find %s in the periodic table ", unexpectedSymbol)
		}
}

func newTestElement(number int, symbol string, name string, weight float64, en float64, radius float64, group int, period int, ) Element{
	return Element{AtomicNumber: number, Symbol: symbol, Name: name, AtomicWeight: weight, Electronegativity: en, VanDerWaalsRadius: radius, Group: group, Period: period}
}
func checkElement(t *testing.T, expected Element, actual Element,) {
	if actual != expected {
		t.Errorf("Expected %s to equal %v, but got %v", expected.Name, expected, actual)
	}
}
