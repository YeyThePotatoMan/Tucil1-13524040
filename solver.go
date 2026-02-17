package main

func bad(r, c int) bool {
	for i := 0; i < r; i++ {
		pc := ans[i]
		if pc == c {
			return true
		}

		if i == r-1 {
			if pc == c-1 || pc == c+1 {
				return true
			}
		}
	}
	return false
}

// no prune
func solve1(r int, live bool) {
	if found || stop {
		return
	}
	cnt++

	if r == n {
		mask := make(map[uint8]bool)
		ok := true
		for i := 0; i < n; i++ {
			c := g[i][ans[i]]
			if mask[c] {
				ok = false
				break
			}
			mask[c] = true
		}
		if ok {
			found = true
		}
		return
	}

	for c := 0; c < n; c++ {
		if !bad(r, c) {
			ans[r] = c
			solve1(r+1, live)
			if found {
				return
			}
			ans[r] = -1
		}
	}
}

// with pruning
func solve2(r int, live bool, used map[uint8]bool) {
	if found || stop {
		return
	}
	cnt++

	if r == n {
		found = true
		return
	}

	for c := 0; c < n; c++ {
		colour := g[r][c]
		if !bad(r, c) && !used[colour] {
			ans[r] = c
			used[colour] = true

			solve2(r+1, live, used)
			if found {
				return
			}

			used[colour] = false
			ans[r] = -1
		}
	}
}
