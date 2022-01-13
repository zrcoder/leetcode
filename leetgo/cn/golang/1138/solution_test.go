package solution

import "testing"

func Test_alphabetBoardPath(t *testing.T) {
	type args struct {
		target string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := alphabetBoardPath(tt.args.target); got != tt.want {
			t.Errorf("%q. alphabetBoardPath() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
"leet"
*/
