# [3102. Minimize Manhattan Distances](https://leetcode.com/problems/minimize-manhattan-distances/description/)


You are given a array points representing integer coordinates of some points on a 2D plane, where points[i] = [xi, yi].

The distance between two points is defined as their 
Manhattan distance
.

Return the minimum possible value for maximum distance between any two points by removing exactly one point.


### The meaning of the question:
We can sequentially remove each point and calculate the maximum Manhattan distance between all points.
There exists a point such that after removing it, the maximum Manhattan distance between the remaining points is minimized.

### Solution1:

At the time of the weekly contest, I used a more intuitive method. For each point, we remove that point and then calculate the maximum Manhattan distance between the remaining points.

If we calculate each point once, the time complexity is O(n^2).

So we can first figure out how to calculate the maximum Manhattan distance.

```
|Xi – Xj| + |Yi – Yj| 

= max(
     Xi - Xj - Yi + Yj,
    -Xi + Xj + Yi - Yj,
    -Xi + Xj - Yi + Yj,
     Xi - Xj + Yi - Yj
)

max(
    |(Xi + Yi) - (Xj + Yj)|,
    |(Xi + Yi) - (Xj - Yj)|
)

```

So we can first calculate the values of (Xi + Yi) and (Xi - Yi) for all points,
then sort, and then we can calculate the maximum Manhattan distance in O(1) time.

Sorting: O(nlogn)