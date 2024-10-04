func merge(nums1 []int, m int, nums2 []int, n int)  {
    n1 := nums1[:m]
    n2 := nums2[:n]

    n1 = append(n1, n2...)
    slices.Sort(n1)
    nums1 = n1
}