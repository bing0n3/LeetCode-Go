// Package dp provides ...
package dp

type NumArray struct {
	memo []int
}

func Constructor(nums []int) NumArray {
	memo := make([]int, len(nums))
	total := 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
		memo[i] = total
	}

	return NumArray{memo: memo}
}

func (this *NumArray) SumRange(i int, j int) int {
	if i == 0 {
		return this.memo[j]
	}
	return this.memo[j] - this.memo[i-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
