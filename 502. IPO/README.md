# [502. IPO](https://leetcode.com/problems/ipo/description/)

## Overview
The goal of this function is to maximize the capital after completing at most k projects,
starting with an initial capital w.

Each project has a required capital to start and a profit associated with it.


## Approach

1. Create a list of project indices to represent each project.
Sort the projects based on the required capital in ascending order.

2. Initialize a max heap to keep track of the most profitable projects
that can be started with the current available capital.

3. Iterate up to k times to select projects.

    For each iteration, add all projects
    that can be started with the current capital to the max heap.

    If the max heap is empty,
    break out of the loop because no more projects can be started.

    Pop the most profitable project from the heap
    and add its profit to the current capital.
