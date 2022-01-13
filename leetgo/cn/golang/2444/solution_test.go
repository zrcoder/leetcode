package solution

import "testing"

func Test_countSubarrays(t *testing.T) {
	type args struct {
		nums []int
		minK int
		maxK int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := countSubarrays(tt.args.nums, tt.args.minK, tt.args.maxK); got != tt.want {
			t.Errorf("%q. countSubarrays() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
[1,3,5,2,7,5]
1
5
*/
