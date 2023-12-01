package utils

func Sort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := i; j < len(a); j++ {
			if a[j] < a[i] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}
