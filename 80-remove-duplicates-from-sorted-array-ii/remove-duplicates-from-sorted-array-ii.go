func removeDuplicates(nums []int) int {
    n := 0
    hm := make(map[int]int)

    for _, i := range nums {
        if hm[i] < 2 {
            hm[i]++
            nums[n] = i
            n++
        }
    }

    return n
}