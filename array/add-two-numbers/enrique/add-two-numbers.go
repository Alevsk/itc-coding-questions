package enrique

func twoSum(nums []int, target int) []int {
	var indexA, indexB int

	valueToIdx := map[int]int{}
	for i, v := range nums {
		difference := target - v

		if idx, ok := valueToIdx[difference]; ok {
			indexA = i
			indexB = idx

			break
		}

		valueToIdx[v] = i
	}

	return []int{indexA, indexB}
}
