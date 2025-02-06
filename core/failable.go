package core

// Failable is an interface wrapper around testing.T
type Failable interface {
	Fail()
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Name() string
	Log(args ...interface{})
	Logf(format string, args ...interface{})
	Helper()
}
