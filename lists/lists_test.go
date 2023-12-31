package lists_test

import (
	"testing"

	"github.com/daqiancode/myutils/lists"
	"github.com/stretchr/testify/assert"
)

func TestUnique(t *testing.T) {
	a := []int{}
	r := lists.Unique(a)
	assert.Nil(t, r)
	a1 := []int{1, 2, 3, 3, 4, 4}
	r1 := lists.Unique(a1)
	assert.EqualValues(t, []int{1, 2, 3, 4}, r1)
}

func TestIntersection(t *testing.T) {
	a := []int{1, 2, 3, 3, 4, 4}
	b := []int{4}
	r := lists.Intersection(a, b)
	assert.EqualValues(t, []int{4}, r)

}

func TestSubtract(t *testing.T) {
	a := []int{1, 2, 3, 3, 4, 4}
	b := []int{1}
	r := lists.Subtract(a, b)
	assert.EqualValues(t, []int{2, 3, 4}, r)

}

func TestUnion(t *testing.T) {
	a := []int{1, 2, 3, 3, 4, 4}
	b := []int{5, 6}
	r := lists.Union(a, b)
	assert.EqualValues(t, []int{1, 2, 3, 4, 5, 6}, r)
}

func TestPaging(t *testing.T) {
	a := []int{1, 2, 3, 3, 4}
	pages := lists.Paging(a, 3)
	assert.EqualValues(t, [][]int{{1, 2, 3}, {3, 4}}, pages)
}

func TestCountPages(t *testing.T) {
	count := lists.CountPages(2, 3)
	assert.EqualValues(t, 1, count)
	count = lists.CountPages(10, 3)
	assert.EqualValues(t, 4, count)
}

func TestSort(t *testing.T) {
	a := []int{100, 2, 12, 2, 34}
	lists.Sort(a, func(i, j int) []int { return []int{a[i] - a[j]} })
	assert.EqualValues(t, []int{2, 2, 12, 34, 100}, a)
}
