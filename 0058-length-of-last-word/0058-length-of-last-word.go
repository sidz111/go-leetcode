func lengthOfLastWord(s string) int {
    // newStr := strings.TrimRight(s, " ")
    // count := 0
    // for i:=len(newStr)-1; i>=0; i--{
    //     if newStr[i]==' '{
    //         break
    //     }
    //     count ++
    // }
    // return count
     i := len(s) - 1

    // 1. skip trailing spaces
    for i >= 0 && s[i] == ' ' {
        i--
    }

    // 2. count last word length
    length := 0
    for i >= 0 && s[i] != ' ' {
        length++
        i--
    }

    return length
}