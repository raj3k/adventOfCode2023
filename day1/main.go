package day1

import (
	"strconv"
	"unicode"
)

func GetCalibrationValues(s string) []int {
	result := []int{}
	n := ""
	for _, c := range s {
		if !unicode.IsSpace(c) {
			if unicode.IsDigit(c) {
				if len(n) != 2 {
					n += string(c)
				} else {
					n = n[0:1] + string(c)
				}
			}
		} else if c == '\n' {
			if len(n) == 1 {
				n = n[0:1] + string(n)
			}
			num, _ := strconv.Atoi(n)
			result = append(result, num)
			n = ""
		} else {
			continue
		}
	}
	return result
}
