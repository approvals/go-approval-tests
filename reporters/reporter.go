package reporters

// Reporter are called on failing approvals.
type Reporter interface {
	// Report is called when the approved and received file do not match.
	Report(approved, received string) bool
}

// FirstWorkingReporter reports using the first possible reporter.
type FirstWorkingReporter struct {
	Reporters []Reporter
}

func (s *FirstWorkingReporter) Report(approved, received string) bool {
	for _, reporter := range s.Reporters {
		result := reporter.Report(approved, received)
		if result {
			return true
		}
	}

	return false
}

// NewFirstWorkingReporter creates in the order reporters are passed in.
func NewFirstWorkingReporter(reporters ...Reporter) Reporter {
	return &FirstWorkingReporter{
		Reporters: reporters,
	}
}
