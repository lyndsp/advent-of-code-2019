package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// CalculateFuelRequiredForFuel calculates fuel for fuel for given mass.
func CalculateFuelRequiredForFuel(mass int) int {
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

// SumFuelForMasses calculates total fuel fot the given input of masses
func SumFuelForMasses(masses []int) int {
	sum := 0
	for _, mass := range masses {
		sum += CalculateFuelRequired(mass)
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

	var totalFuel = SumFuelForMasses(masses)

	println("Total fuel is: ", totalFuel)
}
