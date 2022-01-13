package solution

import "testing"

func Test_bestHand(t *testing.T) {
	type args struct {
		ranks []int
		suits []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := bestHand(tt.args.ranks, tt.args.suits); got != tt.want {
			t.Errorf("%q. bestHand() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
[13,2,3,1,9]
["a","a","a","a","a"]
*/
