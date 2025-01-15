package element

import (
	"fmt"
	"regexp"
	"strconv"
)

type ElementMoles struct { // when creating compounds 
	Element Element
	Moles float64
}

type Compound struct {
	Symbol string
	Elements []ElementMoles
	Mass Mass
	Volume Volume
}

func ParseCompound(compound string, pt *PeriodicTable) ([]ElementMoles, error) {
	if compound =="" {
		return nil, fmt.Errorf("No compound symbols passed")
	}
	var elements []ElementMoles
	re := regexp.MustCompile(`([A-Z][a-z]?)(\d*)`)
	elementMolesMap := make(map[string]float64)
	matches := re.FindAllStringSubmatch(compound, -1)

	for _, match := range matches {
		elementSymbol := match[1]
		count := match[2]

		if count == "" {
			count = "1"
		}
		moles, err := strconv.ParseFloat(count, 64)
		if err != nil {
			return nil, err
		}
		_, found := pt.FindElementBySymbol(elementSymbol)
		if !found {
			return nil, fmt.Errorf("element %s not found in the periodic table", elementSymbol)
		}

		elementMolesMap[elementSymbol] += moles
	}
	for symbol, moles := range elementMolesMap {
		element, _ := pt.FindElementBySymbol(symbol) // We know the element exists, so this is safe
		elements = append(elements, ElementMoles{Element: *element, Moles: moles})
	}
	return elements, nil
}
func (c Compound) MolarMass() float64 {
	totalMass := 0.0
	for _, em := range c.Elements {
		totalMass += em.Element.AtomicWeight * em.Moles
	}
	return totalMass
}
