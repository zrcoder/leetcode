package solution

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := longestPalindrome(tt.args.s); got != tt.want {
			t.Errorf("%q. longestPalindrome() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
"babad"
*/
