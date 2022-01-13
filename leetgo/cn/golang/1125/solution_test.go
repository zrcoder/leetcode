package solution

import (
	"reflect"
	"sort"
	"testing"
)

func Test_smallestSufficientTeam(t *testing.T) {
	type args struct {
		req_skills []string
		people     [][]string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test 1",
			args: args{
				req_skills: []string{"java", "nodejs", "reactjs"},
				people:     [][]string{{"java"}, {"nodejs"}, {"nodejs", "reactjs"}},
			},
			want: []int{0, 2},
		},
		{
			name: "test 2",
			args: args{
				req_skills: []string{"algorithms", "math", "java", "reactjs", "csharp", "aws"},
				people: [][]string{
					{"algorithms", "math", "java"},
					{"algorithms", "math", "reactjs"},
					{"java", "csharp", "aws"},
					{"reactjs", "csharp"},
					{"csharp", "math"},
					{"aws", "java"},
				},
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		got := smallestSufficientTeam(tt.args.req_skills, tt.args.people)
		sort.Ints(got)
		sort.Ints(tt.want)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. smallestSufficientTeam() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
