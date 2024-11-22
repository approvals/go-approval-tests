package approvals_test

import (
	approvals "github.com/approvals/go-approval-tests"
)

func ExampleVerifyString() {
	t = makeExamplesRunLikeTests("ExampleVerifyString")

	approvals.VerifyString(t, "Hello World!")

	printFileContent("examples_test.ExampleVerifyString.received.txt")

	// Output:
	// This produced the file examples_test.ExampleVerifyString.received.txt
	// It will be compared against the examples_test.ExampleVerifyString.approved.txt file
	// and contains the text:
	//
	// Hello World!
}

func ExampleVerifyAllCombinationsFor2() {
	t = makeExamplesRunLikeTests("ExampleVerifyAllCombinationsFor2")

	letters := []string{"aaaaa", "bbbbb", "ccccc"}
	numbers := []int{2, 3}

	functionToTest := func(text interface{}, length interface{}) string {
		return text.(string)[:length.(int)]
	}

	approvals.VerifyAllCombinationsFor2(t, "substring", functionToTest, letters, numbers)
	printFileContent("examples_test.ExampleVerifyAllCombinationsFor2.received.txt")
	// Output:
	// This produced the file examples_test.ExampleVerifyAllCombinationsFor2.received.txt
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

func makeExamplesRunLikeTests(name string) *approvals.TestFailable {
	t = approvals.NewTestFailableWithName(name)
	approvals.UseFolder("")
	return t
}

func ExampleVerifyAllCombinationsFor2_withSkip() {
	t = makeExamplesRunLikeTests("ExampleVerifyAllCombinationsFor2_withSkip")

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
	printFileContent("examples_test.ExampleVerifyAllCombinationsFor2_withSkip.received.txt")
	// Output:
	// This produced the file examples_test.ExampleVerifyAllCombinationsFor2_withSkip.received.txt
	// It will be compared against the examples_test.ExampleVerifyAllCombinationsFor2_withSkip.approved.txt file
	// and contains the text:
	//
	// combineWords
	//
	// [stack,trickle] => stacktrickle
	// [fold,overflow] => foldoverflow
	// [fold,trickle] => foldtrickle
}
