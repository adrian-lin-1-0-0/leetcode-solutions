# [3102. Minimize Manhattan Distances](https://leetcode.com/problems/minimize-manhattan-distances/description/)


You are given a array points representing integer coordinates of some points on a 2D plane, where points[i] = [xi, yi].

The distance between two points is defined as their 
Manhattan distance
.

Return the minimum possible value for maximum distance between any two points by removing exactly one point.

題目的意思是:
我們可以依序拿掉每個點,計算所有點之間的最大曼哈頓距離
存在一個點,使得移除該點後,剩下的點之間的最大曼哈頓距離最小

求移除該點後的最大曼哈頓距離

解法1:

週賽當下是用比較直覺的方法,對於每個點,我們都拿掉該點,然後計算剩下的點之間的最大曼哈頓距離

如果每個點都要計算一次,那麼時間複雜度就是O(n^2)

所以我們可以先推倒如何計算最大曼哈頓距離

|Xi – Xj| + |Yi – Yj| 

= max(
     Xi - Xj - Yi + Yj,
    -Xi + Xj + Yi - Yj,
    -Xi + Xj - Yi + Yj,
     Xi - Xj + Yi - Yj
)

= max(
    ( Xi - Yi) - ( Xj - Yj),
    (-Xi + Yi) - (-Xj + Yj),
    (-Xi - Yi) - (-Xj - Yj),
    ( Xi + Yi) - ( Xj + Yj)
)

