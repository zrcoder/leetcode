package solution

import "testing"

func Test_maxValueOfCoins(t *testing.T) {
	type args struct {
		piles [][]int
		k     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := maxValueOfCoins(tt.args.piles, tt.args.k); got != tt.want {
			t.Errorf("%q. maxValueOfCoins() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
[[1,100,3],[7,8,9]]
2
*/
