package slices

func Map[S1 ~[]E1, S2 ~[]E2, E1 any, E2 any](src S1, dstTmpl S2, f func(E1) E2) S2 {
	var dst S2

	for _, e := range src {
		dst = append(dst, f(e))
	}

	return dst
}
