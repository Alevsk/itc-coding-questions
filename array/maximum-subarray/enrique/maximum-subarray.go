package enrique

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	highestSum := nums[0]
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		v := nums[i]

		if sum > highestSum {
			highestSum = sum
		}

		if v > sum + v {
			sum = v
			continue
		}

		sum += v
	}

	if highestSum > sum {
		sum = highestSum
	}

	return sum
}
