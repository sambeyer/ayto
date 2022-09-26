package main

import (
	"fmt"
	// "sort"
	// "container/heap"
)

type MatchBooth struct {
	Match Match
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

func (c MatchBooth) IsMet(ms MatchSet) bool {
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
	males := []string{"Joshy", "Jack", "Charlie", "Jacob", "Jordan", "Theo", "Josh", "Juan", "Cach", "Ismail"}
	females := []string{"Tasha", "Shae", "Thea", "Vic", "Robyn", "Libby", "Taofiqah", "Tersea", "Sapphia", "Olivia"}
	nextMaleOrdering := Permutations(males)
	nextMatches := func() MatchSet {
		newMales := nextMaleOrdering()
		if newMales == nil {
			return nil
		}
		matches := make([]Match, len(newMales))
		for i, male := range newMales {
			matches[i] = Match{male, females[i]}
		}
		return NewMatchSet(matches)
	}
	conditions := []Condition{
		MatchBooth{Match{"Joshy", "Tasha"}, false},
		MatchBooth{Match{"Josh", "Taofiqah"}, false},
		MatchBooth{Match{"Theo", "Robyn"}, false},
		MatchBooth{Match{"Juan", "Thea"}, false},
		MatchBooth{Match{"Ismail", "Olivia"}, true},
		MatchBooth{Match{"Cach", "Thea"}, false},
		MatchBooth{Match{"Charlie", "Thea"}, false},
		MatchCeremony{
			NewMatchSet([]Match{
				Match{"Jack", "Shae"},
				Match{"Charlie", "Thea"},
				Match{"Jacob", "Vic"},
				Match{"Jordan", "Robyn"},
				Match{"Joshy", "Libby"},
				Match{"Theo", "Tasha"},
				Match{"Josh", "Taofiqah"},
				Match{"Juan", "Tersea"},
				Match{"Cach", "Sapphia"},
				Match{"Ismail", "Olivia"},
			}),
			2,
		},
		MatchCeremony{
			NewMatchSet([]Match{
				Match{"Jacob", "Vic"},
				Match{"Cach", "Taofiqah"},
				Match{"Charlie", "Shae"},
				Match{"Joshy", "Robyn"},
				Match{"Josh", "Libby"},
				Match{"Theo", "Tasha"},
				Match{"Juan", "Thea"},
				Match{"Jack", "Tersea"},
				Match{"Jordan", "Sapphia"},
				Match{"Ismail", "Olivia"},
			}),
			3,
		},
		MatchCeremony{
			NewMatchSet([]Match{
				Match{"Jacob", "Vic"},
				Match{"Cach", "Sapphia"},
				Match{"Charlie", "Shae"},
				Match{"Joshy", "Robyn"},
				Match{"Josh", "Tersea"},
				Match{"Theo", "Taofiqah"},
				Match{"Juan", "Thea"},
				Match{"Jack", "Libby"},
				Match{"Jordan", "Tasha"},
				Match{"Ismail", "Olivia"},
			}),
			2,
		},
		MatchCeremony{
			NewMatchSet([]Match{
				Match{"Joshy", "Vic"},
				Match{"Theo", "Sapphia"},
				Match{"Cach", "Shae"},
				Match{"Jacob", "Robyn"},
				Match{"Josh", "Tersea"},
				Match{"Jack", "Taofiqah"},
				Match{"Charlie", "Thea"},
				Match{"Jordan", "Libby"},
				Match{"Juan", "Tasha"},
				Match{"Ismail", "Olivia"},
			}),
			3,
		},
		MatchCeremony{
			NewMatchSet([]Match{
				Match{"Joshy", "Robyn"},
				Match{"Theo", "Libby"},
				Match{"Cach", "Taofiqah"},
				Match{"Jacob", "Vic"},
				Match{"Josh", "Thea"},
				Match{"Jack", "Tersea"},
				Match{"Charlie", "Sapphia"},
				Match{"Jordan", "Shae"},
				Match{"Juan", "Tasha"},
				Match{"Ismail", "Olivia"},
			}),
			3,
		},
		MatchCeremony{
			NewMatchSet([]Match{
				Match{"Jordan", "Robyn"},
				Match{"Charlie", "Libby"},
				Match{"Juan", "Taofiqah"},
				Match{"Theo", "Vic"},
				Match{"Joshy", "Thea"},
				Match{"Jacob", "Tersea"},
				Match{"Jack", "Sapphia"},
				Match{"Cach", "Shae"},
				Match{"Josh", "Tasha"},
				Match{"Ismail", "Olivia"},
			}),
			2,
		},
		MatchCeremony{
			NewMatchSet([]Match{
				Match{"Jordan", "Taofiqah"},
				Match{"Charlie", "Shae"},
				Match{"Juan", "Tersea"},
				Match{"Theo", "Libby"},
				Match{"Joshy", "Vic"},
				Match{"Jacob", "Thea"},
				Match{"Jack", "Robyn"},
				Match{"Cach", "Sapphia"},
				Match{"Josh", "Tasha"},
				Match{"Ismail", "Olivia"},
			}),
			1,
		},
	}

	possibleMatchSets := []MatchSet{}
	for i:=0; ; i++ {
		matches := nextMatches()
		if i % 100000 == 0 {
			fmt.Println(i)
			fmt.Println(matches)
		}
		if matches == nil {
			break
		}
		allMet := true
		for _, c := range conditions {
			allMet = allMet && c.IsMet(matches)
			if !c.IsMet(matches) {
			}
		}
		if allMet {
			possibleMatchSets = append(possibleMatchSets, matches)
		}
	}

	possibleMatches := NewMatchSet([]Match{})
	for _, ms := range possibleMatchSets {
		for m := range ms {
			possibleMatches.Add(m)
		}
	}

	matchProbability := map[Match]float64{}
	for m := range possibleMatches {
		matchFrequency := 0
		for _, ms := range possibleMatchSets {
			if ms.Has(m) {
				matchFrequency += 1
			}
		}
		matchProbability[m] = float64(matchFrequency)/float64(len(possibleMatchSets))
	}

	// for _, male := range males {
	// 	femaleProbabilities := map[Match]float64
	// 	for _, female := range females {
	// 		match := Match{male, female}
	// 		prob, ok := matchProbability[match]
	// 		if ok {
	// 			femaleProbabilities[match] = prob
	// 		}
	// 	}
	// 	fmt.Println("Probabilities", male, ":")
	// 	sorted = SortMapByFunc(femaleProbabilities, func(_, v float64) {return v})
	// 	for _, el := range sorted {
	// 		var key Match
	// 		for k := range el {
	// 			key = k
	// 		}
	// 		fmt.Println(" ", key.Female, el[key])
	// 	}
	// }

	for _, male := range males {
		fmt.Println("Probabilities", male, ":")
		for _, female := range females {
			prob, ok := matchProbability[Match{male, female}]
			if ok {
				fmt.Println(" ", female, prob)
			}
		}
	}

	for _, female := range females {
		fmt.Println("Probabilities", female, ":")
		for _, male := range males {
			prob, ok := matchProbability[Match{male, female}]
			if ok {
				fmt.Println(" ", male, prob)
			}
		}
	}

	fmt.Println("Number possible match sets:", len(possibleMatchSets))
	fmt.Println("Number Possible matches:", len(matchProbability))

	confirmedMatches := []Match{}
	for m := range possibleMatchSets[0] {
		inAll := true
		for _, ms := range possibleMatchSets {
			if !ms.Has(m) {
				inAll = false
				break
			}
		}
		if inAll {
			confirmedMatches = append(confirmedMatches, m)
		}
	}
	fmt.Printf("Confirmed: %v\n", confirmedMatches)

	// Most likely match set
	maxProb := 0.0
	var maxProbMatchSet MatchSet
	for _, ms := range possibleMatchSets {
		prob := 0.0
		for m := range ms {
			prob += matchProbability[m]
		}
		if prob > maxProb {
			maxProb = prob
			maxProbMatchSet = ms
		}
	}
	fmt.Println("Most likely match set:", maxProbMatchSet)
	for m := range maxProbMatchSet {
		fmt.Println(m, matchProbability[m])
	}
}
