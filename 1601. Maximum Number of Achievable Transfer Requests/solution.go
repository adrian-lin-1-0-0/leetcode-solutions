package solution

func maximumRequests(n int, requests [][]int) int {
    ans :=0
    builds := make([]int,n)
    k := len(requests)

    var dfs func(int,int)
    dfs = func(idx,cnt int){
        if idx == k{
            for i:=0;i<n;i++{
                if builds[i]!=0{
                    return
                }
            }

            ans = max(ans,cnt)
            return
        }

        builds[requests[idx][0]] --
        builds[requests[idx][1]] ++

        dfs(idx+1,cnt+1) 

        builds[requests[idx][0]]++
        builds[requests[idx][1]]--

        dfs(idx+1,cnt)
    }

    dfs(0,0)
    return ans
}