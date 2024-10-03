func majorityElement(nums []int) int {
    n := len(nums) / 2
    f := make(map[int]int)
    j := 0

    for _, i := range nums {
        f[i]++
        if f[i] > n {
            j = i
        }
    }
    return j
}