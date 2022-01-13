package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, minutes int) int {
	if minutes == 0 {
		minutes = 2
	}
	return len(layers) * minutes
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {

	for i := 0; i < len(layers); i++ {
		switch {
		case layers[i] == "noodles":
			noodles += 50
		case layers[i] == "sauce":
			sauce += 0.2
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portion int) []float64 {
	amounts := make([]float64, len(quantities))
	portionNeeded := float64(portion) / 2.0
	for i := 0; i < len(quantities); i++ {
		amounts[i] = quantities[i] * portionNeeded
	}
	return amounts
}
