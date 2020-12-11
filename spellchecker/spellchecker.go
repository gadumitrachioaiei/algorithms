package spellchecker

// Dist returns the smallest editing distance between two ascii strings
func Dist(s1, s2 string) int {
	n, m := len(s1)+1, len(s2)+1
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, m)
		dist[i][0] = i
	}
	for j := 0; j < m; j++ {
		dist[0][j] = j
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			var r int
			if s1[i-1] == s2[j-1] {
				r = dist[i-1][j-1]
			} else {
				r = dist[i-1][j-1] + 1
			}
			dist[i][j] = min(r, dist[i-1][j]+1, dist[i][j-1]+1)
		}
	}
	return dist[n-1][m-1]
}

// TrieNode represents a trie that stores dictionary words, a node per char
// read http://stevehanov.ca/blog/index.php?id=115 for more details
type TrieNode struct {
	word     string // the word that this node represents, if any
	children map[rune]*TrieNode
}

// NewTrieNode creates a new trie
func NewTrieNode() *TrieNode {
	return &TrieNode{children: make(map[rune]*TrieNode)}
}

// InsertWord inserts a new word in the trie, each char being a new node
func (t *TrieNode) InsertWord(w string) {
	node := t
	for _, c := range w {
		if _, ok := node.children[c]; !ok {
			node.children[c] = NewTrieNode()
		}
		node = node.children[c]
	}
	node.word = w
}

// WordMatch represents the matched words together with their editing distance
type WordMatch struct {
	W string
	C int
}

// Search matches word with words in our trie, based on maximum editing distance
func (t *TrieNode) Search(word string, maxDist int) []WordMatch {
	var result []WordMatch
	// dist is the distance between the empty string and word
	dist := make([]int, len([]rune(word))+1)
	for i := 0; i < len(dist); i++ {
		dist[i] = i
	}
	for c, n := range t.children {
		result = append(result, n.search(word, c, dist, maxDist)...)
	}
	return result
}

// search returns a list of words from the trie, that match word, within specified editing distance
// by prefix we understand the list of chars traversed so far, except for the char of current trie node
// currentChar is the added char to the prefix, so the char of the current node
// prefixDist is the distances between the word and the last char in prefix
func (t *TrieNode) search(word string, currentChar rune, prefixDist []int, maxDist int) []WordMatch {
	var result []WordMatch
	// dist is the distance between the current prefix of the trie and word
	dist := distFromPrefixDist(word, currentChar, prefixDist)
	if len(t.word) > 0 && dist[len(dist)-1] <= maxDist {
		result = append(result, WordMatch{W: t.word, C: dist[len(dist)-1]})
	}
	if minArray(dist) > maxDist {
		return result
	}
	for c, n := range t.children {
		result = append(result, n.search(word, c, dist, maxDist)...)
	}
	return result
}

// distFromPrefixDist returns the array with distances between prefix+c and word
// we don't need to know the prefix, just the prefix distance and the added char
func distFromPrefixDist(word string, currentChar rune, prefixDist []int) []int {
	dist := make([]int, len(prefixDist))
	dist[0] = prefixDist[0] + 1
	charCounter := 0
	for _, char := range word {
		var substCost int
		if currentChar != char {
			substCost = 1
		}
		dist[charCounter+1] = min(prefixDist[charCounter]+substCost, prefixDist[charCounter+1]+1, dist[charCounter]+1)
		charCounter++
	}
	return dist
}

func minArray(a []int) int {
	r := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] < r {
			r = a[i]
		}
	}
	return r
}

func min(i, j, k int) int {
	r := i
	if r > j {
		r = j
	}
	if r > k {
		r = k
	}
	return r
}
