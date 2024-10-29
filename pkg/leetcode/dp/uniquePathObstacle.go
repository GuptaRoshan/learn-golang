package dp

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {

	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	dp := make([][]int, m)

	for i := range dp {
		dp[i] = make([]int, n)
	}

	// Search in the first column
	for i := 0; i < m; i++ {

		if obstacleGrid[i][0] == 1 {
			dp[i][0] = 0
			break
		} else {
			dp[i][0] = 1
		}

	}

	// Search in tye first row
	for i := 0; i < n; i++ {

		if obstacleGrid[0][i] == 1 {
			dp[0][1] = 0
			break
		} else {
			dp[0][i] = 1
		}

	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {

			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i][j-1] + dp[i-1][j]
			}

		}
	}

	return dp[m-1][n-1]

}

func TestUniquePathsObstacle() {
	obstacleGrid := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}

	fmt.Print(uniquePathsWithObstacles(obstacleGrid))
}
