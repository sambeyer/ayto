package main

import "fmt"

func PrintStats(PossilbeMatchSets []MatchSet, Males []string, Females []string) {
	possibleMatches := NewMatchSet([]Match{})
	for _, ms := range PossilbeMatchSets {
		for m := range ms {
			possibleMatches.Add(m)
		}
	}

	matchProbability := map[Match]float64{}
	for m := range possibleMatches {
		matchFrequency := 0
		for _, ms := range PossilbeMatchSets {
			if ms.Has(m) {
				matchFrequency += 1
			}
		}
		matchProbability[m] = float64(matchFrequency) / float64(len(PossilbeMatchSets))
	}

	// for _, male := range Males {
	// 	femaleProbabilities := map[Match]float64
	// 	for _, female := range Females {
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

	for _, male := range Males {
		fmt.Println("Probabilities", male, ":")
		for _, female := range Females {
			prob, ok := matchProbability[Match{male, female}]
			if ok {
				fmt.Println(" ", female, prob)
			}
		}
	}

	for _, female := range Females {
		fmt.Println("Probabilities", female, ":")
		for _, male := range Males {
			prob, ok := matchProbability[Match{male, female}]
			if ok {
				fmt.Println(" ", male, prob)
			}
		}
	}

	fmt.Println("Number possible match sets:", len(PossilbeMatchSets))
	fmt.Println("Number Possible matches:", len(matchProbability))

	confirmedMatches := []Match{}
	for m := range PossilbeMatchSets[0] {
		inAll := true
		for _, ms := range PossilbeMatchSets {
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
	for _, ms := range PossilbeMatchSets {
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
