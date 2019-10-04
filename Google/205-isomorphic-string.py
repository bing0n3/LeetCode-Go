from collections import defaultdict


class Solution:
    def isIsomorphic(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False

        seen_s = defaultdict(int)
        seen_t = defaultdict(int)

        for i in range(len(t)):
            if seen_s.get(s[i], -1) == seen_t.get(t[i], -1):
                seen_s[s[i]] = i
                seen_t[t[i]] = i
            else:
                return False

        return True
