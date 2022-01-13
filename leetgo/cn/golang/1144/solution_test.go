package solution

import "testing"

func Test_movesToMakeZigzag(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := movesToMakeZigzag(tt.args.nums); got != tt.want {
			t.Errorf("%q. movesToMakeZigzag() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
[1,2,3]
*/
