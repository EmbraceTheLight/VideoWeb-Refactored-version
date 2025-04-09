package ffmpeg

import (
	"os"
	"path"
	"path/filepath"
	"util/helper/file"
)

// OtherToMP4 converts other video format to mp4 format using ffmpeg
func OtherToMP4(input string, output string) error {
	ffmpegOperator := NewFFmpeg()
	ffmpegOperator.
		AddGlobalArgs("-y").
		AddInputInfo(input).AddOutputInfo(output).
		SetAudioCodec(output, "aac").
		SetVideoCodec(output, "copy")
	return ffmpegOperator.Run()
}

// MakeDASH creates a DASH stream from an input file and saves it to an output folder.
// The input file should be .mp4 format. If the input file is not.mp4 format, it will be converted to.mp4 format first.
// If the fileName is not provided, the output file will be named "output.mpd" automatically.
func MakeDASH(inputFilePath string, outputFolderPath string, mpdFileName string) error {
	if mpdFileName == "" {
		mpdFileName = "output.mpd"
	}
	if filepath.Ext(mpdFileName) != ".mpd" {
		mpdFileName += ".mpd"
	}
	// Convert input file to mp4 format
	if filepath.Ext(inputFilePath) != ".mp4" {
		originInputFilePath := inputFilePath
		inputFilePath = file.ChangeFileExtension(inputFilePath, "mp4")
		defer func() {
			_ = os.Remove(inputFilePath)
		}()
		err := OtherToMP4(originInputFilePath, inputFilePath)
		if err != nil {
			return err
		}
	}

	// Create output folder if not exists
	err := os.MkdirAll(outputFolderPath, os.ModePerm)
	if err != nil {
		return err
	}

	// Use path.Join instead of filepath.Join to avoid windows path separator('\') issue:
	// It seems that ffmpeg command cannot handle windows path separator correctly.
	outputFilePath := path.Join(outputFolderPath, mpdFileName)
	ffmpegOperator := NewFFmpeg().
		AddGlobalArgs("-y").
		AddInputInfo(inputFilePath).
		AddOutputInfo(outputFilePath).
		SetCopyCodec(outputFilePath).ShowCommand()
	return ffmpegOperator.Run()
}
