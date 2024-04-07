package solution

import (
	"math"
	"sort"
)

func minimumDistance(points [][]int) int {

	n := len(points)

	c1 := make([][]int, 0, n)
	c2 := make([][]int, 0, n)
	for i, p := range points {
		c1 = append(c1, []int{p[0] + p[1], i})
		c2 = append(c2, []int{p[0] - p[1], i})
	}

	sort.Slice(c1, func(i, j int) bool {
		return c1[i][0] < c1[j][0]
	})

	sort.Slice(c2, func(i, j int) bool {
		return c2[i][0] < c2[j][0]
	})

	ans := math.MaxInt32
	l, r := 0, n-1

	// fmt.Println(c1)
	// fmt.Println(c2)

	for i := 0; i < n; i++ {

		x, y := c1[l][0], c1[r][0]

		if c1[l][1] == i {
			x = c1[l+1][0]
		}
		if c1[r][1] == i {
			y = c1[r-1][0]
		}
		c1a := diff(x, y)

		x, y = c2[l][0], c2[r][0]

		if c2[l][1] == i {
			x = c2[l+1][0]
		}
		if c2[r][1] == i {
			y = c2[r-1][0]
		}

		c2a := diff(x, y)

		ca := max(c1a, c2a)
		ans = min(ans, ca)
	}
	return ans
}

func diff(a, b int) int {
	if a-b > 0 {
		return a - b
	}

	return b - a
}
