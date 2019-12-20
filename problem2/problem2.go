package problem2

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// LoadOptComputer loads the opt string from a file
func LoadOptComputer(filename string) ([]int, error) {
	infile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(infile)
	scanner.Scan()
	optStrings := strings.Split(scanner.Text(), ",")

	opts := make([]int, len(optStrings))
	for index, opt := range optStrings {
		intOpt, err := strconv.Atoi(opt)
		if err != nil {
			return nil, err
		}

		opts[index] = intOpt
	}

	return opts, nil
}

// RunOptsCodes runs the codes
func RunOptsCodes(inCodes []int) ([]int, error) {
	codes := append([]int(nil), inCodes...)

	for i := 0; i < len(codes); i += 4 {
		if codes[i] == 99 {
			// Done
			fmt.Println("done")
			return codes, nil
		}

		position1 := codes[i+1]
		value1 := codes[position1]

		position2 := codes[i+2]
		value2 := codes[position2]

		returnPosition := codes[i+3]

		switch codes[i] {
		case 1:
			// Addition
			codes[returnPosition] = value1 + value2
		case 2:
			// Multiplication
			codes[returnPosition] = value1 * value2
		default:
			return nil, errors.New("unrecognized code: " + string(codes[i]))
		}
	}

	return codes, nil
}
