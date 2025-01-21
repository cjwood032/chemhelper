package element

import (
	"fmt"
	"log"
	"math"
)


type Prefix float64

const (
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
	getMoles(float64) (float64, error)
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

func NewMass(value float64, options ...interface{}) (Mass, error) {
	if value == 0 {
		return Mass{}, fmt.Errorf("no mass value passed")
	}
	mass := Mass{
		value:  value,
		unit:   gram, 
		prefix: none, 
	}

	for _, opt := range options {
		switch v := opt.(type) {
		case MassUnit:
			mass.unit = v
		case Prefix:
			mass.prefix = v
		default:
			log.Printf("%v is unexpected", v)
		}
	}

	return mass, nil
}


func (v Volume) convertToStandard() (float64, error) {
	return v.value * float64(v.unit), nil
}

func getMoles(p Property, value float64) (float64, error){
	return p.getMoles(value)
}
func (m Mass) getMoles(molarMass float64) (float64,error) {
	standardMass, err := m.convertToStandard()
	if err != nil {
		return 0, err
	}
	return standardMass / molarMass, nil
}
func (c *Compound) getMoles(mass float64) ( error) {
	if mass ==0 {
		return fmt.Errorf("no mass passed")
	}
	if c.MolarMass == 0 {
		err := c.getMolarMass()
		if (err != nil) {
			return err
		}	
	}
	c.Moles = mass/c.MolarMass
	return nil
}
func (v Volume) getMoles(molarity float64) (float64, error) {
	standardVol, err := v.convertToStandard()
	handleError(err)
	return standardVol * molarity, err
}

func (element *ElementMoles) getMoles(mass Mass) error {
	
	moles, err := mass.getMoles(element.Element.AtomicWeight)
	if err != nil {
		return err
	} 
	element.Moles = moles
	return nil
}

func (compound *Compound) getMolesFromMass(mass Mass) error {
	if compound.MolarMass == 0 {
		err := compound.getMolarMass()
		if (err != nil) {
			return err
		}	
	}
	moles, err := mass.getMoles(compound.MolarMass)
	if (err != nil) {
		return err
	}	
	compound.Moles = moles
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
	//I don't really do anything with this yet.
	if err != nil {
		log.Fatal(err)
	}
}

func SetToSigFigs(num float64) float64 {
	multiplier := math.Pow(10, 4) // currently just rounds to 4 decimals for now, sig figs are complex
	return math.Round(num*multiplier) / multiplier
}