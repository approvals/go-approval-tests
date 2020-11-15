package approvals

import (
	"strings"
	"testing"
)

func TestNamer(t *testing.T) {
	var namer = getApprovalName(t)
	if !strings.HasSuffix(namer.name, "TestNamer") {
		t.Fatalf("test name is wrong in namer, got %s", namer.name)
	}
}

func TestNamerFilename(t *testing.T) {
	var namer = getApprovalName(t)
	if !strings.HasSuffix(namer.fileName, "approval_name_test.go") {
		t.Fatalf("test filename is wrong in namer, got %s", namer.fileName)
	}
}

func TestParameterizedTestNames(t *testing.T) {
	for _, tc := range ExampleParameterizedTestcases {
		t.Run(tc.name, func(t *testing.T) {
			var namer = getApprovalName(t)
			if !strings.HasSuffix(namer.name, "TestParameterizedTestNames."+tc.name) {
				t.Fatalf("test name is wrong in namer, got %s", namer.name)
			}
		})
	}
}
