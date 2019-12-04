package main

// CalculateFuelRequired calculates fuel for given mass
func CalculateFuelRequired(mass int) int {
	return (mass / 3) - 2
}

// SumFuelForMasses calculates total fuel fot the given input of masses
func SumFuelForMasses(masses []int) int {
	sum := 0
	for _, mass := range masses {
		sum += CalculateFuelRequired(mass)
	}

	return sum
}

func main() {
}
