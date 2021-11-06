package enrique

func productExceptSelf(nums []int) []int {
	right := getInitializedProductArray(len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		if i+1 > len(nums)-1 {
			continue
		}

		right[i] *= nums[i+1] * right[i+1]
	}

	left := getInitializedProductArray(len(nums))
	for i := 0; i < len(nums); i++ {
		if i-1 < 0 {
			continue
		}

		left[i] *= nums[i-1] * left[i-1]
	}

	productExceptSelf := getInitializedProductArray(len(nums))
	for i, _ := range nums {
		productExceptSelf[i] = right[i] * left[i]
	}

	return productExceptSelf
}

func getInitializedProductArray(size int) []int {
	productArray := make([]int, size)

	for i, _ := range productArray {
		productArray[i] = 1
	}

	return productArray
}
