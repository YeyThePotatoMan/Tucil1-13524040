package main

import (
	"fmt"
)

func main() {
	fmt.Scan(&n)
	g = make([]string, n)
	ans = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&g[i])
	}

	mask := make(map[uint8]bool)
	solve2(0, false, mask)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j == ans[i] {
				fmt.Print("#")
			} else {
				fmt.Print(string(g[i][j]))
			}
		}
		fmt.Println()
	}
}
