# ApprovalTests.go

ApprovalTests for go



# Golden master Verification Library
ApprovalTests allows for easy testing of larger objects, strings and anytthing else that can be saved to a file (images, sounds, csv,  etc...)

#Examples
##In Project
Note: ApprovalTests uses approvaltests to test itself. Therefore are many examples in the code itself.

 * [approvals_test.go](approvals_test.go)

##JSON


```go
func TestVerifyJSON(t *testing.T) {
	jsonb := []byte("{ \"foo\": \"bar\", \"age\": 42, \"bark\": \"woof\" }")
	VerifyJSONBytes(t, jsonb)
}
```
Matches file: approvals_test.TestVerifyJSON.recieved.json

```json
{ \"foo\": \"bar\", \"age\": 42, \"bark\": \"woof\" }
```
