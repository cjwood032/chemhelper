package element

type Group int
const (
    Alkali Group = iota +1
    Alkaline
    Metal 
    Pnictogens = 15
    Chalcogens = 16
    Halogens = 17
    NobleGases = 18

)
func (g Group) String() string {
    switch {
    case g == Alkali:
        return "Alkali Metals"
    case g == Alkaline:
        return "Alkaline Earth Metals"
    case g >= Metal && g <= 14:
        return "Metal"
    case g == Pnictogens:
        return "Pnictogens"
    case g == Chalcogens:
        return "Chalcogens"
    case g == Halogens:
        return "Halogens"
    case g == NobleGases:
        return "Noble Gases"
    default:
        return "Unknown Group"
    }
}
