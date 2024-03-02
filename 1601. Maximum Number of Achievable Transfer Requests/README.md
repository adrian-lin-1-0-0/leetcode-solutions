This is a common knapsack problem. 

For each request, <br>
we can decide whether to use it or not, <br>
then explore all possibilities <br>
and find the maximum result.

The time complexity of this solution is O(2^m*n), Big O of 2 to the power of m times n, <br>
where m is the length of the requests,<br>
and n is the input number n.<br>

```go
func maximumRequests(n int, requests [][]int) int {
    ...
}
```

We declare a DFS (Depth-First Search) function to explore all possibilities. <br>

The function takes two parameters: 
- the current position.
- the number of accepted requests.

```go
dfs = func(idx,cnt int){
    ...
}
```

We also declare an array called buildings to record the number of people in the building after each request change. 

```go
buildings := make([]int,n)
```

When the number in buildings are equal to 0,
it means that the current combination of requests is valid,
and we can update the answer.

```go
for i:=0;i<n;i++{
    if buildings[i]!=0{
        return
    }
}
```


If we accept this request, we update the value of buildings:

buildings's value at the N-th request's first element will increase by 1
buildings's value at the N-th request's second element will decrease by 1


```go
buildings[requests[idx][0]] --
buildings[requests[idx][1]] ++
```

After the first DFS function ends, we will restore the value of buildings:

buildings's value at the N-th request's first element will decrease by 1
buildings's value at the N-th request's second element will increase by 1

```go
buildings[requests[idx][0]]++
buildings[requests[idx][1]]--
```


In this way, we can explore all possibilities and find the maximum result.

```go
func maximumRequests(n int, requests [][]int) int {
    ans :=0
    buildings := make([]int,n)
    k := len(requests)

    var dfs func(int,int)
    dfs = func(idx,cnt int){
        if idx == k{
            for i:=0;i<n;i++{
                if buildings[i]!=0{
                    return
                }
            }

            ans = max(ans,cnt)
            return
        }

        buildings[requests[idx][0]] --
        buildings[requests[idx][1]] ++

        dfs(idx+1,cnt+1) 

        buildings[requests[idx][0]]++
        buildings[requests[idx][1]]--

        dfs(idx+1,cnt)
    }

    dfs(0,0)
    return ans
}
```