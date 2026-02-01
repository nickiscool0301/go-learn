package main

func sortColors1(nums []int) {
	zeroes, ones, twos := 0, 0, 0
	for _, num := range nums {
		if num == 0 {
			zeroes++
		} else if num == 1 {
			ones++
		} else if num == 2 {
			twos++
		}
	}

	for i := 0; i < len(nums); i++ {
		if i < zeroes {
			nums[i] = 0
		} else if i < zeroes+ones {
			nums[i] = 1
		} else {
			nums[i] = 2
		}
	}
}

func sortColors2(nums []int) {
	low, mid, high := 0, 0, len(nums)-1
	for mid <= high {
		switch nums[mid] {
		case 0:
			nums[low], nums[mid] = nums[mid], nums[low]
			mid++
			low++
		case 1:
			mid++
		case 2:
			nums[mid], nums[high] = nums[high], nums[mid]
			mid++
			high--
		}
	}
}
