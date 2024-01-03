package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, minute int) int {
	if minute == 0 {
		minute = 2
	}
	return len(layers) * minute
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodle int, sauce float64) {
	for _, qty := range layers {
		switch qty {
		case "noodles":
			noodle += 50
		case "sauce":
			sauce += 0.2
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	secret := friendsList[len(friendsList)-1]
	myList[len(myList)-1] = secret
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portion int) []float64 {
	calcQuantities := []float64{}
	calcPortion := float64(portion) / 2

	for _, qty := range quantities {
		calcQuantities = append(calcQuantities, qty*calcPortion)
	}

	return calcQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
