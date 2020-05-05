package regular

import (
	"reflect"
	"testing"
)

func TestEmpty_Tags(t *testing.T) {
	type args struct {
		tag string
	}
	tests := []struct {
		name string
		args args
		want EmptyTags
	}{
		{
			name: "regular",
			args: args{
				tag: "json",
			},
			want: EmptyTags{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			em := Empty{}
			if got := em.Tags(tt.args.tag); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tags() = %v, want %v", got, tt.want)
			}
		})
	}
}
