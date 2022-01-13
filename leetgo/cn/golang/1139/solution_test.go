package solution

import "testing"

func Test_largest1BorderedSquare(t *testing.T) {
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
		if got := largest1BorderedSquare(tt.args.grid); got != tt.want {
			t.Errorf("%q. largest1BorderedSquare() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
[[1,1,1],[1,0,1],[1,1,1]]
*/
