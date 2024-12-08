package types

type Pair[T any, U any] struct {
	Fst T
	Snd U
}

type IntPair Pair[int, int]

func AddPair(p1 IntPair, p2 IntPair) IntPair {
	return IntPair{p1.Fst + p2.Fst, p1.Snd + p2.Snd}
}

func DiffPair(p1 IntPair, p2 IntPair) IntPair {
	return IntPair{p1.Fst - p2.Fst, p1.Snd - p2.Snd}
}

func NewIntPair(fst int, snd int) IntPair {
	return IntPair{fst, snd}
}
