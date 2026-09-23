func hasDuplicate(nums []int) bool {
    unique := make(map[int]struct{})
    for _, num := range nums {
        if _, exists := unique[num]; exists {
            return true
        }
        unique[num] = struct{}{}
    }
    return false
}