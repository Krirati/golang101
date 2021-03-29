package grade_test

import (
	"golang101/grade"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Input_80_Should_return_A(t *testing.T) {
	score := 80.0

	actual := grade.Calculate(score)

	expected := "A"
	assert.Equal(t, expected, actual)
}

func Test_Input_79_Should_return_B_plus(t *testing.T) {
	score := 79.0

	actual := grade.Calculate(score)

	expected := "B+"
	assert.Equal(t, expected, actual)
}

func Test_Input_65_Should_return_C_plus(t *testing.T) {
	score := 65.0

	actual := grade.Calculate(score)

	expected := "C+"
	assert.Equal(t, expected, actual)
}

func Test_Input_60_Should_return_C(t *testing.T) {
	score := 60.0

	actual := grade.Calculate(score)

	expected := "C"
	assert.Equal(t, expected, actual)
}

func Test_Input_59_Should_return_D_plus(t *testing.T) {
	score := 59.0

	actual := grade.Calculate(score)

	expected := "D+"

	assert.Equal(t, expected, actual)
}

func Test_Input_99_should_return_A(t *testing.T) {
	score := 99.0

	actual := grade.Calculate(score)

	expected := "A"
	assert.Equal(t, expected, actual)

}
