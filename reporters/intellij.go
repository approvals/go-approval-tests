package reporters

import (
	"fmt"
	"os/exec"

	"github.com/Approvals/ApprovalTests_go/utils"
)

type intellij struct {
}

func NewIntelliJ() Reporter {
	return &intellij{}
}

func (s *intellij) Report(approved, received string) bool {
	programName := "C:/Program Files (x86)/JetBrains/IntelliJ IDEA 2016/bin/idea.exe"

	if !utils.DoesFileExist(programName) {
		return false
	}

	utils.EnsureExists(approved)

	cmd := exec.Command(programName, "diff", received, approved)
	cmd.Start()

	err := cmd.Wait()
	if err != nil {
		panic(fmt.Sprintf("err=%s", err))
	}

	return true
}
