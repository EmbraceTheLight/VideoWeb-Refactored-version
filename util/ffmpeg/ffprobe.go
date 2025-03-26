package ffmpeg

import (
	"fmt"
	"strings"
)

var _ FFProbeCommander = (*FFprobe)(nil) // 确保FFprobe实现接口

type FFProbeCommander interface {
	ffCommander[*FFprobe]
	// ShowEntries displays the specified entries of the input file.
	ShowEntries(inputFilePath string, entries ...string) *FFprobe
	SetOutputFormat(inputFilePath string, outputFormat string) *FFprobe
}

func NewFFprobe() *FFprobe {
	ffprobeOperator := &FFprobe{}
	ffprobeOperator.ffBase = newFFBase()
	ffprobeOperator.ffBase.command = ffprobe
	return ffprobeOperator
}

type FFprobe struct {
	*ffBase
}

func (f *FFprobe) AddGlobalArgs(args ...string) *FFprobe {
	f.ffBase.AddGlobalArgs(args...)
	return f
}

func (f *FFprobe) AddInputInfo(inputFilePath string, args ...string) *FFprobe {
	f.ffBase.AddInputInfo(inputFilePath, args...)
	return f
}

func (f *FFprobe) AddOutputInfo(outputFilePath string, args ...string) *FFprobe {
	f.ffBase.AddOutputInfo(outputFilePath, args...)
	return f
}

func (f *FFprobe) SetAudioCodec(outputFilePath string, codec string) *FFprobe {
	f.ffBase.SetAudioCodec(outputFilePath, codec)
	return f
}

func (f *FFprobe) SetVideoCodec(outputFilePath string, codec string) *FFprobe {
	f.ffBase.SetVideoCodec(outputFilePath, codec)
	return f
}

func (f *FFprobe) SetLogLevel(level string) *FFprobe {
	f.ffBase.SetLogLevel(level)
	return f
}

func (f *FFprobe) ShowEntries(inputFilePath string, entries ...string) *FFprobe {
	args := []string{"-show_entries"}
	entriesStr := strings.Join(entries, ",")
	args = append(args, "format="+entriesStr)
	f.AddInputInfo(inputFilePath, args...)
	return f
}

func (f *FFprobe) SetOutputFormat(inputFilePath string, outputFormat string) *FFprobe {
	args := []string{"-output_format"}
	args = append(args, outputFormat)
	f.ffBase.AddInputInfo(inputFilePath, args...)
	return f
}

func (f *FFprobe) ShowCommand() *FFprobe {
	fmt.Println(f.ffBase.command + " " + f.ffBase.BuildCommand())
	return f
}
