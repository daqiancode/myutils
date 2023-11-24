package rands_test

import (
	"testing"

	"github.com/daqiancode/myutils/rands"
	"github.com/stretchr/testify/assert"
)

func TestRandom(t *testing.T) {
	assert.NotEqual(t, rands.Random(32), rands.Random(32))
	assert.NotEqual(t, rands.RandomHex(32), rands.RandomHex(32))
	assert.NotEqual(t, rands.RandomLower(32), rands.RandomLower(32))
	assert.NotEqual(t, rands.RandomNumber(32), rands.RandomNumber(32))
}
