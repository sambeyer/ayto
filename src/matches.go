package main

type Match struct {
	Male string
	Female string
}

type MatchSet map[Match]struct{}

type PossibleMatches struct {
	all []MatchSet
	possible []int
}

func (ms *MatchSet) Add (m Match) {
	(*ms)[m] = struct{}{}
}

func(ms *MatchSet) Delete(m Match) {
	delete(*ms, m)
}

func(ms *MatchSet) Has(m Match) bool {
	_, ok := (*ms)[m]
	return ok
}

func NewMatchSet(ms []Match) MatchSet {
	out := MatchSet{}
	for _, m := range ms {
		out.Add(m)
	}
	return out
}
