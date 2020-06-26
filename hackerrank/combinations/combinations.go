package combinations

// Get writes to the channel combinations (n, m)
// out is here for writing the combinations as soon as they are available
// it overwrites a previous returned combination, so the caller needs to copy it if it wants to store it
// we close this channel to signal we are done
// it also has a Reply channel, so that the caller can tell us that it wants us to stop, by closing it
func Get(n, m int, out chan chanComm) {
	defer close(out)
	if m > n {
		panic("m needs to be smaller than n")
	}
	if m == 0 {
		emit(nil, out)
		return
	}
	a := make([]int, m)
	var i int
	for i = 0; i < m; i++ {
		a[i] = i
	}
	if !emit(a, out) {
		return
	}
	i = m
	for i > 0 {
		i--
		// a[i] needs to have a small enough value so that we can so that there are enough numbers to fill up the rest of the elements in the combination
		// from i+1 to m-1 there are: m-1-i elemente
		// so a[i] + 1 + m-1-i < n , so: a[i] < n -m + i
		limit := n - m + i
		if a[i] < limit {
			a[i]++
			for j := i + 1; j < m; j++ {
				a[j] = a[j-1] + 1
			}
			for a[m-1] < n {
				if !emit(a, out) {
					return
				}
				a[m-1] = a[m-1] + 1
			}
			i = m
		}
	}
}

// chanComm represents a way to communicate combinations as soon as they are generated
type chanComm struct {
	Data  []int
	Reply chan struct{}
}

func emit(data []int, out chan chanComm) bool {
	comm := chanComm{
		Data:  data,
		Reply: make(chan struct{}),
	}
	out <- comm
	_, isOK := <-comm.Reply
	return isOK
}

// NonDivisibleSubsetSlow returns the length of the longest subset of s,
// such that the sum of any two of its numbers is not divisible by d
// it is very slow, as it loops through s ( a potentially big list ), and it can end up calculating lots of subsets.
func NonDivisibleSubsetSlow(s []int, d int) int {
	var i int
	for i = len(s); i > 1; i-- {
		receive := make(chan chanComm)
		go Get(len(s), i, receive)
		for comm := range receive {
			ss := make([]int, len(comm.Data))
			for j, index := range comm.Data {
				ss[j] = s[index]
			}
			if isCondMet(ss, d) {
				close(comm.Reply)
				return i
			}
			comm.Reply <- struct{}{}
		}
	}
	if len(s) == 0 {
		return 0
	}
	return 1
}

func isCondMet(s []int, d int) bool {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if (s[i]+s[j])%d == 0 {
				return false
			}
		}
	}
	return true
}
