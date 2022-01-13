package solution

import "testing"

func Test_maximalNetworkRank(t *testing.T) {
	type args struct {
		n     int
		roads [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := maximalNetworkRank(tt.args.n, tt.args.roads); got != tt.want {
			t.Errorf("%q. maximalNetworkRank() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
4
[[0,1],[0,3],[1,2],[1,3]]
*/
