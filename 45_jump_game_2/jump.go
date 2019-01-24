package jump_game_2

/**
45. Jump Game II
leecode link: https://leetcode.com/problems/jump-game-ii/
solution: referece to BFS.
*/

func jump(nums []int) int {
	length := len(nums)
	step, start, end := 0, 0, 0
	for end < length-1 {
		step++
		maxend := end + 1
		for i := start; i <= end; i++ {
			if i+nums[i] >= length-1 {
				return step
			}
			if maxend < i+nums[i] {
				maxend = i + nums[i]
			}
		}
		start = end + 1
		end = maxend
	}
	return step
}
