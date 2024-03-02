This is a common knapsack problem. 

For each request, 
we can decide whether to use it or not, 
then explore all possibilities 
and find the maximum result.

The time complexity of this solution is Big O of 2 to the power of m times n, 
where m is the length of the requests,
and n is the input number n.

We declare a DFS (Depth-First Search) function to explore all possibilities. 
The function takes two parameters: 
the current position and the number of accepted requests.

We also declare an array called builds to record the number of people in the build after each request change. 

When the number in builds is equal to 0,
it means that the current combination of requests is valid,
and we can update the answer.

If we accept this request, we update the value of builds:

builds's value at the N-th request's first element will increase by 1
builds's value at the N-th request's second element will decrease by 1

After the first DFS function ends, we will restore the value of builds:

builds's value at the N-th request's first element will decrease by 1
builds's value at the N-th request's second element will increase by 1

In this way, we can explore all possibilities and find the maximum result.