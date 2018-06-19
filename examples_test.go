package approvaltests_test

import (
	"github.com/approvals/go-approval-tests"
)

func ExampleVerifyString() {
	approvaltests.VerifyString(t, "Hello World!")
	printFileContent("examples_test.TestExampleVerifyString.received.txt")

	// Output:
	// This produced the file examples_test.TestExampleVerifyString.received.txt
	// It will be compared against the examples_test.TestExampleVerifyString.approved.txt file
	// and contains the text:
	//
	// Hello World!
}

func ExampleVerifyAllCombinationsFor2() {
	letters := []string{"aaaaa", "bbbbb", "ccccc"}
	numbers := []int{2, 3}

	functionToTest := func(text interface{}, length interface{}) string {
		return text.(string)[:length.(int)]
	}

	approvaltests.VerifyAllCombinationsFor2(t, "substring", functionToTest, letters, numbers)
	printFileContent("examples_test.TestExampleVerifyAllCombinationsFor2.received.txt")
	// Output:
	// This produced the file examples_test.TestExampleVerifyAllCombinationsFor2.received.txt
	// It will be compared against the examples_test.TestExampleVerifyAllCombinationsFor2.approved.txt file
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

func ExampleVerifyAllCombinationsFor2WithSkip() {
	words := []string{"stack", "fold"}
	otherWords := []string{"overflow", "trickle"}

	functionToTest := func(firstWord interface{}, secondWord interface{}) string {
		first := firstWord.(string)
		second := secondWord.(string)
		if first+second == "stackoverflow" {
			return approvaltests.SkipThisCombination
		}
		return first + second
	}

	approvaltests.VerifyAllCombinationsFor2(t, "combineWords", functionToTest, words, otherWords)
	printFileContent("examples_test.TestExampleVerifyAllCombinationsFor2WithSkip.received.txt")
	// Output:
	// 	This produced the file examples_test.TestExampleVerifyAllCombinationsFor2WithSkip.received.txt
	// It will be compared against the examples_test.TestExampleVerifyAllCombinationsFor2WithSkip.approved.txt file
	// and contains the text:

	// combineWords

	// [stack,trickle] => stacktrickle
	// [fold,overflow] => foldoverflow
	// [fold,trickle] => foldtrickle
}
