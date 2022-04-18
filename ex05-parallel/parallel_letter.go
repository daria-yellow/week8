package letter

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, val := range s {
		m[val] += 1
	}
	return m
}

func ConcurrentFrequency(ss []string) FreqMap {
	ch := make(chan FreqMap)

	f := func(s string) {
		ch <- Frequency(s)
	}

	for i, _ := range ss {
		go f(ss[i])
	}

	m := FreqMap{}
	for _ = range ss {
		for i, val := range <-ch {
			m[i] += val
		}
	}
	return m
}
