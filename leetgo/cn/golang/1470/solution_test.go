package solution

import (
	"reflect"
	"testing"
)

func Test_shuffle(t *testing.T) {
	type args struct {
		nums []int
		n    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := shuffle(tt.args.nums, tt.args.n); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. shuffle() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
[2,5,1,3,4,7]
3
*/
