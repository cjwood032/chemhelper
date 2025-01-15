package element


func (element *Element) getMolesFromMass(mass Mass) (moles float64){
	return convertMassToGrams(mass) / element.AtomicWeight 
}

func getMolesFromVolume(volume Volume, molarity float64) (moles float64){
	return convertVolumeToLiters(volume) * molarity
}

