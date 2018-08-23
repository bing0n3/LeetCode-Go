/**
1. Two Sum
https://leetcode.com/problems/two-sum/description/
**/
package two_sum

import "fmt"

func twoSum(nums []int, target int) []int {
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
