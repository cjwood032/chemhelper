package element


func (element *Element) getMolesFromMass(mass Mass) (moles float64){
	return convertMassToGrams(mass).value / element.AtomicWeight 
}

func getMolesFromVolume(volume Volume, molarity float64) (moles float64){
	return convertVolumeToLiters(volume).value * molarity
}

func (c Compound) MolarMass() float64 {
	totalMass := 0.0
	for _, em := range c.Elements {
		totalMass += em.Element.AtomicWeight * em.Moles
	}
	return totalMass
}