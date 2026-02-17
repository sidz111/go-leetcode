func findTheDifference(s string, t string) byte {
    var result byte

    for i := 0; i < len(s); i++ {
        result ^= s[i]
    }

    for i := 0; i < len(t); i++ {
        result ^= t[i]
    }

    return result
}