package utils

func MergeMaps[M map[K]V, K comparable, V comparable](first, second map[K]V) map[K]V {
	merged := first
	for k, v := range second {
		merged[k] = v
	}

	return merged
}
