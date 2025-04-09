package ffmpeg

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

type ffCommander[T interface{ *ffBase | *FFprobe | *FFmpeg }] interface {
	AddGlobalArgs(args ...string) T
	AddInputInfo(inputFilePath string, args ...string) T
	AddOutputInfo(outputFilePath string, args ...string) T
	SetAudioCodec(outputFilePath string, codec string) T
	SetVideoCodec(outputFilePath string, codec string) T
	SetLogLevel(level string) T
	ShowCommand() T
	Run() error
	RunCombinedOutput() ([]byte, error)
}

type ffBase struct {
	// command is the ffmpeg-family command name.
	// It contains ffmpeg, ffprobe, and ffplay.
	command string

	// globalArgs are the global arguments for ffmpeg-family command.
	globalArgs globalArgs

	// inputFileArgs are the input arguments for ffmpeg-family command.
	// It is a map, the key is the input file path, the value is the arguments for this input file.
	inputFileArgs inputArgs

	// outputArgs are the output arguments for ffmpeg-family command.
	// It is a map, the key is the input file path, the value is the arguments for this input file.
	outputArgs outputArgs

	outputFormat string
}

func newFFBase() *ffBase {
	return &ffBase{
		globalArgs:    globalArgs{},
		inputFileArgs: make(inputArgs),
		outputArgs:    make(outputArgs),
	}
}

func (f *ffBase) AddInputInfo(inputFilePath string, arguments ...string) *ffBase {
	f.inputFileArgs.addInput(inputFilePath, arguments...)
	return f
}

func (f *ffBase) AddOutputInfo(outputFilePath string, arguments ...string) *ffBase {
	f.outputArgs.addOutput(outputFilePath, arguments...)
	return f
}

func (f *ffBase) AddGlobalArgs(args ...string) *ffBase {
	f.globalArgs.addGlobal(args...)
	return f
}

func (f *ffBase) SetLogLevel(level string) *ffBase {
	f.AddGlobalArgs("-v", level)
	return f
}

func (f *ffBase) SetVideoCodec(outputFile string, codec string) *ffBase {
	f.AddOutputInfo(outputFile, "-c:v", codec)
	return f
}

func (f *ffBase) SetAudioCodec(outputFile string, codec string) *ffBase {
	f.AddOutputInfo(outputFile, "-c:a", codec)
	return f
}

func (f *ffBase) Run() error {
	args := f.buildArguments()
	command := exec.Command(f.command, args...)
	var stderr bytes.Buffer
	command.Stderr = &stderr

	err := command.Run()
	if err != nil {
		return errors.New(stderr.String())
	}
	return nil
}

func (f *ffBase) RunCombinedOutput() ([]byte, error) {
	args := f.buildArguments()
	command := exec.Command(f.command, args...)
	output, err := command.CombinedOutput()
	return output, err
}

func (f *ffBase) BuildCommand() string {
	return strings.Join(f.buildArguments(), " ")
}

func (f *ffBase) ShowCommand() *ffBase {
	fmt.Println(f.command + " " + f.BuildCommand())
	return f
}

func (f *ffBase) buildArguments() []string {
	var args []string

	// add global arguments
	args = append(args, f.globalArgs...)

	// add input arguments
	for inputFilePath, arguments := range f.inputFileArgs {
		args = append(args, arguments...)
		args = append(args, "-i", inputFilePath)
	}

	// add output arguments
	for outputFilePath, arguments := range f.outputArgs {
		args = append(args, arguments...)
		args = append(args, outputFilePath)
	}

	return args
}
