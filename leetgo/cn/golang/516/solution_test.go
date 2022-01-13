package solution

import "testing"

func Test_longestPalindromeSubseq(t *testing.T) {
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
		if got := longestPalindromeSubseq(tt.args.s); got != tt.want {
			t.Errorf("%q. longestPalindromeSubseq() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
"bbbab"
*/
