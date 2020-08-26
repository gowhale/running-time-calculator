package main

// go test -v running_time_calculator.go running_time_calculator_test.go

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTheTests(t *testing.T) {
	assert.EqualValues(t, 1, 1)
}

func TestCalculateAveragePaceFivekmRun(t *testing.T) {
	assert.EqualValues(t, "5:00 /km", calculateAveragePace(5, 25, 0))
}
