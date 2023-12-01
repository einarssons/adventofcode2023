package utils

func ReverseSlice[T any](s []T) {
	l := len(s)
	nrSwaps := l / 2
	for i := 0; i < nrSwaps; i++ {
		s[i], s[l-1-i] = s[l-1-i], s[i]
	}
}
