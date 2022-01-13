package solution

import "testing"

func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := convert(tt.args.s, tt.args.numRows); got != tt.want {
			t.Errorf("%q. convert() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
"PAYPALISHIRING"
3
*/
