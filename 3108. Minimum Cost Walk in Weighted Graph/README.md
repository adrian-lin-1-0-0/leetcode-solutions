## 3108. Minimum Cost Walk in Weighted Graph

The problem entails traversing back and forth between different nodes in a weighted graph. 

And we can pass through a node multiple times.

Each time we pass through an edge, we can take the bitwise AND operation on the cost. 
Thus, we initially compute the bitwise AND of the costs of all connected nodes. 

Utilize the Union-find data structure to merge all connected nodes 
and simultaneously calculate the total cost. 

Finally, we only need to verify whether two nodes are connected. 
If they are, return the cost; otherwise, return -1.