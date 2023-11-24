package lists

import (
	"sort"
	"strconv"
	"strings"
)

// Map returns a new slice containing the results of applying fn to each
func Map[T any, V any](arr []T, fn func(v T, i int) V) []V {
	res := make([]V, len(arr))
	for i, v := range arr {
		res[i] = fn(v, i)
	}
	return res
}

// Reduce returns a new slice containing the results of applying fn to each
func Reduce[T any, V any](arr []T, fn func(v T, i int, result V) V, result V) V {
	for i, v := range arr {
		result = fn(v, i, result)
	}
	return result
}

// Filter returns a new slice containing the results of applying fn to each, keep filter(v, i) == true
func Filter[T any](arr []T, filter func(v T, i int) bool) []T {
	var res []T
	for i, v := range arr {
		if filter(v, i) {
			res = append(res, v)
		}
	}
	return res
}

// Sort sorts the slice according to the fileds function.
// less: return order of fields(like sql order by), like na, asc,desc, [0,-1,1]
func Sort(arr any, less func(i, j int) []int) {
	sort.Slice(arr, func(i, j int) bool {
		result := less(i, j)
		for _, v := range result {
			if v == 0 {
				continue
			}
			return v < 0
		}
		return false
	})
}

func Int64s2Strs(arr []int64) []string {
	r := make([]string, len(arr))
	for i, v := range arr {
		r[i] = strconv.FormatInt(v, 10)
	}
	return r
}

func Strs2Int64s(arr []string) ([]int64, error) {
	r := make([]int64, len(arr))
	var err error
	for i, v := range arr {
		r[i], err = strconv.ParseInt(strings.TrimSpace(v), 10, 64)
		if err != nil {
			return nil, err
		}
	}
	return r, nil
}
func Unique[T comparable](vs []T) []T {
	m := make(map[T]bool, len(vs))
	var r []T
	for _, v := range vs {
		if m[v] {
			continue
		}
		r = append(r, v)
		m[v] = true

	}
	return r
}

// Concat arrays with copy
func Concat[T any](arrs ...[]T) []T {
	n := 0
	for _, arr := range arrs {
		n += len(arr)
	}
	res := make([]T, n)
	i := 0
	for _, arr := range arrs {
		copy(res[i:], arr)
		i += len(arr)
	}
	return res
}

func Subtract[T comparable](a []T, b []T) []T {
	if len(a) == 0 {
		return nil
	}
	m := AsBoolMap(b)
	a = Unique(a)
	var r []T
	for _, v := range a {
		if !m[v] {
			r = append(r, v)
		}
	}
	return r
}

func Union[T comparable](vs ...[]T) []T {
	var r []T
	for _, v := range vs {
		r = append(r, v...)
	}
	return Unique(r)
}

func Intersection[T comparable](a, b []T) []T {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}
	m := AsBoolMap(b)
	a = Unique(a)
	var r []T
	for _, v := range a {
		if m[v] {
			r = append(r, v)
		}
	}
	return r
}

func Paging[T any](page []T, pageSize int) [][]T {
	if len(page) == 0 {
		return nil
	}
	var r [][]T
	for i := 0; i < len(page); i += pageSize {
		end := i + pageSize
		if end > len(page) {
			end = len(page)
		}
		r = append(r, page[i:end])
	}
	return r
}
func CountPages(total, pageSize int) int {
	if total == 0 || pageSize == 0 {
		return 0
	}
	return (total + pageSize - 1) / pageSize
}

func GroupBy[T any, E comparable](arr []T, by func(v T) E) map[E][]T {
	r := make(map[E][]T)
	var key E
	for _, v := range arr {
		key = by(v)
		if b, ok := r[key]; ok {
			r[key] = append(b, v)
		} else {
			r[key] = []T{v}
		}
	}
	return r
}

// GroupByOne group by one value
func GroupByOne[T any, E comparable](arr []T, keepFirst bool, by func(v T) E) map[E]T {
	r := make(map[E]T)
	var key E
	for _, v := range arr {
		key = by(v)
		if _, ok := r[key]; ok {
			if !keepFirst {
				r[key] = v
			}
		}
	}
	return r
}

func Contains[T comparable](arr []T, value T) bool {
	return Index(arr, value) >= 0
}

// Index returns the index of the first instance of value in arr, or -1 if value is not present in arr.
func Index[T comparable](arr []T, value T) int {
	for i, v := range arr {
		if v == value {
			return i
		}
	}
	return -1
}

func IndexBy[T any](arr []T, fn func(v T, i int) bool) int {
	for i, v := range arr {
		if fn(v, i) {
			return i
		}
	}
	return -1
}

// Count returns the count of value in arr
func Count[T comparable](arr []T, value T) int {
	var count int
	for _, v := range arr {
		if v == value {
			count++
		}
	}
	return count
}

// CountBy returns the count of value in arr by given function
func CountBy[T any](arr []T, fn func(v T, i int) bool) int {
	var count int
	for i, v := range arr {
		if fn(v, i) {
			count++
		}
	}
	return count
}

// Any Returns true if arr contain any true value
func Any(arr []bool) bool {
	for _, v := range arr {
		if v {
			return true
		}
	}
	return false
}

// Any Returns true if all values are true
func All(arr []bool) bool {
	for _, v := range arr {
		if !v {
			return false
		}
	}
	return true
}

// Reindex 按indexes设置索引
func Reindex[T any](arr []T, indexes []int) {
	for i, v := range indexes {
		arr[i], arr[v] = arr[v], arr[i]
	}
}

// AsBoolMap arr -> [value]bool
func AsBoolMap[T comparable](vs []T) map[T]bool {
	m := make(map[T]bool, len(vs))
	for _, v := range vs {
		m[v] = true
	}
	return m
}

// ValueIndex arr -> [value]index
func AsIndexMap[T comparable](arr []T) map[T]int {
	r := make(map[T]int, len(arr))
	for i, v := range arr {
		r[v] = i
	}
	return r
}

// Count arr -> [value]count
func AsCountMap[T comparable](arr []T) map[T]int {
	r := make(map[T]int, len(arr))
	for _, v := range arr {
		r[v]++
	}
	return r
}
