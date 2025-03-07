func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    nums1 = append(nums1, nums2...)

    slices.Sort(nums1)

    med := math.Round(float64((len(nums1)-1))/2) // rounds up
    n1 := nums1[int(med)] // get the middle number

    if len(nums1) % 2 == 0 {
        n2 := nums1[int(med-1)] // rounds down
        return ((float64(n1) + float64(n2)) / float64(2))
    }
    return float64(n1)
}