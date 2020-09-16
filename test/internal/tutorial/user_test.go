package tutorial

import "testing"

func TestUser_TagsBson(t *testing.T) {
	obj := User{}
	tags := obj.TagsBson()
	t.Logf("%+v", tags)
}
