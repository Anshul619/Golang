package main

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	var out Ints

	for _, v := range i {
		if filter(v) {
			out = append(out, v)
		}
	}

	return out
}

func (i Ints) Discard(filter func(int) bool) Ints {
	var out Ints

	for _, v := range i {
		if !filter(v) {
			out = append(out, v)
		}
	}

	return out
}

func (l Lists) Keep(filter func([]int) bool) Lists {

	var out Lists

	for _, v := range l {
		if filter(v) {
			out = append(out, v)
		}
	}

	return out
}

func (s Strings) Keep(filter func(string) bool) Strings {
	var out Strings

	for _, v := range s {
		if filter(v) {
			out = append(out, v)
		}
	}

	return out
}
