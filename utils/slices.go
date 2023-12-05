package utils

func Contains[S []T, T comparable](slice S, element T) bool {
	for _, value := range slice {
		if value == element {
			return true
		}
	}
	return false
}
