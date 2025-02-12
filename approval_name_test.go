package approvals

import (
	"strings"
	"testing"
)

func TestNamer(t *testing.T) {
	name, _ := getApprovalName(t)
	if !strings.HasSuffix(name, "TestNamer") {
		t.Fatalf("test name is wrong in namer, got %s", name)
	}
}

func TestNamerFilename(t *testing.T) {
	_, fileName := getApprovalName(t)
	if !strings.HasSuffix(fileName, "approval_name_test.go") {
		t.Fatalf("test filename is wrong in namer, got %s", fileName)
	}
}

func TestParameterizedTestNames(t *testing.T) {
	for _, tc := range ExampleParameterizedTestcases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			name, _ := getApprovalName(t)
			sanitizedName := strings.Replace(tc.name, " ", "_", -1)
			if !strings.Contains(name, "TestParameterizedTestNames.") &&
				!strings.HasSuffix(name, sanitizedName) {
				t.Fatalf("parameterized test name is wrong in namer, got %s", name)
			}
		})
	}
}
