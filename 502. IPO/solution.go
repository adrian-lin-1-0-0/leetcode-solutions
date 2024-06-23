package solution

import (
	"container/heap"
	"sort"
)

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/*
The goal of this function is to maximize the capital after completing at most k projects,
starting with an initial capital w.

Each project has a required capital to start and a profit associated with it.


Overview of the Approach:

1.
Create a list of project indices to represent each project.
Sort the projects based on the required capital in ascending order.

2.
Initialize a max heap to keep track of the most profitable projects
that can be started with the current available capital.

3.
Iterate up to k times to select projects.

For each iteration, add all projects
that can be started with the current capital to the max heap.

If the max heap is empty,
break out of the loop because no more projects can be started.

Pop the most profitable project from the heap
and add its profit to the current capital.

*/

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	n := len(profits)
	projects := make([]int, 0, n)

	for i := 0; i < n; i++ {
		projects = append(projects, i)
	}

	sort.Slice(projects, func(i, j int) bool {
		return capital[projects[i]] < capital[projects[j]]
	})

	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	i := 0

	for j := 0; j < k; j++ {
		// Add all projects that can be started with the current capital to the max heap.
		for i < n && capital[projects[i]] <= w {
			heap.Push(maxHeap, profits[projects[i]])
			i++
		}

		if maxHeap.Len() == 0 {
			break
		}

		w += heap.Pop(maxHeap).(int)
	}

	return w
}
