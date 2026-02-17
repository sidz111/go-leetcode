func lengthOfLastWord(s string) int {
    newStr := strings.TrimRight(s, " ")
    count := 0
    for i:=len(newStr)-1; i>=0; i--{
        if newStr[i]==' '{
            break
        }
        count ++
    }
    return count
}