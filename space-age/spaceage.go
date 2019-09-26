package space

// Planet is the planet on which we will calculate  age
type Planet string

// The duration of one complete revolution of earth in seconds
const earthRevolutionDuration = 31557600

// Age calculates the age relative to a planet rotation cycle
func Age(ageInSeconds float64, planet Planet) float64 {

	Planets := map[Planet]float64{
		"Earth":   earthRevolutionDuration,
		"Mercury": earthRevolutionDuration * 0.2408467,
		"Venus":   earthRevolutionDuration * 0.61519726,
		"Mars":    earthRevolutionDuration * 1.8808158,
		"Jupiter": earthRevolutionDuration * 11.862615,
		"Saturn":  earthRevolutionDuration * 29.447498,
		"Uranus":  earthRevolutionDuration * 84.016846,
		"Neptune": earthRevolutionDuration * 164.79132,
	}
	return calculateAge(Planets[planet], ageInSeconds)
}
func calculateAge(planetRevolution float64, ageInSeconds float64) float64 {
	return ageInSeconds / planetRevolution
}
