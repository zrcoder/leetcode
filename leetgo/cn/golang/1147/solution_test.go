package solution

import "testing"

func Test_longestDecomposition(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := longestDecomposition(tt.args.text); got != tt.want {
			t.Errorf("%q. longestDecomposition() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
"ghiabcdefhelloadamhelloabcdefghi"
*/
