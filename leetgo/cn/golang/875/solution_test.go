package solution

import "testing"

func Test_minEatingSpeed(t *testing.T) {
	type args struct {
		piles []int
		h     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := minEatingSpeed(tt.args.piles, tt.args.h); got != tt.want {
			t.Errorf("%q. minEatingSpeed() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
[3,6,7,11]
8
*/
