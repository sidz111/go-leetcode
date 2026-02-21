func distributeCandies(candyType []int) int {
    uniqueCandies := make(map[int]bool)
    for _, candy := range candyType {
        uniqueCandies[candy] = true
    }

    maxTypes := len(uniqueCandies)
    allowedToEat := len(candyType) / 2

    return int(math.Min(float64(maxTypes), float64(allowedToEat)))
}