#!/usr/bin/env python
# -*- coding: utf-8 -*-

class Solution:
    def combinationSum4(self, nums: List[int], target: int) -> int:
        
        dp = [0] * (target + 1)
        dp[0] = 1
        
        for i in range(target + 1):
            j = 0
            while  j < len(nums):
                # not in order
                if  i - nums[j] >= 0:
                    dp[i] += dp[i-nums[j]] 
                
                j+=1
                
        return dp[target]
