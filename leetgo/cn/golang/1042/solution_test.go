package solution

import (
	"reflect"
	"testing"
)

func Test_gardenNoAdj(t *testing.T) {
	type args struct {
		n     int
		paths [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := gardenNoAdj(tt.args.n, tt.args.paths); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. gardenNoAdj() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
3
[[1,2],[2,3],[3,1]]
*/
