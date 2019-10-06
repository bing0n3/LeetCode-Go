#!/usr/bin/env python
# -*- coding: utf-8 -*-

class Solution:
    def numSquares(self, n: int) -> int:
        if n <= 0:
            return 0
        
        lst = []
        
        i = 1
        
        while i * i <= n:
            lst.append(i*i)
            i += 1
            
        queue = {n}
        cnt = 0
        while queue:
            cnt+= 1
            temp = set()
            for x in queue:
                for y in lst:
                    if  x == y:
                        return cnt
                    if x < y:
                        break 
                    temp.add(x-y)
            queue = temp
            
        
        
        return cnt
            
            
            

