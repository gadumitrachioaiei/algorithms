package mapper

import (
	"math/rand"
	"reflect"
	"sort"
	"strconv"
	"testing"
	"time"
)

func TestMap(t *testing.T) {
	for _, tc := range getTestCases() {
		tc.calculateExpectedWithSort()
		MapData(tc.as, tc.bs, tc.isMap, tc.fMap)
		if !reflect.DeepEqual(tc.as, tc.expectedAs) {
			t.Errorf("%s got as\n%v\nexpected\n%v\n", tc.name, tc.as, tc.expectedAs)
		}
	}
}

func TestMapNaive(t *testing.T) {
	for _, tc := range getTestCases() {
		tc.calculateExpected()
		MapNaive(tc.as, tc.bs)
		if !reflect.DeepEqual(tc.as, tc.expectedAs) {
			t.Errorf("%s got as\n%v\nexpected\n%v\n", tc.name, tc.as, tc.expectedAs)
		}
	}
}

func TestMapWithMapNaive(t *testing.T) {
	as, bs := getData(1000)
	tc := TestCase{
		as:         as,
		bs:         bs,
		expectedAs: copy(as),
	}
	MapData(tc.as, tc.bs, tc.isMap, tc.fMap)
	MapNaive(tc.expectedAs, tc.bs)
	sort.Slice(tc.expectedAs, func(i, j int) bool { return tc.expectedAs[i].id < tc.expectedAs[j].id })
	if !reflect.DeepEqual(tc.as, tc.expectedAs) {
		t.Errorf("got as\n%v\nexpected\n%v\n", tc.as, tc.expectedAs)
	}
}

func BenchmarkMapSmall(b *testing.B) {
	benchmarkMap(b, 100)
}

func BenchmarkMapMedium(b *testing.B) {
	benchmarkMap(b, 5000)
}

func BenchmarkMapBig(b *testing.B) {
	benchmarkMap(b, 50000)
}

func BenchmarkMapNaiveSmall(b *testing.B) {
	benchmarkMapNaive(b, 100)
}

func BenchmarkMapNaiveMedium(b *testing.B) {
	benchmarkMapNaive(b, 5000)
}

func BenchmarkMapNaiveBig(b *testing.B) {
	benchmarkMapNaive(b, 50000)
}

func benchmarkMap(b *testing.B, size int) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		as, bs := getData(size)
		var isMap = func(i, j int) bool {
			return as[i].id == bs[j].id
		}
		var fMap = func(i, j, k int) {
			as[i].bs = bs[j:k]
		}
		b.StartTimer()
		MapData(as, bs, isMap, fMap)
	}
}

func benchmarkMapNaive(b *testing.B, size int) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		as, bs := getData(size)
		b.StartTimer()
		MapNaive(as, bs)
	}
}

type A struct {
	id int
	bs []B
}

type B struct {
	id int
	n  string
}

type As []A

func (as As) Len() int {
	return len(as)
}

func (as As) Equal(i, j int) bool {
	return as[i].id == as[j].id
}

type Bs []B

func (bs Bs) Len() int {
	return len(bs)
}

func MapData(as []A, bs []B, isMap func(i, j int) bool, fMap func(i, j, k int)) {
	sort.Slice(as, func(i, j int) bool { return as[i].id < as[j].id })
	sort.Slice(bs, func(i, j int) bool { return bs[i].id < bs[j].id })
	Map(As(as), Bs(bs), isMap, fMap)
}

type TestCase struct {
	name       string
	as         []A
	bs         []B
	m          map[int][2]int
	expectedAs []A
}

func (tc *TestCase) isMap(i, j int) bool {
	return tc.as[i].id == tc.bs[j].id
}

func (tc *TestCase) fMap(i, j, k int) {
	tc.as[i].bs = tc.bs[j:k]
}

func (tc *TestCase) calculateExpected() {
	tc.expectedAs = copy(tc.as)
	for i := range tc.m {
		tc.expectedAs[i].bs = tc.bs[tc.m[i][0]:tc.m[i][1]]
	}
}

func (tc *TestCase) calculateExpectedWithSort() {
	tc.calculateExpected()
	sort.Slice(tc.expectedAs, func(i, j int) bool { return tc.expectedAs[i].id < tc.expectedAs[j].id })
}

func (tc *TestCase) calculateExpectedWithMapNaive() {
	asOld := copy(tc.as)
	MapNaive(tc.as, tc.bs)
	tc.expectedAs = tc.as
	tc.as = asOld
	sort.Slice(tc.expectedAs, func(i, j int) bool { return tc.expectedAs[i].id < tc.expectedAs[j].id })
}

func getTestCases() []*TestCase {
	var testCases = []*TestCase{
		{
			name: "nothing",
			as:   nil,
			bs:   nil,
		},
		{
			name: "nobs",
			as:   []A{{id: 1}, {id: 2}},
			bs:   nil,
		},
		{
			name: "tc1",
			as:   []A{{id: 1}, {id: 2}},
			bs:   []B{{id: 1, n: "n11"}, {id: 1, n: "n12"}, {id: 3, n: "n31"}},
			m:    map[int][2]int{0: [2]int{0, 2}},
		},
		{
			name: "tc2",
			as:   []A{{id: 2}, {id: 3}},
			bs:   []B{{id: 1, n: "n11"}, {id: 1, n: "n12"}, {id: 3, n: "n31"}},
			m:    map[int][2]int{1: [2]int{2, 3}},
		},
		{
			name: "tc3",
			as:   []A{{id: 1}, {id: 3}},
			bs:   []B{{id: 1, n: "n11"}, {id: 1, n: "n12"}, {id: 3, n: "n31"}},
			m:    map[int][2]int{0: [2]int{0, 2}, 1: [2]int{2, 3}},
		},
		{
			name: "tc4",
			as:   []A{{id: 2}, {id: 3}},
			bs:   []B{{id: 1, n: "n11"}, {id: 1, n: "n12"}, {id: 2, n: "n21"}, {id: 3, n: "n31"}},
			m:    map[int][2]int{0: [2]int{2, 3}, 1: [2]int{3, 4}},
		},
		{
			name: "tc5",
			as:   []A{{id: 1}, {id: 1}, {id: 2}},
			bs:   []B{{id: 1, n: "n11"}, {id: 1, n: "n12"}, {id: 2, n: "n21"}, {id: 3, n: "n31"}},
			m:    map[int][2]int{0: [2]int{0, 2}, 1: [2]int{0, 2}, 2: [2]int{2, 3}},
		},
	}
	return testCases
}

func copy(as []A) []A {
	if as == nil {
		return nil
	}
	result := make([]A, len(as))
	for i := range as {
		result[i] = as[i]
	}
	return result
}

func getData(size int) ([]A, []B) {
	var (
		as []A
		bs []B
	)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		as = append(as, A{id: i})
		bs = append(bs, B{id: i, n: "n" + strconv.Itoa(i)})
		bs = append(bs, B{id: i, n: "m" + strconv.Itoa(i)})
		if i%100 == 0 {
			as = append(as, A{id: i})
			as = append(as, A{id: size + rand.Intn(1000)})
			bID := size + 1001 + rand.Intn(1000)
			bs = append(bs, B{id: bID, n: "r" + strconv.Itoa(bID)})
		}
	}
	return as, bs
}

// MapNaive is the naive solution of our problem
func MapNaive(as []A, bs []B) {
	m := make(map[int][]B, len(bs)/3)
	for i := range bs {
		m[bs[i].id] = append(m[bs[i].id], bs[i])
	}
	for i := 0; i < len(as); i++ {
		as[i].bs = m[as[i].id]
	}
}
