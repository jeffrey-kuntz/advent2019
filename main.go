package main

import (
	"fmt"
	"os"

	"github.com/jeffrey-kuntz/advent2019/math"
	"github.com/jeffrey-kuntz/advent2019/problem1"
	"github.com/jeffrey-kuntz/advent2019/problem2"
)

func main() {
	weights, err := problem1.LoadModuleWeights("data/problem1.txt")
	if err != nil {
		fmt.Println("error loading module weights: ", err)
		os.Exit(0)
	}

	fuelReq := problem1.CalculateAllFuelPlusFuel(weights)
	totalFuel := math.Sum(fuelReq)

	fmt.Println("problem 1: total fuel: ", totalFuel)

	optValues, err := problem2.LoadOptComputer("data/problem2.txt")

	// As per instructions replace position 1 with 12 and position 2 with 2
	// 1202 error, part 1
	optValues[1] = 12
	optValues[2] = 2
	newOptValues := problem2.RunOptsCodes(optValues)

	fmt.Println("problem 2: final opt: ", newOptValues)

	// day 2 part 2
	done := false
	programValue := 0
	for noun := 0; noun <= 99 && !done; noun++ {
		for verb := 0; verb <= 99 && !done; verb++ {
			fmt.Println("stuff: ", noun, verb)
			fmt.Println("opt val: ", optValues)
			optValues[1] = noun
			optValues[2] = verb

			newOptValues = problem2.RunOptsCodes(optValues)
			if newOptValues[0] == 19690720 {
				programValue = (100 * noun) + verb
				done = true
			}
		}
	}

	fmt.Println("problem 2: part 2: ", programValue)

}
