package solution

func MoveZeroes(nums []int) []int {
	left := 0

	for right := 0; right < len(nums); right++ {
		if nums[right] != 0 {
			temp := nums[right]
			nums[right] = nums[left]
			nums[left] = temp

			left++
		}
	}
	return nums

}
