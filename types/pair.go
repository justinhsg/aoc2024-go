package types

import "github.com/aoc-2024-go/utils"

type Pair[T any, U any] struct {
	Fst T
	Snd U
}

type IntPair Pair[int, int]

func (p IntPair) Destruct() (int, int) {
	return p.Fst, p.Snd
}

func AddPair(p1 IntPair, p2 IntPair) IntPair {
	return IntPair{p1.Fst + p2.Fst, p1.Snd + p2.Snd}
}

func (p IntPair) MultScalar(x int) IntPair {
	return IntPair{p.Fst * x, p.Snd * x}
}

func DiffPair(p1 IntPair, p2 IntPair) IntPair {
	return IntPair{p1.Fst - p2.Fst, p1.Snd - p2.Snd}
}

func DistPair(p1 IntPair, p2 IntPair) int {
	diffPair := DiffPair(p1, p2)
	return utils.Abs(diffPair.Fst) + utils.Abs(diffPair.Snd)
}

func NewIntPair(fst int, snd int) IntPair {
	return IntPair{fst, snd}
}

func SortByFstInt(x IntPair, y IntPair) int {
	if x.Fst == y.Fst {
		return x.Snd - y.Snd
	} else {
		return x.Fst - y.Fst
	}
}

func SortBySndInt(x IntPair, y IntPair) int {
	if x.Snd == y.Snd {
		return x.Fst - y.Fst
	} else {
		return x.Snd - y.Snd
	}
}
