package reporters

import (
	"fmt"
	"os"

	"github.com/approvals/go-approval-tests/internal/log"
	"github.com/approvals/go-approval-tests/utils"
)

var (
	directory                = ".approval_tests_temp"
	filenameWithoutExtension = directory + "/approval_script"
	filename                 = ""
)

type approvalScript struct{}

// Deprecated: instead run .approval_tests_temp/approve_all.py
func NewReporterThatCreatesAnApprovalScript() Reporter {
	initializeFile()
	return &approvalScript{}
}

func (s *approvalScript) Report(approved, received string) bool {
	move := getMoveCommandText(approved, received) + "\n"

	utils.AppendToFile(filename, move)

	return true
}

func initializeFile() {
	if filename != "" {
		return
	}

	filename = filenameWithoutExtension + ".sh"

	log.InitializeTempDirectory()

	// create the file and make it executable in one step
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Println("Error creating file: ", err)
		return
	}
	file.Close()

	utils.AppendToFile(filename, "#!/bin/bash\n")

	fmt.Println("You can run the approval script by executing: ", filename)

}
