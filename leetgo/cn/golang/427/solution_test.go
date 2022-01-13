package solution

import (
	"reflect"
	"testing"
)

func Test_construct(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := construct(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. construct() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
[[0,1],[1,0]]
*/
