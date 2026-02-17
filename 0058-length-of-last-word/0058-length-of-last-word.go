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


    end := len(s) - 1

    // skip trailing spaces
    for end >= 0 && s[end] == ' ' {
        end--
    }

    // find last space
    start := end
    for start >= 0 && s[start] != ' ' {
        start--
    }

    return end - start
}