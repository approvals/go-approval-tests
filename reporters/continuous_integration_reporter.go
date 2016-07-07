package reporters

type continuousIntegration struct{}

func NewContinuousIntegrationReporter() Reporter {
	return &continuousIntegration{}
}

func (s *continuousIntegration) Report(approved, received string) bool {
	return false
}
