# [992. Subarrays with K Different Integers](https://leetcode.com/problems/subarrays-with-k-different-integers/description/)

The question requires us to find subarrays
within the given element set with a total number equal to K
and calculate the total number of these subarrays.

We can use the Sliding Window to calculate two things:
- First, the total number of subarrays with a total number in range from 1 to K
- Second, the total number of subarrays with a total number in range from 1 to K-1

And the answer is the difference between the two.

The time complexity is O(n) and the space complexity is O(n).