package core

import "io"

type ApprovalNamer interface {
	Compare(approvalFile, receivedFile string, reader io.Reader) error
	GetReceivedFile(extWithDot string) string
	GetApprovalFile(extWithDot string) string
}
