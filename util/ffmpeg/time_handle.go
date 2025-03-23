package ffmpeg

import "strconv"

const prec = 3 // precision of float time

// roundOff rounds off the float value to the nearest integer
func roundOff(f float64) int64 {
	var decimal = f - float64(int(f))
	if decimal >= 0.5 {
		return int64(f) + 1
	}
	return int64(f)
}

// formatFloatTime formats the float time to string with 3 decimal places
func formatFloatTime(t float64, precision int) string {
	if precision <= 0 {
		precision = prec
	}
	return strconv.FormatFloat(t, 'f', precision, 64)
}
