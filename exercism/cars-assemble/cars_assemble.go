package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) / 100 * successRate
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	if carsCount == 0 {
		return uint(0)
	}

	if carsCount < 10 {
		return uint(carsCount * 10000)
	} else {
		var tens = carsCount / 10
		var rest = carsCount % (tens * 10)
		var cost = (tens * 95000) + (rest * 10000)
		return uint(cost)
	}
}
