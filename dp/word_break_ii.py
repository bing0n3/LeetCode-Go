class Solution:
    def wordBreak(self, s: str, wordDict: List[str]) -> List[str]:
        n = len(s)

        # dfs backtracking

        seen = dict()

        return self.dfs(s, wordDict, seen)

    def dfs(self, s, wordDict, seen):
        if s in seen:
            return seen[s]

        res = []

        if s == "":
            return [""]

        for word in wordDict:
            if s.startswith(word):
                subList = self.dfs(s[len(word) :], wordDict, seen)
                for sub in subList:
                    if sub == "":
                        res.append(word)
                    else:
                        res.append(word + " " + sub)
        seen[s] = res
        return res
