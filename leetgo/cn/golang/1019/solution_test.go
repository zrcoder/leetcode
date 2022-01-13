package solution

import (
	"reflect"
	"testing"
)

func Test_nextLargerNodes(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := nextLargerNodes(tt.args.head); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. nextLargerNodes() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
[2,1,5]
*/
