package xslices

func Map[S1 ~[]E1, S2 []E2, E1 any, E2 any](s1 S1, fn func(E1) E2) S2 {
	s2 := make(S2, len(s1), cap(s1))
	for i := range s1 {
		s2[i] = fn(s1[i])
	}
	return s2
}
