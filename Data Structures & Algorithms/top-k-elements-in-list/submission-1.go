func topKFrequent(nums []int, k int) []int {
    freq := make(map[int]int, len(nums))
    for _, num := range nums {
        freq[num]++
    }

    buckets := make([][]int, len(nums)+1)
    for num, f := range freq {
        buckets[f] = append(buckets[f], num)
    }

    res := make([]int, 0, k)
    for i := len(buckets) - 1; i >= 0 && len(res) < k; i-- {
        res = append(res, buckets[i]...)
    }
    return res
}