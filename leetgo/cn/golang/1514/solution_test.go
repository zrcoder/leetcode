package solution

import "testing"

func Test_maxProbability(t *testing.T) {
	type args struct {
		n        int
		edges    [][]int
		succProb []float64
		start    int
		end      int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := maxProbability(tt.args.n, tt.args.edges, tt.args.succProb, tt.args.start, tt.args.end); got != tt.want {
			t.Errorf("%q. maxProbability() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
3
[[0,1],[1,2],[0,2]]
[0.5,0.5,0.2]
0
2
*/
