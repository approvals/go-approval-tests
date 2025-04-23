package approvals

import (
	"fmt"
	"testing"
)

func TestWithParameters(t *testing.T) {
	t.Parallel()

	values := []string{"Test1", "Test2"}
	for _, value := range values {
		result := fmt.Sprintf("Testing with parameter: %s", value)
		VerifyString(t, result, Options().ForFile().WithAdditionalInformation(value))
	}
}


