class Solution(object):
    def oddEvenJumps(self, A):
        """
        :type A: List[int]
        :rtype: int
        """
    
        N = len(A)


        
        # mono decrease stack
        odd_next, even_next = [None] * N, [None] * N
        
        
        
        # order index by value ascending
        B = sorted(range(N), key = lambda i: A[i])
        
        stack = []
        
        for i in B:
            while stack and stack[-1] < i:
                odd_next[stack.pop()] = i
            stack.append(i)
            
        stack = []
        for i in sorted(range(N), key = lambda i: -A[i]):
            while stack and stack[-1] < i:
                even_next[stack.pop()] = i
            stack.append(i)
            
        # dp
        
        odd = [False] * N
        even = [False] * N
        odd[N-1] = even[N-1] = True
        
        for i in range(N-2, -1, -1):
            if odd_next[i] is not None:
                odd[i] = even[odd_next[i]]
            if even_next[i] is not None:
                even[i] = odd[even_next[i]]
                
        return sum(odd)
