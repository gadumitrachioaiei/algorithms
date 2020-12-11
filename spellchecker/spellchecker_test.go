package spellchecker

import (
	"bufio"
	"os"
	"reflect"
	"sort"
	"testing"
)

func TestDist(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case0",
			args{"cuvant", "cuvant"},
			0,
		},
		{
			"case1",
			args{"thou shalt", "you should"},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Dist(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("Dist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearch(t *testing.T) {
	f, err := os.Open("/usr/share/dict/words")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	trie := NewTrieNode()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		trie.InsertWord(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		t.Fatal(err)
	}
	expectedWordsMatch := []WordMatch{
		{"gobber", 1}, {"goofer", 1}, {"goober", 0}, {"goobers", 1},
		{"gooder", 1}, {"gooier", 1}, {"Goober", 1},
	}
	wordsMatch := trie.Search("goober", 1)
	sort.Slice(expectedWordsMatch, func(i, j int) bool {
		return expectedWordsMatch[i].W < expectedWordsMatch[j].W
	})
	sort.Slice(wordsMatch, func(i, j int) bool {
		return wordsMatch[i].W < wordsMatch[j].W
	})
	if !reflect.DeepEqual(wordsMatch, expectedWordsMatch) {
		t.Fatalf("got matches: \n%v\n, expected: \n%v\n", wordsMatch, expectedWordsMatch)
	}
}
