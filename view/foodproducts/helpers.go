package foodproducts

import "strconv"

func formatFloat(f float64) string {
	if f == 0 {
		return strconv.FormatFloat(f, 'f', 1, 32)
	}

	return strconv.FormatFloat(f, 'f', -1, 32)
}

// func formatFloat[T float32 | float64](f T) string {
// 	if f == 0 {
// 		return strconv.FormatFloat(float64(f), 'f', 1, 32)
// 	}
//
// 	return strconv.FormatFloat(float64(f), 'f', -1, 32)
// }
