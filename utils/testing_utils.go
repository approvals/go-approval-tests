package utils

import "testing"

func AssertEqual(t *testing.T, expected, actual interface{}, message string) {
	if expected != actual {
		t.Fatalf(message+"\n[expected != actual]\n[%s != %s]", expected, actual)
	}
}
