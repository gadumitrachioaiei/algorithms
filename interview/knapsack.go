package interview

func Knapsack(s, v []int, c int) int {
	t := make([][]int, len(s)+1)
	for i := 0; i < len(t); i++ {
		t[i] = make([]int, c+1)
	}
	for i := 0; i < len(s); i++ {
		for j := 0; j <= c; j++ {
			t[i+1][j] = t[i][j]
			if j-s[i] >= 0 && t[i][j-s[i]]+v[i] > t[i+1][j] {
				t[i+1][j] = t[i][j-s[i]] + v[i]
			}
		}
	}
	maxValue := 0
	for i := 1; i <= len(s); i++ {
		if t[i][c] > maxValue {
			maxValue = t[i][c]
		}
	}
	return maxValue
}
