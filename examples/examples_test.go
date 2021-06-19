package approvals_test

import (
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func TestExampleVerifyString(t *testing.T) {
	approvals.VerifyString(t, "Hello World!")
	printFileContent("examples_test.TestExampleVerifyString.received.txt")

	// Output:
	// This produced the file examples_test.TestExampleVerifyString.received.txt
	// It will be compared against the examples_test.ExampleVerifyString.approved.txt file
	// and contains the text:
	//
	// Hello World!
}

func TestExampleVerifyAllCombinationsFor2(t *testing.T) {
	letters := []string{"aaaaa", "bbbbb", "ccccc"}
	numbers := []int{2, 3}

	functionToTest := func(text interface{}, length interface{}) string {
		return text.(string)[:length.(int)]
	}

	approvals.VerifyAllCombinationsFor2(t, "substring", functionToTest, letters, numbers)
	printFileContent("examples_test.TestExampleVerifyAllCombinationsFor2.received.txt")
	// Output:
	// This produced the file examples_test.TestExampleVerifyAllCombinationsFor2.received.txt
	// It will be compared against the examples_test.ExampleVerifyAllCombinationsFor2.approved.txt file
	// and contains the text:
	//
	// substring
	//
	//
	// [aaaaa,2] => aa
	// [aaaaa,3] => aaa
	// [bbbbb,2] => bb
	// [bbbbb,3] => bbb
	// [ccccc,2] => cc
	// [ccccc,3] => ccc
}

func TestExampleVerifyAllCombinationsFor2_withSkip(t *testing.T) {
	words := []string{"stack", "fold"}
	otherWords := []string{"overflow", "trickle"}

	functionToTest := func(firstWord interface{}, secondWord interface{}) string {
		first := firstWord.(string)
		second := secondWord.(string)
		if first+second == "stackoverflow" {
			return approvals.SkipThisCombination
		}
		return first + second
	}

	approvals.VerifyAllCombinationsFor2(t, "combineWords", functionToTest, words, otherWords)
	printFileContent("examples_test.TestExampleVerifyAllCombinationsFor2_withSkip.received.txt")
	// Output:
	// This produced the file examples_test.TestExampleVerifyAllCombinationsFor2_withSkip.received.txt
	// It will be compared against the examples_test.TestExampleVerifyAllCombinationsFor2_withSkip.approved.txt file
	// and contains the text:
	//
	// combineWords
	//
	//
	// [stack,trickle] => stacktrickle
	// [fold,overflow] => foldoverflow
	// [fold,trickle] => foldtrickle
}
