package solution

import "testing"

func Test_evaluateTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := evaluateTree(tt.args.root); got != tt.want {
			t.Errorf("%q. evaluateTree() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
[2,1,3,null,null,0,1]
*/
