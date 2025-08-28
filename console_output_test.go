package approvals

import (
	"fmt"
	"os"
	"testing"
)

func TestVerifyOutput(t *testing.T) {
	console := NewConsoleOutput()
	defer console.Close()

	fmt.Print("Hello, World!")
	
	console.VerifyOutput(t)
}

func TestVerifyError(t *testing.T) {
	console := NewConsoleOutput()
	defer console.Close()

	fmt.Fprint(os.Stderr, "Error message!")
	
	console.VerifyError(t)
}

func TestVerifyAll(t *testing.T) {
	console := NewConsoleOutput()
	defer console.Close()

	fmt.Print("Standard output")
	fmt.Fprint(os.Stderr, "Error output")
	
	console.VerifyAll(t)
}
