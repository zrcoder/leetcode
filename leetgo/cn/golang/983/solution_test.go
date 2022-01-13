package solution

import "testing"

func Test_mincostTickets(t *testing.T) {
	type args struct {
		days  []int
		costs []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := mincostTickets(tt.args.days, tt.args.costs); got != tt.want {
			t.Errorf("%q. mincostTickets() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
[1,4,6,7,8,20]
[2,7,15]
*/
