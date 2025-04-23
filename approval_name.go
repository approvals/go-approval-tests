package approvals

import (
	"github.com/approvals/go-approval-tests/core"
)

func getApprovalNameCreator() core.ApprovalNamerCreator {
	return CreateTemplatedCustomNamerCreator("{TestSourceDirectory}/{ApprovalsSubdirectory}/{TestFileName}.{TestCaseName}{AdditionalInformation}.{ApprovedOrReceived}.{FileExtension}")
}
