package solution

import (
	"reflect"
	"testing"
)

func Test_checkArithmeticSubarrays(t *testing.T) {
	type args struct {
		nums []int
		l    []int
		r    []int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := checkArithmeticSubarrays(tt.args.nums, tt.args.l, tt.args.r); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. checkArithmeticSubarrays() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
[4,6,5,9,3,7]
[0,0,2]
[2,3,5]
*/
