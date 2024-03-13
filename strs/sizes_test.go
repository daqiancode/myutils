package strs_test

import (
	"testing"

	"github.com/daqiancode/myutils/strs"
	"github.com/stretchr/testify/assert"
)

func TestSizes(t *testing.T) {
	r, err := strs.ParseSize("80x80")
	assert.Nil(t, err)
	assert.Equal(t, []int{80, 80}, r)
	r1, err := strs.ParseSize("80")
	assert.Nil(t, err)
	assert.Equal(t, []int{80, 80}, r1)
}
