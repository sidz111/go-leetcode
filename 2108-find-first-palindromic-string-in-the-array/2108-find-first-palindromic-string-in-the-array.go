func firstPalindrome(words []string) string {
    for _, s := range words {
		isPalindrome := true

		for j := 0; j < len(s)/2; j++ {
			if s[j] != s[len(s)-1-j] {
				isPalindrome = false
				break
			}
		}

		if isPalindrome {
			return s
		}
	}

	return ""
}