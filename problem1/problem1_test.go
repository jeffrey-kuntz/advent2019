package problem1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWeight(t *testing.T) {
	fuel := CalculateFuel(14)
	assert.Equal(t, 2, fuel)
}

func TestFuel(t *testing.T) {
	fuel := CalculateFuelPlusFuel(14)
	assert.Equal(t, 2, fuel)
}

func TestFuel1(t *testing.T) {
	fuel := CalculateFuelPlusFuel(1969)
	assert.Equal(t, 966, fuel)
}

func TestFuel2(t *testing.T) {
	fuel := CalculateFuelPlusFuel(100756)
	assert.Equal(t, 50346, fuel)
}
