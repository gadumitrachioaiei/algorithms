package interview

func LongestSubsequence(s string) string {
	table := make([]string, len(s)) // list with the longest subsequence ending in s[i]
	table[0] = s[0:1]
	var result string
	for i := 1; i < len(s); i++ {
		var sequence string
		for j := 0; j < i; j++ {
			if s[j] <= s[i] {
				if len(table[j]) > len(sequence) {
					sequence = table[j]
				}
			}
		}
		table[i] = sequence + s[i:i+1]
		if len(result) < len(table[i]) {
			result = table[i]
		}
	}
	return result
}

func LongestSubsequence2(s string) string {
	table := make([]int, len(s))       // list with the length of the longest subsequence ending in s[i]
	predecessor := make([]int, len(s)) // maps the above subsequence with the index of the previous subsequence from which this was constructed
	table[0] = 1
	for i := 0; i < len(predecessor); i++ {
		predecessor[i] = -1
	}
	var resultIndex int // which subsequence is the longest
	for i := 1; i < len(s); i++ {
		var max int
		for j := 0; j < i; j++ {
			if s[j] <= s[i] {
				if max < table[j] {
					max = table[j]
					predecessor[i] = j
				}
			}
		}
		table[i] = max + 1
		if table[resultIndex] < table[i] {
			resultIndex = i
		}
	}
	// trace back the longest subsequence, using the predecessor
	result := make([]byte, table[resultIndex])
	i := resultIndex
	index := len(result) - 1
	for i > -1 {
		result[index] = s[i]
		index--
		i = predecessor[i]
	}
	return string(result)
}
