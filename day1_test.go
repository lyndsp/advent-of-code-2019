/*
The Fuel Counter-Upper needs to know the total fuel requirement.
To find it, individually calculate the fuel needed for the mass of each module (your puzzle input),
 then add together all the fuel values.
*/

package main

import (
	"fmt"
	"testing"
)

var fuelForMasstests = []struct {
	fuel int
	mass int
}{
	{2, 12},
	{2, 14},
	{654, 1969},
}

func TestDay1FuelForSingleModuleMass(t *testing.T) {

	for _, tt := range fuelForMasstests {
		t.Run(fmt.Sprint(tt.mass), func(t *testing.T) {
			out := CalculateFuelRequired(tt.mass)
			if out != tt.fuel {
				t.Errorf("got %v, want %v", out, tt.fuel)
			}
		})
	}
}

func TestDay1SumFuelForMasses(t *testing.T) {
	moduleMasses := [2]int{12, 14}

	got := SumFuelForMasses(moduleMasses)

	want := 4

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
