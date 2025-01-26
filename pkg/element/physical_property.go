package element

import (
	"fmt"
	"log"

	"github.com/shopspring/decimal"
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

type MassUnit float64

const (
	// Catch for things like Ton, and Slug
	unknownMass MassUnit = -1
	//Metric
	gram    MassUnit = 1
	// Imperial
	ounce   MassUnit = 28.349
	pound   MassUnit = 453.592
)

type Mass struct {
	value decimal.Decimal
	unit  MassUnit
	prefix Prefix
}


type Volume struct {
	value decimal.Decimal
	unit  Prefix 
}

type Property interface {
	convertToStandard() (decimal.Decimal, error)
	getMoles(decimal.Decimal) (decimal.Decimal, error)
}

func convertToStandardValue(p Property) (decimal.Decimal, error) {
    return p.convertToStandard()
}

func (m Mass) convertToStandard() (decimal.Decimal, error) {	
		return m.value.Mul( decimal.NewFromFloat(float64(m.unit))).Mul(decimal.NewFromFloat(float64(m.prefix))), nil
	}

func NewMass(value decimal.Decimal, options ...interface{}) (Mass, error) {
	if value.Equal(decimal.Zero) {
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

func (v Volume) convertToStandard() (decimal.Decimal, error) {
	return v.value.Mul(decimal.NewFromFloat(float64(v.unit))), nil
}

func getMoles(p Property, value decimal.Decimal) (decimal.Decimal, error){
	return p.getMoles(value)
}

func (m Mass) getMoles(molarMass decimal.Decimal) (decimal.Decimal, error) {
	standardMass, err := m.convertToStandard()
	if err != nil {
		return decimal.Zero, err
	}
	return standardMass.Div(molarMass), nil
}

func (c *Compound) getMoles(mass decimal.Decimal) ( error) {
	if mass .Equal(decimal.Zero) {
		return fmt.Errorf("no mass passed")
	}
	if c.MolarMass .Equal(decimal.Zero) {
		err := c.getMolarMass()
		if (err != nil) {
			return err
		}	
	}
	c.Moles = mass.Div(c.MolarMass)
	return nil
}

func (v Volume) getMoles(molarity decimal.Decimal) (decimal.Decimal, error) {
	standardVol, err := v.convertToStandard()
	handleError(err)
	return standardVol.Mul(molarity), err
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
	if compound.MolarMass .Equal(decimal.Zero) {
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
	var totalMass decimal.Decimal = decimal.Zero
	for _, em := range compound.Elements {
		totalMass = totalMass.Add(em.Element.AtomicWeight.Mul(em.Moles))
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
