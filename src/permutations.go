package main

// https://stackoverflow.com/questions/30226438/generate-all-permutations-in-go

func nextPerm(p []int) {
    for i := len(p) - 1; i >= 0; i-- {
        if i == 0 || p[i] < len(p)-i-1 {
            p[i]++
            return
        }
        p[i] = 0
    }
}

func getPerm[T any] (orig []T, p []int) []T {
    result := make([]T, len(orig))
    copy(result, orig)
    for i, v := range p {
        result[i], result[i+v] = result[i+v], result[i]
    }
    return result
}

func Permutations[T any](orig []T) func() []T {
    first := true
    p := make([]int, len(orig))
    return func() []T {
        if !first {
            nextPerm(p)
        } else {
            first = false
        }
        if p[0] < len(p) {
            return getPerm(orig, p)
        }
        return nil
    }
}
