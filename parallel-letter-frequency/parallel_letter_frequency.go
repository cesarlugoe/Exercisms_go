package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap
func ConcurrentFrequency(ss []string) FreqMap {
	m1 := FreqMap{}
	m2 := FreqMap{}
	m3 := FreqMap{}

	c1 := make(chan FreqMap, 1)
	c2 := make(chan FreqMap, 1)
	c3 := make(chan FreqMap, 1)

	go calculateRuneFrecuency(ss[0], c1)
	go calculateRuneFrecuency(ss[1], c2)
	go calculateRuneFrecuency(ss[2], c3)

	m1 = <-c1
	m2 = <-c2
	m3 = <-c3

	for k, v := range m1 {
		m2[k] += v
	}

	for k, v := range m2 {
		m3[k] += v
	}

	return m3
}

func calculateRuneFrecuency(s string, c chan<- FreqMap) {
	cm := FreqMap{}
	for _, r := range s {
		cm[r]++
	}
	c <- cm
}
