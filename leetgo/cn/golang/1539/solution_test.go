package solution

import "testing"

func Test_findKthPositive(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := findKthPositive(tt.args.arr, tt.args.k); got != tt.want {
			t.Errorf("%q. findKthPositive() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
[2,3,4,7,11]
5
*/
