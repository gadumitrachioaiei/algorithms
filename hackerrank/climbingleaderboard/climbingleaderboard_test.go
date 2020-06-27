package climbingleaderboard

import (
	"io/ioutil"
	"reflect"
	"strconv"
	"testing"
)

func TestClimbingLeaderboard(t *testing.T) {
	type args struct {
		scores []int32
		alice  []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			"case1",
			args{[]int32{3, 2, 1}, []int32{1, 5, 2}},
			[]int32{3, 1, 2},
		},
		{
			"case2",
			args{[]int32{100, 100, 50, 40, 40, 20, 10}, []int32{5, 25, 50, 120}},
			[]int32{6, 4, 2, 1},
		},
		{
			"case3",
			args{[]int32{100, 90, 90, 80, 75, 60}, []int32{50, 65, 77, 90, 102}},
			[]int32{6, 5, 4, 2, 1},
		},
		{
			"case4",
			args{[]int32{3, 2, 1, 1}, []int32{1, 5, 2}},
			[]int32{3, 1, 2},
		},
		{
			"case5",
			args{[]int32{3}, []int32{1, 5, 2}},
			[]int32{2, 1, 2},
		},
		{
			"case5",
			args{[]int32{}, []int32{1, 5, 2}},
			[]int32{1, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbingLeaderboard(tt.args.scores, tt.args.alice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("climbingLeaderboard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClimbingLeaderboardLarge(t *testing.T) {
	data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		t.Fatal(err)
	}
	var scores []int32
	var alice []int32
	lastIndex := 0
	newLines := 0
	data = append(data, ' ')
	for i, c := range data {
		if newLines == 1 {
			if c == '\n' {
				newLines = 2
				lastIndex = i + 1
			}
			continue
		}
		if c == '\n' {
			newLines++
		}
		if c == ' ' || c == '\n' {
			score, err := strconv.Atoi(string(data[lastIndex:i]))
			if err != nil {
				t.Fatal(lastIndex, i, err)
			}
			if newLines == 2 {
				alice = append(alice, int32(score))
			} else {
				scores = append(scores, int32(score))
			}
			lastIndex = i + 1
		}
	}
	ranks := climbingLeaderboard(scores, alice)
	expected := make([]int32, len(alice))
	for i := 0; i < len(expected); i++ {
		expected[i] = 199784
	}
	if !reflect.DeepEqual(ranks, expected) {
		t.Fatal("not expected")
	}
}
