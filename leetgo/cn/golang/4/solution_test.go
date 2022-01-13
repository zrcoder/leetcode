package solution

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
			t.Errorf("%q. findMedianSortedArrays() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
[1,3]
[2]
*/
