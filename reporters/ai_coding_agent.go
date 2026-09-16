package reporters

import (
	"os"
	"strconv"
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

func NewAICodingAgentReporter() Reporter {
	return &aiCodingAgent{}
}

func IsAICodingAgent() bool {
	for _, key := range AICodingAgentEnvVars {
		value, exists := os.LookupEnv(key)
		if !exists || value == "" {
			continue
		}
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

	return NewSystemoutReporter().Report(approved, received)
}
