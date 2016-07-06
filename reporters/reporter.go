package reporters

type Reporter interface {
	Report(approved, received string) bool
}

