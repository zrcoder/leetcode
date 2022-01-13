package solution

import "testing"

func Test_maximalPathQuality(t *testing.T) {
	type args struct {
		values  []int
		edges   [][]int
		maxTime int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := maximalPathQuality(tt.args.values, tt.args.edges, tt.args.maxTime); got != tt.want {
			t.Errorf("%q. maximalPathQuality() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
[0,32,10,43]
[[0,1,10],[1,2,15],[0,3,10]]
49
*/
