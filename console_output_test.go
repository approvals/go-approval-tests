package approvals

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestVerifyOutput(t *testing.T) {
	console := NewConsoleOutput()

	fmt.Print("Hello, World!")
	
	console.Close()
	
	output := console.GetOutput()
	t.Logf("Captured output: '%s'", output)
	
	console.VerifyOutput(t)
}

func TestVerifyError(t *testing.T) {
	console := NewConsoleOutput()
	defer console.Close()

	fmt.Fprint(os.Stderr, "Error message!")
	
	time.Sleep(10 * time.Millisecond)
	
	console.VerifyError(t)
}

func TestVerifyAll(t *testing.T) {
	console := NewConsoleOutput()
	defer console.Close()

	fmt.Print("Standard output")
	fmt.Fprint(os.Stderr, "Error output")
	
	time.Sleep(10 * time.Millisecond)
	
	console.VerifyAll(t)
}
