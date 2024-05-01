package lsproduct

import (
	"errors"
	"strconv"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if len(digits) < span || span < 0 {
		return 0, errors.New("span must be smaller than string length")
	}

	series := []string{}
	for i := 0; i < len(digits); i++ {
		if i+span > len(digits) {
			break
		}
		series = append(series, digits[i:i+span])
	}
	
	max := 0
	for _, digit := range series {
		total := 1
		for _, r := range digit {
			num, err := strconv.Atoi(string(r))
			if err != nil {
				return 0, err
			}
			total *= num
		}
		if total > max {
			max = total
		}
	}

	return int64(max), nil
}
