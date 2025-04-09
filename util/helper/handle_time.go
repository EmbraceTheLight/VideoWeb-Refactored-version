package helper

import "fmt"

// SecondToTime 秒数转时间字符串以及时分秒数组，可以按需取用两个返回值中的任何一个
func SecondToTime(second int64) (string, []int64) {
	hour := second / 3600
	minute := (second % 3600) / 60
	second = second % 60
	return fmt.Sprintf("%02d:%02d:%02d", hour, minute, second), []int64{hour, minute, second}
}
