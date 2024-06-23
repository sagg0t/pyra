package utils

import "strconv"

func FormatFloat[T float32 | float64](f T) string {
	f64 := float64(f)
	if f64-float64(int64(f64)) == 0 {
		return strconv.FormatInt(int64(f64), 10)
	}

	return strconv.FormatFloat(f64, 'f', -1, 32)
}
