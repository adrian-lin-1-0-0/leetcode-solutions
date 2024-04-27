# [3123. Find Edges in Shortest Paths](https://leetcode.com/problems/find-edges-in-shortest-paths/description/)

Using Dijkstra's algorithm to find the shortest paths from node 0 to all other nodes.

Starting from the node-`n-1` and moving towards the node-`0`,
if there exists an edge whose weight is equal to 
the difference in shortest path distances between the two nodes,
then this edge lies on the shortest path.

For example:
```
distances[x]  = distances[y] + weight(x, y)
```

implies that the edge x-y lies on the shortest path.
