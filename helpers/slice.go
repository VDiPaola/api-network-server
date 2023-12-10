package helpers

func Remove[t any](s []t, i int) []t {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
