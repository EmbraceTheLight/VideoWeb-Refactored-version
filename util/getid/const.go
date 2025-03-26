package getid

import "time"

var (
	baseTime int64
)

func init() {
	tmp := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano()
	baseTime = tmp / 1e6
}
