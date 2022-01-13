package solution

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := isValid(tt.args.s); got != tt.want {
			t.Errorf("%q. isValid() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
"()"
*/
