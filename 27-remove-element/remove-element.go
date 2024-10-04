func removeElement(nums []int, val int) int {
    n := 0

    for _, i := range nums {
        if i != val {
            nums[n] = i
            n++
        }
    }
    return n
}