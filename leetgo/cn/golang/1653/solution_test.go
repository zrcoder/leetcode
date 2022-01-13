package solution

import "testing"

func Test_minimumDeletions(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := minimumDeletions(tt.args.s); got != tt.want {
			t.Errorf("%q. minimumDeletions() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
"aababbab"
*/
