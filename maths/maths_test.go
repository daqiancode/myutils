package maths_test

import (
	"testing"

	"github.com/daqiancode/myutils/maths"
	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	r := maths.Min([]int{1, 2, 3})
	assert.Equal(t, 1, r)
	r = maths.Max([]int{1, 2, 3})
	assert.Equal(t, 3, r)
}
