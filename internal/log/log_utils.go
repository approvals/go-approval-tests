package log

import (
	"fmt"
	"os"
	"sync"
)

var dirOnce sync.Once

const approvalTempdirectory = ".approval_tests_temp"

func InitializeTempDirectory() {
	dirOnce.Do(func() {
		// create the file and setup the parent directory if needed
		err := os.MkdirAll(approvalTempdirectory, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating directory: ", err)

			return
		}

		// create a .gitignore file containing *
		file, err := os.OpenFile(approvalTempdirectory+"/.gitignore", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			fmt.Println("Error creating file: ", err)

			return
		}
		defer file.Close()
		file.WriteString("*\n")
	})
}
