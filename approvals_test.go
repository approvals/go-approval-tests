package approvaltests

import (
	"os"
	"testing"

	"github.com/approvals/go-approval-tests/reporters"
)

func TestMain(m *testing.M) {
	r := UseReporter(reporters.NewBeyondCompareReporter())
	defer r.Close()

	os.Exit(m.Run())
}

func TestVerifyStringApproval(t *testing.T) {
	r := UseReporter(reporters.NewIntelliJReporter())
	defer r.Close()

	VerifyString(t, "Hello World!")
}

func TestReporterFromSetup(t *testing.T) {
	VerifyString(t, "Hello World!")
}

func TestVerifyJSONStruct(t *testing.T) {
	json := struct {
		Title string
		Name  string
		Age   int
	}{
		Title: "Hello World!",
		Name:  "Peter Pan",
		Age:   100,
	}

	VerifyJSONStruct(t, json)
}

func TestVerifyJSONBytes(t *testing.T) {
	jsonb := []byte("{ \"foo\": \"bar\", \"age\": 42, \"bark\": \"woof\" }")
	VerifyJSONBytes(t, jsonb)
}

func TestVerifyBadJSONBytes(t *testing.T) {
	jsonb := []byte("{ foo: \"bar\", \"age\": 42, \"bark\": \"woof\" }")
	VerifyJSONBytes(t, jsonb)
}

func TestVerifyMap(t *testing.T) {
	m := map[string]string{
		"dog": "bark",
		"cat": "meow",
	}

	VerifyMap(t, m)
}
