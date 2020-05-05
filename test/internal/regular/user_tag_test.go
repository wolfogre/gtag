package regular

import (
	"reflect"
	"testing"
)

func TestUser_Tags(t *testing.T) {
	type fields struct {
		Id    int
		Name  string
		Email string
		age   int
	}
	type args struct {
		tag string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   UserTags
	}{
		{
			name: "regular",
			fields: fields{
				Id:    0,
				Name:  "",
				Email: "",
				age:   0,
			},
			args: args{
				tag: "json",
			},
			want: UserTags{
				Id:    "id",
				Name:  "name",
				Email: "email",
				age:   "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := User{
				Id:    tt.fields.Id,
				Name:  tt.fields.Name,
				Email: tt.fields.Email,
				age:   tt.fields.age,
			}
			if got := us.Tags(tt.args.tag); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tags() = %v, want %v", got, tt.want)
			}
		})
	}
}
