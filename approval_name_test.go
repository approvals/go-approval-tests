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
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			var namer = getApprovalName(t)
			sanitizedName := strings.Replace(tc.name, " ", "_", -1)
			if !strings.Contains(namer.name, "TestParameterizedTestNames.") &&
				!strings.HasSuffix(namer.name, sanitizedName) {
				t.Fatalf("parameterized test name is wrong in namer, got %s", namer.name)
			}
		})
	}
}
