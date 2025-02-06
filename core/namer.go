package core

type ApprovalNamer interface {
	GetName() string
	GetReceivedFile(extWithDot string) string
	GetApprovalFile(extWithDot string) string
}

type ApprovalNamerCreator func(t Failable) ApprovalNamer
