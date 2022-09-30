package main

// "sort"
// "container/heap"

import (
	"fmt"
	"os"
)

type TruthBooth struct {
	Match   Match
	Perfect bool
}

type MatchCeremony struct {
	Matches MatchSet
	Correct int
}

type Set[T comparable] interface {
	Add(T)
	Delete(T)
	Has(T) bool
}

type Condition interface {
	IsMet(MatchSet) bool
}

func (c TruthBooth) IsMet(ms MatchSet) bool {
	if ms.Has(c.Match) {
		return c.Perfect
	}
	return !c.Perfect
}

func (c MatchCeremony) IsMet(ms MatchSet) bool {
	nSameMatches := 0
	for m := range ms {
		if c.Matches.Has(m) {
			nSameMatches += 1
		}
	}
	return nSameMatches == c.Correct
}

// func SortMapByFunc[K any, V any](a map[K]V, f func(K, V) float64) []map[K]V {
// 	keys := make([]K, 0, len(a))
// 	values := make([]V, 0, len(a))
// 	for k := range a {
// 		keys = append(keys, k)
// 	}
// 	keys.Sort()

// 	out := make([]map[K]V, len(keys))
// 	for i, key := range keys {
// 		out[i] = map[K]V{key: a[key]}
// 	}
// 	return out
// }

func main() {
	jsonData, err := os.ReadFile("data/uk.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	showData := ParseShowData(jsonData)
	var conditions []Condition
	// conditions = append(conditions, showData.TruthBooths...)
	// ^ this doesn't work because Go is dumb, no idea why
	for _, tb := range showData.TruthBooths {
		conditions = append(conditions, tb)
	}
	for _, mc := range showData.MatchCeremonies {
		conditions = append(conditions, mc)
	}

	possibleMatches := GetPossibleMatches(
		showData.Males,
		showData.Females,
		conditions,
	)
	PrintStats(possibleMatches, showData.Males, showData.Females)

	// males := []string{"Joshy", "Jack", "Charlie", "Jacob", "Jordan", "Theo", "Josh", "Juan", "Cach", "Ismail"}
	// females := []string{"Tasha", "Shae", "Thea", "Vic", "Robyn", "Libby", "Taofiqah", "Tersea", "Sapphia", "Olivia"}
	// conditions := []Condition{
	// 	TruthBooth{Match{"Joshy", "Tasha"}, false},
	// 	TruthBooth{Match{"Josh", "Taofiqah"}, false},
	// 	TruthBooth{Match{"Theo", "Robyn"}, false},
	// 	TruthBooth{Match{"Juan", "Thea"}, false},
	// 	TruthBooth{Match{"Ismail", "Olivia"}, true},
	// 	TruthBooth{Match{"Cach", "Thea"}, false},
	// 	TruthBooth{Match{"Charlie", "Thea"}, false},
	// 	MatchCeremony{
	// 		NewMatchSet([]Match{
	// 			{"Jack", "Shae"},
	// 			{"Charlie", "Thea"},
	// 			{"Jacob", "Vic"},
	// 			{"Jordan", "Robyn"},
	// 			{"Joshy", "Libby"},
	// 			{"Theo", "Tasha"},
	// 			{"Josh", "Taofiqah"},
	// 			{"Juan", "Tersea"},
	// 			{"Cach", "Sapphia"},
	// 			{"Ismail", "Olivia"},
	// 		}),
	// 		2,
	// 	},
	// 	MatchCeremony{
	// 		NewMatchSet([]Match{
	// 			{"Jacob", "Vic"},
	// 			{"Cach", "Taofiqah"},
	// 			{"Charlie", "Shae"},
	// 			{"Joshy", "Robyn"},
	// 			{"Josh", "Libby"},
	// 			{"Theo", "Tasha"},
	// 			{"Juan", "Thea"},
	// 			{"Jack", "Tersea"},
	// 			{"Jordan", "Sapphia"},
	// 			{"Ismail", "Olivia"},
	// 		}),
	// 		3,
	// 	},
	// 	MatchCeremony{
	// 		NewMatchSet([]Match{
	// 			{"Jacob", "Vic"},
	// 			{"Cach", "Sapphia"},
	// 			{"Charlie", "Shae"},
	// 			{"Joshy", "Robyn"},
	// 			{"Josh", "Tersea"},
	// 			{"Theo", "Taofiqah"},
	// 			{"Juan", "Thea"},
	// 			{"Jack", "Libby"},
	// 			{"Jordan", "Tasha"},
	// 			{"Ismail", "Olivia"},
	// 		}),
	// 		2,
	// 	},
	// 	MatchCeremony{
	// 		NewMatchSet([]Match{
	// 			{"Joshy", "Vic"},
	// 			{"Theo", "Sapphia"},
	// 			{"Cach", "Shae"},
	// 			{"Jacob", "Robyn"},
	// 			{"Josh", "Tersea"},
	// 			{"Jack", "Taofiqah"},
	// 			{"Charlie", "Thea"},
	// 			{"Jordan", "Libby"},
	// 			{"Juan", "Tasha"},
	// 			{"Ismail", "Olivia"},
	// 		}),
	// 		3,
	// 	},
	// 	MatchCeremony{
	// 		NewMatchSet([]Match{
	// 			{"Joshy", "Robyn"},
	// 			{"Theo", "Libby"},
	// 			{"Cach", "Taofiqah"},
	// 			{"Jacob", "Vic"},
	// 			{"Josh", "Thea"},
	// 			{"Jack", "Tersea"},
	// 			{"Charlie", "Sapphia"},
	// 			{"Jordan", "Shae"},
	// 			{"Juan", "Tasha"},
	// 			{"Ismail", "Olivia"},
	// 		}),
	// 		3,
	// 	},
	// 	MatchCeremony{
	// 		NewMatchSet([]Match{
	// 			{"Jordan", "Robyn"},
	// 			{"Charlie", "Libby"},
	// 			{"Juan", "Taofiqah"},
	// 			{"Theo", "Vic"},
	// 			{"Joshy", "Thea"},
	// 			{"Jacob", "Tersea"},
	// 			{"Jack", "Sapphia"},
	// 			{"Cach", "Shae"},
	// 			{"Josh", "Tasha"},
	// 			{"Ismail", "Olivia"},
	// 		}),
	// 		2,
	// 	},
	// 	MatchCeremony{
	// 		NewMatchSet([]Match{
	// 			{"Jordan", "Taofiqah"},
	// 			{"Charlie", "Shae"},
	// 			{"Juan", "Tersea"},
	// 			{"Theo", "Libby"},
	// 			{"Joshy", "Vic"},
	// 			{"Jacob", "Thea"},
	// 			{"Jack", "Robyn"},
	// 			{"Cach", "Sapphia"},
	// 			{"Josh", "Tasha"},
	// 			{"Ismail", "Olivia"},
	// 		}),
	// 		1,
	// 	},
	// }

	// possibleMatchSets := GetPossibleMatches(males, females, conditions)
	// PrintStats(possibleMatchSets, males, females)
}
