package solution

import "testing"

func Test_minimizeArrayValue(t *testing.T) {
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
		if got := minimizeArrayValue(tt.args.nums); got != tt.want {
			t.Errorf("%q. minimizeArrayValue() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
[3,7,1,6]
*/
