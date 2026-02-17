func isPalindrome(s string) bool {
    var scl []rune

    for _, ch:= range strings.ToLower(s){
        if unicode.IsLetter(ch) || unicode.IsDigit(ch){
            scl = append(scl, ch)
        }
    }
    left, right:=0, len(scl)-1

    for left < right{
        if scl[left]!=scl[right]{
            return false
        }
        left ++
        right--
    }
    return true
}