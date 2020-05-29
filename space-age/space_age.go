package space


// Planet is string
type Planet string

var continent = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}
//Age is exported
func Age(sec float64, p Planet) float64 {
	var earthSeconds float64 = float64(31557600)
	if p == "Earth" {
		return sec / earthSeconds
	}
	return (sec / earthSeconds) / continent[p]

}
