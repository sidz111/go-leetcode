func distributeCandies(candyType []int) int {
    limit := len(candyType) / 2
    unique := make(map[int]struct{})

    for _, c := range candyType {
        unique[c] = struct{}{}

        if len(unique) == limit {
            return limit
        }
    }
    return len(unique)
}