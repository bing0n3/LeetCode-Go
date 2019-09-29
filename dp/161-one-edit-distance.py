class Solution:
    def isOneEditDistance(self, s: str, t: str) -> bool:
        m = len(s) 
        n = len(t)
        
        if abs(n-m) > 1:
            return False
        
        
        for i in range(min(m,n)):
            if s[i] != t[i]:
                if m == n:
                    return s[i+1:] == t[i+1:]
                if m > n:
                    return s[i+1:] == t[i:]
                if m < n: 
                    return s[i:] == t[i+1:]
                
        return abs(n-m) == 1
