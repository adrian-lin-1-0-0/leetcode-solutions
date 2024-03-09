# [1626. Best Team With No Conflicts](https://leetcode.com/problems/best-team-with-no-conflicts/description/)

This is a common dynamic programming problem.

Our every choice is influenced by the previous choices, and we can start by sorting ages and scores. This way, we only need to consider the previous selection.

```python
n = len(scores)

players = [(ages[i], scores[i]) for i in range(n)]
players.sort(key=lambda x: (x[0], x[1]))
```

First, we use depth-first search (DFS) to traverse all possibilities and find the maximum result. In this process, we need two parameters:

Current position
Score of the previous selection

```python
def dfs(i, prev):
    if i == n:
        return 0

    if players[i][1] >= prev:
        ans1 = dfs(i + 1, prev)
        ans2 = dfs(i + 1, players[i][1]) + players[i][1]

        return max(ans1,ans2)

    return dfs(i + 1, prev)

```

Then, we use memoization to optimize the time complexity.

```python

@cache 
def dfs(i, prev):
    ...
```


```python
class Solution:
    def bestTeamScore(self, scores: List[int], ages: List[int]) -> int:
        n = len(scores)
        
        players = [(ages[i], scores[i]) for i in range(n)]
        players.sort(key=lambda x: (x[0], x[1]))

        @cache
        def dfs(i, prev):
            if i == n:
                return 0

            if players[i][1] >= prev:
                ans1 = dfs(i + 1, prev)
                ans2 = dfs(i + 1, players[i][1]) + players[i][1]

                return max(ans1,ans2)

            return dfs(i + 1, prev)

        return dfs(0, 0)

```


```python
class Solution:
    def bestTeamScore(self, scores: List[int], ages: List[int]) -> int:
        
        
        arr = sorted(zip(scores, ages), key = lambda x: (x[1], x[0]))
        
        n = len(arr)
        
        @cache
        def dfs(idx):
            
            if idx == n:
                return 0
            
            ans = arr[idx][0]
            
            for i in range(idx + 1, n):
                
                if arr[idx][0] <= arr[i][0]:
                    ans = max(ans, arr[idx][0] + dfs(i))
                    
            return ans
        
        return max([dfs(i) for i in range(n)])
```