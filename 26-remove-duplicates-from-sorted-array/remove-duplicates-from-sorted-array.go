func removeDuplicates(nums []int) int {
    n := 0
    hm := make(map[int]bool)

    for _, i := range nums {
        if !hm[i] {
            hm[i] = true
            nums[n] = i
            n++
        }
    }

    return n
}