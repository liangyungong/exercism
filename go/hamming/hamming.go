package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Length not equal!")
	}

	diff := 0

	for idx, _ := range a {
		if a[idx] != b[idx] {
			diff++
		}
	}
	return diff, nil
}
