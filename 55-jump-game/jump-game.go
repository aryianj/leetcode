func canJump(nums []int) bool {
    j := nums[0]

    for i := 0; i < len(nums); i++ {
        if i > j {
            return false
        }
        j = max(j, i + nums[i]) 
    }

    return true
}