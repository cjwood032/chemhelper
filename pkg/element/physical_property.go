package element

import (
	"fmt"
	"log"
)


type Prefix float64

const (
	unknownPrefix Prefix = 0
	none Prefix = 1
	kilo Prefix = 1000
	hecto Prefix = 100
	deca Prefix = 10
	deci Prefix = 0.1
	centi Prefix = 0.01
	milli Prefix = 0.001
	micro Prefix = 0.000001
)
type MassUnit string

const (
	// Catch for things like Ton, and Slug
	unknownMass MassUnit = "unknown"
	//Metric
	gram    MassUnit = "gram"
	// Imperial
	ounce   MassUnit = "ounce"
	pound   MassUnit = "pound"
)

type Mass struct {
	value float64
	unit  MassUnit
	prefix Prefix
}


type Volume struct {
	value float64
	unit  Prefix 
}

type Property interface {
	convertToStandard() (float64, error)
	getMoles(float64) float64
}

func convertToStandardValue(p Property) (float64, error) {
	
    return p.convertToStandard()
}

func (m Mass) convertToStandard() (float64, error) {	
	
	switch m.unit {
	case pound:
		return m.value * 453.592 * float64(m.prefix), nil
	case ounce:
		return m.value * 28.349 * float64(m.prefix), nil
	case gram:
		return m.value * float64(m.prefix), nil
	default:
		return 0, fmt.Errorf("unknown unit passed")
	}
}

func (v Volume) convertToStandard() (float64, error) {
	
	return v.value * float64(v.unit), nil
}

func getMoles(p Property, value float64) {
	p.getMoles(value)
}
func (m Mass) getMoles(molarMass float64) float64 {
	standardMass, err := m.convertToStandard()
	handleError(err)
	return standardMass / molarMass
}
func (v Volume) getMoles(molarity float64) float64 {
	standardVol, err := v.convertToStandard()
	handleError(err)
	return standardVol * molarity
}

func (element *ElementMoles) getMoles(mass Mass) {
	element.Moles = mass.getMoles(element.Element.AtomicWeight)
}

func (compound *Compound) getMolesFromMass(mass Mass) error {
	if compound.MolarMass == 0 {
		err := compound.getMolarMass()
		if (err != nil) {
			return err
		}	
	}
	compound.Moles = mass.getMoles(compound.MolarMass)
	return nil
}

func (compound *Compound) getMolarMass() error {
	if len(compound.Elements) == 0 {
		return fmt.Errorf("no elements passed")
	}
	totalMass := 0.0
	for _, em := range compound.Elements {
		totalMass += em.Element.AtomicWeight * em.Moles
	}
	compound.MolarMass = totalMass
	return nil
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}