package slices

func Filter[S ~[]E, E any](src S, f func(E) bool) S {
	var dst S

	for _, e := range src {
		if f(e) {
			dst = append(dst, e)
		}
	}

	return dst
}
