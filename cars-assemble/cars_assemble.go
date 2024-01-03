package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(float64(productionRate/60) * successRate / 100)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var totalCost uint
	groupCost := 95000
	eachCost := 10000

	if carsCount > 9 {
		group := carsCount / 10
		each := carsCount - group*10

		totalCost = uint(group*groupCost + each*eachCost)
		return totalCost
	}

	totalCost = uint(carsCount * eachCost)
	return totalCost
}
