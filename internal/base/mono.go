package base

func Mono(nums []int) bool {
	mono_decreasing := true
	mono_increasing := true

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			mono_decreasing = false
		}
		if nums[i] < nums[i-1] {
			mono_increasing = false
		}
	}
	return mono_decreasing || mono_increasing
}
