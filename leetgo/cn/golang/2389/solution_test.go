package solution

import (
	"reflect"
	"testing"
)

func Test_answerQueries(t *testing.T) {
	type args struct {
		nums    []int
		queries []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := answerQueries(tt.args.nums, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. answerQueries() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

/* sample test case:
[4,5,2,1]
[3,10,21]
*/
