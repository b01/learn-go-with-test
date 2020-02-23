package iteration

func Repeat(c string, n  int) string {
	var r string

	for i := 0; i < n; i++ {
		r += c
	}

	return r
}
