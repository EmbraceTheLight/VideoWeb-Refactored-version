package ffmpeg

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

// MakeDASH creates a DASH stream from an input file and saves it to an output folder
func MakeDASH(inputFilePath string, outputFolderPath string, fileName string) error {
	if fileName == "" {
		fileName = "output.mpd"
	}

	//Get absolute path of output folder,
	//because the releative path will be cleaned by the filepath.Join() function
	if outputFolderPath[len(outputFolderPath)-1] == '/' {
		outputFolderPath = outputFolderPath[:len(outputFolderPath)-1]
	}
	outputFilePath := outputFolderPath + "/" + fileName
	ffmpegOperator := NewFFmpeg()
	ffmpegOperator.
		AddGlobalArgs("-y").
		AddInputInfo(inputFilePath).
		AddOutputInfo(outputFilePath).
		SetCopyCodec(outputFilePath).ShowCommand()
	return ffmpegOperator.Run()
}
