package solution

import (
	"reflect"
	"testing"
)

func Test_findContinuousSequence(t *testing.T) {
	type args struct {
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := findContinuousSequence(tt.args.target); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. findContinuousSequence() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
9
*/
