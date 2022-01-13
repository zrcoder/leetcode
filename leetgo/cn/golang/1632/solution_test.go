package solution

import (
	"reflect"
	"testing"
)

func Test_matrixRankTransform(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := matrixRankTransform(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. matrixRankTransform() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
[[1,2],[3,4]]
*/
