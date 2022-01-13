package solution

import "testing"

func Test_maxDistance(t *testing.T) {
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
		if got := maxDistance(tt.args.grid); got != tt.want {
			t.Errorf("%q. maxDistance() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
[[1,0,1],[0,0,0],[1,0,1]]
*/
