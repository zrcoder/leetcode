package solution

import (
	"reflect"
	"testing"
)

func Test_reverseOddLevels(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := reverseOddLevels(tt.args.root); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. reverseOddLevels() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
[2,3,5,8,13,21,34]
*/
