package solution

import "testing"

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := removeElement(tt.args.nums, tt.args.val); got != tt.want {
			t.Errorf("%q. removeElement() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
[3,2,2,3]
3
*/
