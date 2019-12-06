package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// CalculateFuelRequiredForMassAndFuel calculates fuel for fuel for given mass.
func CalculateFuelRequiredForMassAndFuel(mass int) int {
	var fuelRequired = CalculateFuelRequired(mass)

	var totalFuelRequiredForMassAndFuel = 0

	for {
		totalFuelRequiredForMassAndFuel += fuelRequired

		if fuelRequired <= 6 {
			break
		}

		fuelRequired = CalculateFuelRequired(fuelRequired)
	}

	return totalFuelRequiredForMassAndFuel
}

// CalculateFuelRequired calculates fuel for given mass
func CalculateFuelRequired(mass int) int {
	return (mass / 3) - 2
}

// SumFuelForMasses calculates total fuel for the given input of masses
func SumFuelForMasses(masses []int) int {
	sum := 0
	for _, mass := range masses {
		sum += CalculateFuelRequired(mass)
	}

	return sum
}

// SumFuelForMassesAndFuel calculates total fuel for the given input of masses and each mass's fuel
func SumFuelForMassesAndFuel(masses []int) int {
	sum := 0
	for _, mass := range masses {
		sum += CalculateFuelRequiredForMassAndFuel(mass)
	}

	return sum
}

func main() {
	file, err := os.Open("day1-input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var masses = []int{}
	for _, line := range lines {
		var mass = 0

		mass, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println(err)
			return
		}
		masses = append(masses, mass)
	}

	var totalFuel = SumFuelForMassesAndFuel(masses)

	println("Total fuel is: ", totalFuel)
}
