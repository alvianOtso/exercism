package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	i := 0
	for n > 0 {
		if n == 1 {
			return i, nil
		}
		if n%2 == 0 {
			n = n / 2
		} else {
			n = (n * 3) + 1
		}
		i++
	}

	return 0, errors.New("error")
}
