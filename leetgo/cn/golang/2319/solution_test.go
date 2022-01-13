package solution

import "testing"

func Test_checkXMatrix(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := checkXMatrix(tt.args.grid); got != tt.want {
			t.Errorf("%q. checkXMatrix() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
[[2,0,0,1],[0,3,1,0],[0,5,2,0],[4,0,0,2]]
*/
