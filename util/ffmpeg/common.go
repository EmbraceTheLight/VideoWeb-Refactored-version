package ffmpeg

const (
	ffprobe = "ffprobe"
	ffmpeg  = "ffmpeg"
	ffplay  = "ffplay"
)

type globalArgs []string
type inputArgs map[string][]string
type outputArgs map[string][]string

func addArgs[T ~[]string](s T, args ...string) T {
	for _, arg := range args {
		s = append(s, arg)
	}
	return s
}

func (ga *globalArgs) addGlobal(args ...string) {
	*ga = addArgs(*ga, args...)
}

func (ia *inputArgs) addInput(inputFile string, args ...string) {
	(*ia)[inputFile] = addArgs((*ia)[inputFile], args...)
}

func (oa *outputArgs) addOutput(outputFile string, args ...string) {
	(*oa)[outputFile] = addArgs((*oa)[outputFile], args...)
}
