package solution

import (
	"reflect"
	"testing"
)

func Test_restoreMatrix(t *testing.T) {
	type args struct {
		rowSum []int
		colSum []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := restoreMatrix(tt.args.rowSum, tt.args.colSum); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. restoreMatrix() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
[3,8]
[4,7]
*/
