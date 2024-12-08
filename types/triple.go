package types

type Triple[T any, U any, V any] struct {
	Fst T
	Snd U
	Thd V
}

type IntTriple Triple[int, int, int]

func NewIntTriple(fst int, snd int, thd int) IntTriple {
	return IntTriple{fst, snd, thd}
}
