package reporters

type sublimeMerge struct{}

// NewSublimeMergeReporter creates a new reporter for the SublimeMerge diff tool.
func NewSublimeMergeReporter() Reporter {
	return &sublimeMerge{}
}

func (s *sublimeMerge) Report(approved, received string) bool {
	return launchProgram("smerge", approved, "mergetool", received, approved, "-o", approved)
}
