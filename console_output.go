package approvals

import (
	"bytes"
	"fmt"
	"os"
	"sync"

	"github.com/approvals/go-approval-tests/core"
)

type ConsoleOutput struct {
	originalStdout *os.File
	originalStderr *os.File
	stdoutBuffer   *bytes.Buffer
	stderrBuffer   *bytes.Buffer
	stdoutReader   *os.File
	stderrReader   *os.File
	stdoutWriter   *os.File
	stderrWriter   *os.File
	wg             sync.WaitGroup
	mu             sync.RWMutex
	closed         bool
}

func NewConsoleOutput() *ConsoleOutput {
	console := &ConsoleOutput{
		originalStdout: os.Stdout,
		originalStderr: os.Stderr,
		stdoutBuffer:   &bytes.Buffer{},
		stderrBuffer:   &bytes.Buffer{},
	}

	stdoutReader, stdoutWriter, _ := os.Pipe()
	stderrReader, stderrWriter, _ := os.Pipe()

	console.stdoutReader = stdoutReader
	console.stderrReader = stderrReader
	console.stdoutWriter = stdoutWriter
	console.stderrWriter = stderrWriter

	os.Stdout = stdoutWriter
	os.Stderr = stderrWriter

	console.wg.Add(2)
	go func() {
		defer console.wg.Done()
		buf := make([]byte, 1024)
		for {
			n, err := stdoutReader.Read(buf)
			if n > 0 {
				console.mu.Lock()
				console.stdoutBuffer.Write(buf[:n])
				console.mu.Unlock()
			}
			if err != nil {
				break
			}
		}
	}()
	go func() {
		defer console.wg.Done()
		buf := make([]byte, 1024)
		for {
			n, err := stderrReader.Read(buf)
			if n > 0 {
				console.mu.Lock()
				console.stderrBuffer.Write(buf[:n])
				console.mu.Unlock()
			}
			if err != nil {
				break
			}
		}
	}()

	return console
}

func (c *ConsoleOutput) ensureClosed() {
	c.mu.Lock()
	if c.closed {
		c.mu.Unlock()
		return
	}
	c.stdoutWriter.Close()
	c.stderrWriter.Close()
	c.closed = true
	c.mu.Unlock()
	
	c.wg.Wait()
}

func (c *ConsoleOutput) GetOutput() string {
	c.ensureClosed()
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.stdoutBuffer.String()
}

func (c *ConsoleOutput) GetError() string {
	c.ensureClosed()
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.stderrBuffer.String()
}

func (c *ConsoleOutput) VerifyOutput(t core.Failable) {
	VerifyString(t, c.GetOutput())
}

func (c *ConsoleOutput) VerifyError(t core.Failable) {
	VerifyString(t, c.GetError())
}

func (c *ConsoleOutput) VerifyAll(t core.Failable) {
	combined := fmt.Sprintf("STDOUT:\n%s\nSTDERR:\n%s", c.GetOutput(), c.GetError())
	VerifyString(t, combined)
}

func (c *ConsoleOutput) Close() error {
	c.ensureClosed()
	
	c.stdoutReader.Close()
	c.stderrReader.Close()
	
	os.Stdout = c.originalStdout
	os.Stderr = c.originalStderr
	
	return nil
}
