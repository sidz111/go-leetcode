func twoSum(nums []int, target int) []int {
    m := make(map[int]int, 5)
    for i, v:= range nums{
        if val, ok:= m[target-v]; ok{
            return []int{val, i}
        } 
        m[v]=i
    }
    return []int{}
}