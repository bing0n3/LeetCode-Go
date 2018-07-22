package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for l := i + 1; l < len(nums); l++ {
			if nums[l]+nums[i] == target {
				return []int{nums[i], nums[l]}
			}
		}
	}
	return []int{}

}

func twoSum2(nums []int, target int) []int {
	buff := make(map[int]int)
	for i, num := range nums {
		val, ok := buff[target-num]
		if ok {
			return []int{val, i}
		}
		buff[nums[i]] = i
	}
	return nil
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 8}
	for _, i := range twoSum(numbers, 9) {
		fmt.Println(i)
	}
}
