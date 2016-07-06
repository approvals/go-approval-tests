package reporters

type intellij struct{}

func NewIntelliJ() Reporter {
	return &intellij{}
}

func (s *intellij) Report(approved, received string) bool {
	xs := []string{"diff", approved, received}
	programName := "C:/Program Files (x86)/JetBrains/IntelliJ IDEA 2016/bin/idea.exe"

	return launchProgram(programName, approved, xs...)
}
