package solution

import (
	"reflect"
	"testing"
)

func Test_findDuplicateSubtrees(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []*TreeNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := findDuplicateSubtrees(tt.args.root); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. findDuplicateSubtrees() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
[1,2,3,4,null,2,4,null,null,4]
*/
