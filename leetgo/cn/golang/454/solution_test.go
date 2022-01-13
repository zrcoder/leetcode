package solution

import "testing"

func Test_fourSumCount(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		nums3 []int
		nums4 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := fourSumCount(tt.args.nums1, tt.args.nums2, tt.args.nums3, tt.args.nums4); got != tt.want {
			t.Errorf("%q. fourSumCount() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
[1,2]
[-2,-1]
[-1,2]
[0,2]
*/
