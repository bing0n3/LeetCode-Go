#!/usr/bin/env python
# -*- coding: utf-8 -*-

def kSum(self, A, k, target):
        # write your code here
        # knapback
        dp = [[[0 for v in range(target+1)] for j in range(k+1)] for i in range(2)]
        dp[0][0][0] = 1
        dp[1][0][0] = 1
        
        for i in range(1, len(A)+1):
            for j in range(1, k+1):
                for v in range(1,target+1):
                    if v >= A[i-1]:
                        dp[i%2][j][v] = dp[(i-1)%2][j][v] + dp[(i-1)%2][j-1][v-A[i-1]]
                    else:
                        dp[i%2][j][v] = dp[(i-1)%2][j][v]
                        
                    
        
        return dp[len(A)%2][k][target]

