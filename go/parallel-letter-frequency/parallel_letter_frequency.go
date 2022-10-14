package letter

import (
	"sync"
)

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

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	// fmt.Println("Waiting 0")
	var wg sync.WaitGroup

	channel := make(chan FreqMap, len(l))

	counter := make(FreqMap)

	for _, s := range l {
		wg.Add(1)
		go func(t string) {
			defer wg.Done()
			channel <- Frequency(t)
		}(s)
	}

	wg.Wait()
	close(channel)
	for fm := range channel {
		for letter, count := range fm {
			if counter[letter] > 0 {
				counter[letter] += count
			} else {
				counter[letter] = count
			}
		}
	}
	return counter
}
