package index

// 双指针从后往前
func merge(nums1 []int, m int, nums2 []int, n int) {
	maxLength := len(nums1)
	k := maxLength - 1
	i := m - 1
	j := n - 1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
	for j>=0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}
