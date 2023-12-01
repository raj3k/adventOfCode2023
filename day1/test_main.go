package day1

import "testing"

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func TestGetCalibrationValues(t *testing.T) {
	const input = `1abc2
	pqr3stu8vwx
	a1b2c3d4e5f
	treb7uchet
	`

	expected := []int{12, 38, 15, 77}

	received := GetCalibrationValues(input)

	if !Equal(expected, received) {
		t.Errorf("Expected: %v, Received: %v", expected, received)
	}

}
