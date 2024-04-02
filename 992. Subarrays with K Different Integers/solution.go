package solution

// 1~k - 1~k-1 == k
func help(nums []int, k int) int {
	n := len(nums)
	sets := make(map[int]int)
	left, right, ans := 0, 0, 0

	for ; right < n; right++ {

		sets[nums[right]]++

		for len(sets) > k {

			sets[nums[left]]--

			if sets[nums[left]] == 0 {
				delete(sets, nums[left])
			}
			left++
		}

		ans += right - left + 1
	}
	return ans
}

func subarraysWithKDistinct(nums []int, k int) int {
	return help(nums, k) - help(nums, k-1)
}
