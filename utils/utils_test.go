package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains_ShouldReturnFalse(t *testing.T) {
	s := []string{"my", "array", "of", "string"}
	arrayContainsString := Contains(s, "arrays")
	assert.False(t, arrayContainsString)
}

func TestContains_ShouldReturnTrue(t *testing.T) {
	s := []string{"my", "array", "of", "string"}
	arrayContainsString := Contains(s, "of")
	assert.True(t, arrayContainsString)
}

func TestTrimLeftZeros(t *testing.T) {
	zeroPadedAddress := "0000000000000000000000004cf32fa4a59963379b419cb254e2bbbaccee35ea"
	expected := "04cf32fa4a59963379b419cb254e2bbbaccee35ea"

	actual := TrimLeftZeroes(zeroPadedAddress)

	assert.EqualValues(t, expected, actual)
}
