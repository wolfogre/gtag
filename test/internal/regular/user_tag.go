// Code generated by gtag. DO NOT EDIT.
// See: https://github.com/wolfogre/gtag

//go:generate gtag -types Empty,User,UserName -tags bson,json .
package regular

import (
	"reflect"
	"strings"
)

var (
	valueOfUser = User{}
	typeOfUser  = reflect.TypeOf(valueOfUser)

	_                = valueOfUser.Id
	fieldOfUserId, _ = typeOfUser.FieldByName("Id")
	tagOfUserId      = fieldOfUserId.Tag

	_                  = valueOfUser.Name
	fieldOfUserName, _ = typeOfUser.FieldByName("Name")
	tagOfUserName      = fieldOfUserName.Tag

	_                   = valueOfUser.Email
	fieldOfUserEmail, _ = typeOfUser.FieldByName("Email")
	tagOfUserEmail      = fieldOfUserEmail.Tag

	_                 = valueOfUser.age
	fieldOfUserage, _ = typeOfUser.FieldByName("age")
	tagOfUserage      = fieldOfUserage.Tag
)

// UserTags indicate tags of type User
type UserTags struct {
	Id    string
	Name  string
	Email string
	age   string
}

// Tags return specified tags of User
func (User) Tags(tag string, convert ...func(string) string) UserTags {
	conv := func(in string) string { return strings.TrimSpace(strings.Split(in, ",")[0]) }
	if len(convert) > 0 && convert[0] != nil {
		conv = convert[0]
	}
	_ = conv
	return UserTags{
		Id:    conv(tagOfUserId.Get(tag)),
		Name:  conv(tagOfUserName.Get(tag)),
		Email: conv(tagOfUserEmail.Get(tag)),
		age:   conv(tagOfUserage.Get(tag)),
	}
}

// TagsBson is alias of Tags("bson")
func (v User) TagsBson() UserTags {
	return v.Tags("bson")
}

// TagsJson is alias of Tags("json")
func (v User) TagsJson() UserTags {
	return v.Tags("json")
}

var (
	valueOfUserName = UserName{}
	typeOfUserName  = reflect.TypeOf(valueOfUserName)

	_                       = valueOfUserName.First
	fieldOfUserNameFirst, _ = typeOfUserName.FieldByName("First")
	tagOfUserNameFirst      = fieldOfUserNameFirst.Tag

	_                      = valueOfUserName.Last
	fieldOfUserNameLast, _ = typeOfUserName.FieldByName("Last")
	tagOfUserNameLast      = fieldOfUserNameLast.Tag
)

// UserNameTags indicate tags of type UserName
type UserNameTags struct {
	First string
	Last  string
}

// Tags return specified tags of UserName
func (UserName) Tags(tag string, convert ...func(string) string) UserNameTags {
	conv := func(in string) string { return strings.TrimSpace(strings.Split(in, ",")[0]) }
	if len(convert) > 0 && convert[0] != nil {
		conv = convert[0]
	}
	_ = conv
	return UserNameTags{
		First: conv(tagOfUserNameFirst.Get(tag)),
		Last:  conv(tagOfUserNameLast.Get(tag)),
	}
}

// TagsBson is alias of Tags("bson")
func (v UserName) TagsBson() UserNameTags {
	return v.Tags("bson")
}

// TagsJson is alias of Tags("json")
func (v UserName) TagsJson() UserNameTags {
	return v.Tags("json")
}
