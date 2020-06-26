package combinations

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/gadumitrachioaiei/algorithms/hackerrank/nondivisiblesubset"
)

func TestNonDivisibleSubsetSlow(t *testing.T) {
	s := make([]int, 30)
	d := 47
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(s); i++ {
		s[i] = rand.Intn(100)
	}
	fmt.Println(s)
	calculated := NonDivisibleSubsetSlow(s, d)
	expected := nondivisiblesubset.Find(s, d)
	if calculated != expected {
		t.Fatalf("got: %d, expected: %d, s:\n%v\n", calculated, expected, s)
	}
	fmt.Println(calculated)
}
