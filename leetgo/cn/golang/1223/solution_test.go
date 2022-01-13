package solution

import "testing"

func Test_dieSimulator(t *testing.T) {
	type args struct {
		n       int
		rollMax []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := dieSimulator(tt.args.n, tt.args.rollMax); got != tt.want {
			t.Errorf("%q. dieSimulator() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
2
[1,1,2,2,2,3]
*/
