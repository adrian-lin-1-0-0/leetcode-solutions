package solution

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