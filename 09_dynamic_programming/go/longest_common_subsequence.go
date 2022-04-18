package main

import "fmt"

func main() {
	fmt.Println(lcs("fish", "hish"))
	fmt.Println(lcs("vista", "hish"))
}

func lcs(a, b string) int {
	m := len(a)
	n := len(b)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

func max(m int, n int) int {
	if m >= n {
		return m
	}
	return n
}
