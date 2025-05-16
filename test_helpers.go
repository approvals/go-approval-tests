package approvals

type TestFailable struct {
	failed bool
	name   string
}

func (s *TestFailable) Name() string              { return s.name }
func (s *TestFailable) Error(args ...interface{}) { s.failed = true }
func (s *TestFailable) Helper()                   {}
func (s *TestFailable) Failed() bool              { return s.failed }

func NewTestFailableWithName(name string) *TestFailable {
	return &TestFailable{name: name}
}
