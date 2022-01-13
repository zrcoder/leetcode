package solution

import "testing"

func Test_balancedString(t *testing.T) {
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
		if got := balancedString(tt.args.s); got != tt.want {
			t.Errorf("%q. balancedString() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
"QWER"
*/
