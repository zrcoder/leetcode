package solution

import "testing"

func Test_simplifyPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := simplifyPath(tt.args.path); got != tt.want {
			t.Errorf("%q. simplifyPath() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
"/home/"
*/
