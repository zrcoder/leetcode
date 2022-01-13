// Code generated by https://github.com/j178/leetgo.

package main

import (
	"testing"

	. "github.com/j178/leetgo/testutils/go"
)

var testcases = `
input:
[1,1,3,4]
[4,4,1,1]
2
output:
15

input:
[1,1]
[1,1]
2
output:
2
`

func Test_miceAndCheese(t *testing.T) {
	targetCaseNum := -1
	if err := RunTestsWithString(t, miceAndCheese, testcases, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
