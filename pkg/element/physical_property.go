package element


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
	unknown MassUnit = "unknown"
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
	convertToStandard() float64
}

func (m Mass) convertToStandard() float64 {	
	
	switch m.unit {
	case pound:
		return m.value * 453.592 * float64(m.prefix)
	case ounce:
		return m.value * 28.349 * float64(m.prefix)
	default:
		return m.value * float64(m.prefix)
	}
}

func (v Volume) convertToStandard() float64 {
	return v.value * float64(v.unit)
}
func convertToStandardValue(p Property) float64 {
	
    return p.convertToStandard()
}
func (element *Element) getMolesFromMass(mass Mass) (moles float64){
	return mass.convertToStandard() / element.AtomicWeight 
}

func getMolesFromVolume(volume Volume, molarity float64) (moles float64){
	return volume.convertToStandard() * molarity
}
