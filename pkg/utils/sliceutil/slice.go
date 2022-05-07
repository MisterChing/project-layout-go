package sliceutil

func InStringSlice(needle string, haystack []string) bool {
	for _, v := range haystack {
		if needle == v {
			return true
		}
	}
	return false
}

func InIntSlice(needle int, haystack []int) bool {
	for _, v := range haystack {
		if needle == v {
			return true
		}
	}
	return false
}
