package problem2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpt(t *testing.T) {
	opts := []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}

	retOpts := RunOptsCodes(opts)
	assert.Equal(t, 3500, retOpts[0])
}

func TestOpt1(t *testing.T) {
	opts := []int{1, 0, 0, 0, 99}

	retOpts := RunOptsCodes(opts)
	assert.Equal(t, 2, retOpts[0])
}

func TestOpt2(t *testing.T) {
	opts := []int{2, 3, 0, 3, 99}

	retOpts := RunOptsCodes(opts)
	assert.Equal(t, 6, retOpts[3])
}

func TestOpt3(t *testing.T) {
	opts := []int{2, 4, 4, 5, 99, 0}

	retOpts := RunOptsCodes(opts)
	assert.Equal(t, 9801, retOpts[5])
}

func TestOpt4(t *testing.T) {
	opts := []int{1, 1, 1, 4, 99, 5, 6, 0, 99}

	retOpts := RunOptsCodes(opts)
	assert.Equal(t, 30, retOpts[0])
}
