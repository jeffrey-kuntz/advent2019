package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jeffrey-kuntz/advent2019/math"
	"github.com/jeffrey-kuntz/advent2019/problem1"
	"github.com/jeffrey-kuntz/advent2019/problem2"
	"github.com/jeffrey-kuntz/advent2019/problem3"
)

func main() {
	runProblem1 := flag.Bool("problem1", false, "run problem 1")
	runProblem2 := flag.Bool("problem2", false, "run problem 2")
	runProblem3 := flag.Bool("problem3", false, "run problem 3")

	flag.Parse()

	if *runProblem1 {
		weights, err := problem1.LoadModuleWeights("data/problem1.txt")
		if err != nil {
			fmt.Println("error loading module weights: ", err)
			os.Exit(0)
		}

		fuelReq := problem1.CalculateAllFuelPlusFuel(weights)
		totalFuel := math.Sum(fuelReq)

		fmt.Println("problem 1: total fuel: ", totalFuel)
	}

	if *runProblem2 {
		optValues, err := problem2.LoadOptComputer("data/problem2.txt")

		// As per instructions replace position 1 with 12 and position 2 with 2
		// 1202 error, part 1
		optValues[1] = 12
		optValues[2] = 2
		newOptValues, err := problem2.RunOptsCodes(optValues)
		if err != nil {
			fmt.Println("mailformed opt codes: ", err)
		}

		fmt.Println("problem 2: part 1: ", newOptValues)

		// day 2 part 2
		done := false
		programValue := 0
		for noun := 0; noun <= 99 && !done; noun++ {
			for verb := 0; verb <= 99 && !done; verb++ {
				fmt.Println("stuff: ", noun, verb)
				fmt.Println("opt val: ", optValues)
				optValues[1] = noun
				optValues[2] = verb

				newOptValues, err = problem2.RunOptsCodes(optValues)
				if err != nil {
					fmt.Println("error while running opts: ", err)
					os.Exit(-1)
				}

				if newOptValues[0] == 19690720 {
					programValue = (100 * noun) + verb
					done = true
				}
			}
		}

		fmt.Println("problem 2: part 2: ", programValue)
	}

	// Problem 3
	if runProblem3 {
		wire1, wire2, err := problem3.LoadWires("data/problem3.txt")
		if err != nil {
			fmt.Println("error while parsing wire data: ", err)
			os.Exit(-1)
		}

		origin := math.NewPoint(0, 0)

		firstPoints, err := problem3.GeneratePoints(origin, wire1)
		if err != nil {
			fmt.Println("error while parsing wire data: ", err)
			os.Exit(-1)
		}

		secondPoints, err := problem3.GeneratePoints(origin, wire2)
		if err != nil {
			fmt.Println("error while parsing wire data: ", err)
			os.Exit(-1)
		}

		intersections := problem3.Intersection(firstPoints, secondPoints)
		distances := problem3.Distances(origin, intersections)

		fmt.Println("problem 3: ", distances[0])
	}
}
