package ffmpeg_test

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
	"util/ffmpeg"
)

func TestConvertToMP4(t *testing.T) {
	inputFilePath := "./testfiles/input/E3Trailer.avi"
	outputFilePath := "./testfiles/output/E3Trailer.mp4"
	ffmpegOperator := ffmpeg.NewFFmpeg()
	ffmpegOperator.AddGlobalArgs("-y").
		AddInputInfo(inputFilePath).
		AddOutputInfo(outputFilePath).
		SetAudioCodec(outputFilePath, "aac").
		SetVideoCodec(outputFilePath, "copy")
	err := ffmpegOperator.Run()
	require.NoError(t, err)

	err = ffmpeg.OtherToMP4(inputFilePath, outputFilePath)
	require.NoError(t, err)

}

func TestGetVideoDuration(t *testing.T) {
	inputFilePath := "./testfiles/input/E3Trailer.avi"
	duration, err := ffmpeg.GetVideoDuration(inputFilePath)
	require.NoError(t, err)
	fmt.Println(duration)
}

func TestMakeDash(t *testing.T) {
	inputFilePath := "./testfiles/input/MakingOfMusic.avi"
	outputFilePath := "./testfiles/dash"
	err := ffmpeg.MakeDASH(inputFilePath, outputFilePath, "dash.mpd")
	require.NoError(t, err)
}
