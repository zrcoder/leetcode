package solution

import "testing"

func Test_bestTeamScore(t *testing.T) {
	type args struct {
		scores []int
		ages   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := bestTeamScore(tt.args.scores, tt.args.ages); got != tt.want {
			t.Errorf("%q. bestTeamScore() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
[1,3,5,10,15]
[1,2,3,4,5]
*/
