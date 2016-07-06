package reporters

func NewDiffReporter() Reporter {
	return NewFirstWorkingReporter(
		NewIntelliJ(),
		NewBeyondCompareReporter())
}
