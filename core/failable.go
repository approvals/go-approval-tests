package core

// Failable is an interface wrapper around testing.T
type Failable interface {
	Error(args ...interface{})
	Name() string
	Helper()
}
