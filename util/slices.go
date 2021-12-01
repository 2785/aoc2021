package util

import "strconv"

func IntsFromStrings(s []string) ([]int, error) {
	r := make([]int, len(s))
	for i, v := range s {
		num, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		r[i] = num
	}
	return r, nil
}
