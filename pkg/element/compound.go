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
	Mass MassUnit
	Volume VolumeUnit
}

func parseCompound(compound string, pt *PeriodicTable) ([]ElementMoles, error) {
	var elements []ElementMoles
	re := regexp.MustCompile(`([A-Z][a-z]?)(\d*)`)
	
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
		element, found := pt.FindElementBySymbol(elementSymbol)
		if !found {
			return nil, fmt.Errorf("element %s not found in the periodic table", elementSymbol)
		}

		elements = append(elements, ElementMoles{Element: *element, Moles: moles})
	}

	return elements, nil
}