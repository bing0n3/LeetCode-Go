package jump_game

/*
55 Jump Game

Description https://leetcode.com/problems/jump-game/
only have maxmium n chance to jump for n length array. simialr to game jump 2, reference to bfs
*/
func canJump(nums []int) bool {
	length := len(nums)
	start, end := 0, 0
	step := length

	for step > 0 {
		maxend := end
		for i := start; i <= end; i++ {
			if i+nums[i] >= length-1 {
				return true
			}
			if maxend < i+nums[i] {
				maxend = i + nums[i]
			}
		}
		start = end + 1
		end = maxend
		step--
	}
	return false
}
