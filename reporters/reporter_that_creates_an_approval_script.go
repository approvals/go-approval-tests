package reporters

import (
	"fmt"
	"os"

	"github.com/approvals/go-approval-tests/utils"
)

var (
	directory                = ".approval_tests_temp"
	filenameWithoutExtention = directory + "/approval_script"
	filename                 = ""
)

type approvalScript struct{}

// NewAllFailingTestReporter copies move file command to your clipboard
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

	filename = filenameWithoutExtention + ".sh"

	// create the file and setup the parent directory if needed
	err := os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory: ", err)
		return
	}

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
