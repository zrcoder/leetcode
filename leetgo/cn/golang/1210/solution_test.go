package solution

import "testing"

func Test_minimumMoves(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := minimumMoves(tt.args.grid); got != tt.want {
			t.Errorf("%q. minimumMoves() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
[[0,0,0,0,0,1],[1,1,0,0,1,0],[0,0,0,0,1,1],[0,0,1,0,1,0],[0,1,1,0,0,0],[0,1,1,0,0,0]]
*/
