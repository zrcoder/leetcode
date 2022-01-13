package solution

import "testing"

func Test_countOdds(t *testing.T) {
	type args struct {
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := countOdds(tt.args.low, tt.args.high); got != tt.want {
			t.Errorf("%q. countOdds() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
3
7
*/
