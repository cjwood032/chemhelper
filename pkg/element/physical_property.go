package element

//the iota is equal to the conversion to grams
type MassUnit float64
const (
gram MassUnit = 1
ounce MassUnit = 28.349
pound MassUnit = 453.592
kilogram MassUnit = 1000
hectogram MassUnit = 100
dekagram MassUnit = 10
decigram MassUnit = 1/10
centigram MassUnit = 1/100
milligram MassUnit = 1/1000
)
type VolumeUnit float64
const (
	liter VolumeUnit = 1
	deciliter VolumeUnit = 1/10
	centiliter VolumeUnit = 1/100
	milliliter VolumeUnit = 1/1000

)

type Mass struct{
	value float64
	unit MassUnit 
}

type Volume struct {
	value float64
	unit VolumeUnit 
}
//grams is the standard
func convertMassToGrams (m Mass) Mass{
	return Mass{value: m.value * float64(m.unit)}
}

func convertVolumeToLiters (v Volume) Volume{
	return Volume{value: v.value * float64(v.unit)}
}
