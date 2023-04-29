package test

func complexArray(pair ...int) []complex128 {
	arr := make([]complex128, len(pair)/2)
	for i := 1; i < len(pair); i += 2 {
		arr[i/2] = complex(float64(pair[i-1]), float64(pair[i]))
	}
	return arr
}

func compareArrays[T comparable](arr1, arr2 []T) bool {
	if len(arr1) != len(arr2) {
		return false
	} else {
		for i, v := range arr1 {
			if arr2[i] != v {
				return false
			}
		}
		return true
	}
}
