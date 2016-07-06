package reporters

type Reporter interface {
	Report(approved, received string) bool
}

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

func NewFirstWorkingReporter(reporters ...Reporter) Reporter {
	return &FirstWorkingReporter{
		Reporters: reporters,
	}
}
