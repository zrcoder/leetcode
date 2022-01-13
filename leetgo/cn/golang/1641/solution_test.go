package solution

import "testing"

func Test_countVowelStrings(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := countVowelStrings(tt.args.n); got != tt.want {
			t.Errorf("%q. countVowelStrings() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
1
*/
