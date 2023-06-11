package _two_sum

// @submit start
func twoSum(nums []int, target int) []int {
	idx := map[int]int{}
	for i, v := range nums {
		if j, ok := idx[target-v]; ok {
			return []int{j, i}
		}
		idx[v] = i
	}
	return nil
}

// @submit end
