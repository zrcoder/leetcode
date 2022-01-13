package solution

import "testing"

func Test_minimumRecolors(t *testing.T) {
	type args struct {
		blocks string
		k      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := minimumRecolors(tt.args.blocks, tt.args.k); got != tt.want {
			t.Errorf("%q. minimumRecolors() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
"WBBWWBBWBW"
7
*/
