package problem1

import (
	"bufio"
	"os"
	"strconv"
)

// CalculateAllFuelPlusFuel takes the weight of multipel modules and gets
// the fuel requirements
func CalculateAllFuelPlusFuel(weights []int) []int {
	var allFuel []int

	for _, weight := range weights {
		fuel := CalculateFuelPlusFuel(weight)
		allFuel = append(allFuel, fuel)
	}

	return allFuel
}

// CalculateFuelPlusFuel figures out how much fuel is needed given the
// weight of the module and the weight of the fuel
func CalculateFuelPlusFuel(weight int) int {
	fuel := ((weight / 3) - 2)
	if fuel < 0 {
		return 0
	}

	return fuel + CalculateFuelPlusFuel(fuel)
}

// CalculateAllFuel takes the weight of multipel modules and gets
// the fuel requirements
func CalculateAllFuel(weights []int) []int {
	var allFuel []int

	for _, weight := range weights {
		fuel := CalculateFuel(weight)
		allFuel = append(allFuel, fuel)
	}

	return allFuel
}

// CalculateFuel takes the weight and returns fuel requirements
func CalculateFuel(weight int) int {
	return ((weight / 3) - 2)
}

// LoadModuleWeights loads the data that is needed for problem 1
func LoadModuleWeights(filename string) ([]int, error) {
	infile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer infile.Close()

	var weights []int

	scanner := bufio.NewScanner(infile)
	for scanner.Scan() {
		weight, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}

		weights = append(weights, weight)
	}

	return weights, nil
}
