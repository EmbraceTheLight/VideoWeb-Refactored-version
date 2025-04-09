package ffmpeg

import (
	"strconv"
	"strings"
)

// GetVideoDuration gets the duration of the video file.
// The duration is rounded off to the nearest integer in seconds.
func GetVideoDuration(inputFilePath string) (int64, error) {
	ffprobeOperator := NewFFprobe()

	outData, err := ffprobeOperator.
		AddInputInfo(inputFilePath).
		AddGlobalArgs("-v", "quiet").
		ShowEntries(inputFilePath, "duration").
		SetOutputFormat(inputFilePath, "csv=p=0").
		RunCombinedOutput()
	if err != nil {
		return 0, err
	}
	tmp, _ := strconv.ParseFloat(strings.TrimRight(string(outData), "\r\n"), 64) //注意去除字符串末尾的\r\n
	return roundOff(tmp), nil
}
