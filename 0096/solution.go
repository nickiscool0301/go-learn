package main

/*
 * If we choose i as root:
 	* Left subtree has i-1 nodes
  	* Right subtree has n-i nodes
  	dp[n] = number of unique BSTs with n nodes
   	total for root i = dp[i-1] * dp[n-i]
 *
 *
*/

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for nodes := 2; nodes <= n; nodes++ {
		for root := 1; root <= nodes; root++ {
			dp[nodes] += dp[root-1] * dp[nodes-root]
		}
	}
	return dp[n]
}
