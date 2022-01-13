package solution

import "testing"

func Test_btreeGameWinningMove(t *testing.T) {
	type args struct {
		root *TreeNode
		n    int
		x    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := btreeGameWinningMove(tt.args.root, tt.args.n, tt.args.x); got != tt.want {
			t.Errorf("%q. btreeGameWinningMove() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
[1,2,3,4,5,6,7,8,9,10,11]
11
3
*/
