package ffmpeg

import "fmt"

var _ FFmpegCommander = (*FFmpeg)(nil)

type FFmpegCommander interface {
	ffCommander[*FFmpeg]
	SetSegmentTime(outputFilePath string, duration float64) *FFmpeg
	SetFormatDash(outputFilePath string) *FFmpeg
}

type FFmpeg struct {
	*ffBase
}

func NewFFmpeg() *FFmpeg {
	ffmpegOperator := &FFmpeg{}
	ffmpegOperator.ffBase = newFFBase()
	ffmpegOperator.ffBase.command = ffmpeg
	return ffmpegOperator
}

func (f *FFmpeg) AddGlobalArgs(args ...string) *FFmpeg {
	f.ffBase.globalArgs = append(f.ffBase.globalArgs, args...)
	return f
}

func (f *FFmpeg) AddInputInfo(inputFilePath string, args ...string) *FFmpeg {
	f.ffBase.AddInputInfo(inputFilePath, args...)
	return f
}

func (f *FFmpeg) AddOutputInfo(outputFilePath string, args ...string) *FFmpeg {
	f.ffBase.AddOutputInfo(outputFilePath, args...)
	return f
}

func (f *FFmpeg) SetAudioCodec(outputFilePath string, codec string) *FFmpeg {
	f.ffBase.SetAudioCodec(outputFilePath, codec)
	return f
}

func (f *FFmpeg) SetVideoCodec(outputFilePath string, codec string) *FFmpeg {
	f.ffBase.SetVideoCodec(outputFilePath, codec)
	return f
}

func (f *FFmpeg) SetCopyCodec(outputFilePath string) *FFmpeg {
	f.ffBase.SetVideoCodec(outputFilePath, "copy")
	f.ffBase.SetAudioCodec(outputFilePath, "copy")
	return f
}

func (f *FFmpeg) SetLogLevel(level string) *FFmpeg {
	f.ffBase.SetLogLevel(level)
	return f
}

func (f *FFmpeg) SetSegmentTime(outputFilePath string, duration float64) *FFmpeg {
	f.ffBase.AddOutputInfo(outputFilePath, "-segment_time", formatFloatTime(duration, 0))
	return f
}

func (f *FFmpeg) SetFormatDash(outputFilePath string) *FFmpeg {
	f.ffBase.AddOutputInfo(outputFilePath, "-f", "dash")
	return f
}

func (f *FFmpeg) ShowCommand() *FFmpeg {
	fmt.Println(f.ffBase.command + " " + f.ffBase.BuildCommand())
	return f
}
