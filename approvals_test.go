package ApprovalTests_go

import (
	"testing"

	"github.com/approvals/ApprovalTests_go/reporters"
)

func TestMain(m *testing.M) {
	r := UseReporter(reporters.NewBeyondCompareReporter())
	defer r.Close()

	m.Run()
}

func TestVerifyStringApproval(t *testing.T) {
	r := UseReporter(reporters.NewIntelliJReporter())
	defer r.Close()

	VerifyString(t, "Hello World!")
}

func TestReporterFromSetup(t *testing.T) {
	VerifyString(t, "Hello World!")
}

func TestVerifyJSONBytes(t *testing.T) {
	jsonb := []byte("{ \"foo\": \"bar\", \"age\": 42, \"bark\": \"woof\" }")
	VerifyJSONBytes(t, jsonb)
}

func TestVerifyBadJSONBytes(t *testing.T) {
	jsonb := []byte("{ foo: \"bar\", \"age\": 42, \"bark\": \"woof\" }")
	VerifyJSONBytes(t, jsonb)
}
