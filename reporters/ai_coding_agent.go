package reporters

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/approvals/go-approval-tests/utils"
)

// AICodingAgentEnvVars are the environment variables that indicate the tests are
// being run by an AI coding agent. Append to it to support an agent not listed here.
var AICodingAgentEnvVars = []string{
	"AI_AGENT",               // generic opt-in
	"APPROVAL_TESTS_AGENT",   // generic opt-in
	"CLAUDECODE",             // Claude Code
	"CLAUDE_CODE_ENTRYPOINT", // Claude Code
	"CODEX_SANDBOX",          // OpenAI Codex CLI
	"CURSOR_AGENT",           // Cursor
	"CURSOR_TRACE_ID",        // Cursor
	"GEMINI_CLI",             // Gemini CLI
	"OPENCODE",               // opencode
}

type aiCodingAgent struct{}

// NewAICodingAgentReporter creates a new reporter for AI coding agents.
//
// The reporter checks the environment variables in AICodingAgentEnvVars and, when
// one is set, prints both file contents and the command to approve the result
// rather than launching a diff tool the agent cannot see.
func NewAICodingAgentReporter() Reporter {
	return &aiCodingAgent{}
}

// IsAICodingAgent reports whether an AI coding agent environment was detected.
func IsAICodingAgent() bool {
	for _, key := range AICodingAgentEnvVars {
		value, exists := os.LookupEnv(key)
		if !exists || value == "" {
			continue
		}
		// Values are often names rather than booleans, so only an explicit false opts out.
		if set, err := strconv.ParseBool(value); err == nil && !set {
			continue
		}
		return true
	}

	return false
}

func (s *aiCodingAgent) Report(approved, received string) bool {
	if !IsAICodingAgent() {
		return false
	}

	approvedFull, _ := filepath.Abs(approved)
	receivedFull, _ := filepath.Abs(received)

	status := "approval files did not match"
	if !utils.DoesFileExist(approved) {
		status = "result never approved"
	}

	fmt.Printf("=== APPROVAL TEST FAILED ===\n")
	fmt.Printf("status: %s\n", status)
	fmt.Printf("approved: %s\n", approvedFull)
	fmt.Printf("received: %s\n", receivedFull)
	printAgentSection("APPROVED", approvedFull)
	printAgentSection("RECEIVED", receivedFull)
	fmt.Printf("--- APPROVE WITH ---\n%s\n", getMoveCommandText(approved, received))
	fmt.Printf("=== END APPROVAL TEST FAILED ===\n")

	return true
}

func printAgentSection(label, path string) {
	content, err := utils.ReadFile(path)
	if err != nil {
		content = "** file missing **\n"
	}
	fmt.Printf("--- %s ---\n%s", label, content)
	if len(content) > 0 && content[len(content)-1] != '\n' {
		fmt.Println()
	}
}
