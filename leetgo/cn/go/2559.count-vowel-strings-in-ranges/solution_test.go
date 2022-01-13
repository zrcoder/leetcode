// Code generated by https://github.com/j178/leetgo.

package main

import (
	"testing"

	. "github.com/j178/leetgo/testutils/go"
)

var testcases = `
input:
["aba","bcb","ece","aa","e"]
[[0,2],[1,4],[1,1]]
output:
[2,3,0]

input:
["a","e","i"]
[[0,2],[0,1],[2,2]]
output:
[3,2,1]
`

func Test_vowelStrings(t *testing.T) {
	targetCaseNum := -1
	if err := RunTestsWithString(t, vowelStrings, testcases, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}