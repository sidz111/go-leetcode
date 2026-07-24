func firstPalindrome(words []string) string {
    for _, s := range words {
		j := 0

		for ; j < len(s)/2; j++ {
			if s[j] != s[len(s)-1-j] {
				break
			}
		}

		if j == len(s)/2 {
			return s
		}
	}

	return ""
}