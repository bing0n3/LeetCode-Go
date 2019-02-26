package permutationsequence

import (
	"sort"
)

func getPermutation(n int, k int) string {

	numbers := []int{}
	for i := 1; i <= n; i++ {
		numbers = append(numbers, i)
	}

	factorial := []int{1}
	for i := 1; i <= n; i++ {
		factorial = append(factorial, factorial[i-1]*(i+1))
	}

	res := []byte{}

	k--

	for i := n - 1; i >= 1; i-- {
		index := k / factorial[i-1]
		res = append(res, byte(numbers[index])+'0')
		numbers = append(numbers[:index], numbers[index+1:]...)
		k = k % factorial[i-1]

	}

	res = append(res, byte(numbers[0])+'0')

	return string(res)
}

// getPermutation using backtracking
func getPermutation_backtracking(n int, k int) string {
	result := []string{}
	nums := []byte{}
	for i := 1; i <= n; i++ {
		nums = append(nums, byte(i)+'0')
	}

	backtracking(nums, &result, 0)

	sort.Strings(result)

	return result[k-1]
}

func backtracking(nums []byte, result *[]string, index int) {
	if index == len(nums) {
		*result = append(*result, string(nums))
		return
	}
	for i := index; i < len(nums); i++ {
		nums[i], nums[index] = nums[index], nums[i]
		backtracking(nums, result, index+1)
		nums[i], nums[index] = nums[index], nums[i]
	}
}
