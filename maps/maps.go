package maps

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func Update[K comparable, V any](m map[K]V, values map[K]V) map[K]V {
	for k, v := range values {
		m[k] = v
	}
	return m
}

func Merge[K comparable, V any](ms ...map[K]V) map[K]V {
	r := make(map[K]V)
	for _, m := range ms {
		for k, v := range m {
			r[k] = v
		}
	}
	return r
}
func MergeWith[K comparable, V any](f func(K, V, V) V, ms ...map[K]V) map[K]V {
	r := make(map[K]V)
	for _, m := range ms {
		for k, v := range m {
			if ov, ok := r[k]; ok {
				r[k] = f(k, ov, v)
			} else {
				r[k] = v
			}
		}
	}
	return r
}

func Filter[K comparable, V any](m map[K]V, f func(K, V) bool) map[K]V {
	r := make(map[K]V)
	for k, v := range m {
		if f(k, v) {
			r[k] = v
		}
	}
	return r
}

func Map[K comparable, V1, V2 any](m map[K]V1, f func(K, V1) V2) map[K]V2 {
	r := make(map[K]V2)
	for k, v := range m {
		r[k] = f(k, v)
	}
	return r
}

func ToList[K comparable, V any](m map[K]V) ([]K, []V) {
	return Keys(m), Values(m)
}

func ToMap[K comparable, V any](keys []K, values []V) map[K]V {
	r := make(map[K]V)
	for i, k := range keys {
		r[k] = values[i]
	}
	return r
}
