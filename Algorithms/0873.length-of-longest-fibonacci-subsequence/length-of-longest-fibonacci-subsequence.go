package problem0873

import "sort"

func lenLongestFibSubseq(a []int) int {
	size := len(a)
	res := 0

	for k := 2; k < size; k++ {
		l, r := 0, k-1
		for l < r {
			s := a[l] + a[r]

			if s < a[k] {
				l++
				continue
			}

			if a[k] < s {
				r--
				continue
			}

			// 此时找到了第一组 a[l] + a[r] = a[k]
			count := 3
			i, j := r, k
			// 反复利用二分查找法，查看下个数是否在 a 中
			for {
				next := a[i] + a[j]
				i, j = j, j+sort.SearchInts(a[j:], next)
				if j == size || a[j] != next {
					break
				}
				count++
			}

			res = max(res, count)
			l++
			r--
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
