package regular

import (
	"reflect"
	"strings"
	"testing"
)

func TestEmpty_Tags(t *testing.T) {
	type args struct {
		tag     string
		convert []func(string) string
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
		{
			name: "convert ToUpper",
			args: args{
				tag:     "json",
				convert: []func(string) string{strings.ToUpper},
			},
			want: EmptyTags{},
		},
		{
			name: "convert nil",
			args: args{
				tag:     "json",
				convert: []func(string) string{nil},
			},
			want: EmptyTags{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			em := Empty{}
			if got := em.Tags(tt.args.tag, tt.args.convert...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmpty_TagsJson(t *testing.T) {
	tests := []struct {
		name string
		want EmptyTags
	}{
		{
			name: "regular",
			want: EmptyTags{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Empty{}
			if got := v.TagsJson(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TagsJson() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmpty_TagsBson(t *testing.T) {
	tests := []struct {
		name string
		want EmptyTags
	}{
		{
			name: "regular",
			want: EmptyTags{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Empty{}
			if got := v.TagsBson(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TagsBson() = %v, want %v", got, tt.want)
			}
		})
	}
}
