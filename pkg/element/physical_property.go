package element


type MetricPrefix float64

const (
	none MetricPrefix = 1
	kilo MetricPrefix = 1000
	hecto MetricPrefix = 100
	deca MetricPrefix = 10
	deci MetricPrefix = 0.1
	centi MetricPrefix = 0.01
	milli MetricPrefix = 0.001
	micro MetricPrefix = 0.000001
	nano MetricPrefix = 0.000000001
)
type MassUnit string

const (
	// Metric
	gram    MassUnit = "gram"
	// Imperial
	ounce   MassUnit = "ounce"
	pound   MassUnit = "pound"
)

type Mass struct {
	value float64
	unit  MassUnit
	prefix MetricPrefix
}


type Volume struct {
	value float64
	unit  MetricPrefix // e.g., centi, milli
}

type Convertible interface {
	convertToStandard() float64
}

func (m Mass) convertToStandard() float64 {	
	switch m.unit {
	case pound:
		return m.value * 453.592
	case ounce:
		return m.value * 28.349
	default:
		// Handle metric units with the prefix
		return m.value * float64(m.prefix) // Convert to grams by applying the metric prefix
	}
}

func (v Volume) convertToStandard() float64 {
	// Convert volume to liters by applying the metric prefix
	return v.value * float64(v.unit)
}

func convertMassToGrams(m Mass) float64 {
	return m.convertToStandard()
}

func convertVolumeToLiters(v Volume) float64 {
	return v.convertToStandard()
}

func convertToStandardValue(c Convertible) float64 {
	return c.convertToStandard()
}


func (element *Element) getMolesFromMass(mass Mass) (moles float64){
	return convertMassToGrams(mass) / element.AtomicWeight 
}

func getMolesFromVolume(volume Volume, molarity float64) (moles float64){
	return convertVolumeToLiters(volume) * molarity
}