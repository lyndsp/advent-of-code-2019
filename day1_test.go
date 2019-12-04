/*
The Fuel Counter-Upper needs to know the total fuel requirement.
To find it, individually calculate the fuel needed for the mass of each module (your puzzle input),
 then add together all the fuel values.

Fuel required to launch a given module is based on its mass.
Specifically, to find the fuel required for a module,
	take its mass, divide by three, round down, and subtract 2.
*/

package main

import "testing"

func TestDay1FuelForSingleModuleMass12(t *testing.T) {
	got := CalculateFuelRequired(12)

	want := 2

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
