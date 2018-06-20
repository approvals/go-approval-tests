package reporters

type fileMerge struct{}

// NewFileMergeReporter creates a new reporter for the Xcode filemerge diff tool.
func NewFileMergeReporter() Reporter {
	return &fileMerge{}
}

func (s *fileMerge) Report(approved, received string) bool {
	xs := []string{"--nosplash", "-left", received, "-right", approved}
	programName := "/Applications/Xcode.app/Contents/Applications/FileMerge.app/Contents/MacOS/FileMerge"

	return launchProgram(programName, approved, xs...)
}
