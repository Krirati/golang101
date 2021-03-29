package average_test

import (
	"golang101/grade"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAverage_1_2_3_5_Should_return_2_75(t *testing.T) {
	// arrange
	input := []float64{1, 2, 3, 5}

	//action
	actual := grade.Average(input)

	//assery
	expected := 2.75
	assert.Equal(t, expected, actual)
}

func TestAvertage_3_3_3_3_should_return_somthing(t *testing.T) {
	// arrange
	input := []float64{3, 3, 3, 3}
	//action
	actual := grade.Average(input)
	expected := 3.0

	assert.Equal(t, expected, actual)
}
