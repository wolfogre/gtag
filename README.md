# gtag

<img src="./assets/gtag.png" width="130" alt="gtag logo">

[![Go Reference](https://pkg.go.dev/badge/github.com/wolfogre/gtag.svg)](https://pkg.go.dev/github.com/wolfogre/gtag)
[![Actions](https://github.com/wolfogre/gtag/actions/workflows/test.yaml/badge.svg)](https://github.com/wolfogre/gtag/actions)
[![Codecov](https://codecov.io/gh/wolfogre/gtag/branch/master/graph/badge.svg)](https://codecov.io/gh/wolfogre/gtag)
[![Go Report Card](https://goreportcard.com/badge/github.com/wolfogre/gtag)](https://goreportcard.com/report/github.com/wolfogre/gtag)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/wolfogre/gtag)](https://github.com/wolfogre/gtag/blob/master/go.mod)
[![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/wolfogre/gtag)](https://github.com/wolfogre/gtag/releases)

Help you to get golang struct's tags elegantly.

## Quick start

### 1. Define your struct

A source file `user.go`:

```go
package tutorial

type User struct {
	Id    int    `bson:"_id"`
	Name  string `bson:"name"`
	Email string `bson:"email"`
}
```

## 2. Run gtag

Run

```bash
go run github.com/wolfogre/gtag/cmd/gtag -types User -tags bson .
```

and you will get file user_tag.go:

```go
// Code generated by gtag. DO NOT EDIT.
// See: https://github.com/wolfogre/gtag

//go:generate go run github.com/wolfogre/gtag/cmd/gtag -types User -tags bson .
package tutorial

import (
	"reflect"
	"strings"
)

var (
	// ...
)

// UserTags indicate tags of type User
type UserTags struct {
	Id    string // `bson:"_id"`
	Name  string // `bson:"name"`
	Email string // `bson:"email"`
}

// Tags return specified tags of User
func (*User) Tags(tag string, convert ...func(string) string) UserTags {
	conv := func(in string) string { return strings.TrimSpace(strings.Split(in, ",")[0]) }
	if len(convert) > 0 {
		conv = convert[0]
	}
	if conv == nil {
		conv = func(in string) string { return in }
	}
	return UserTags{
		Id:    conv(tagOfUserId.Get(tag)),
		Name:  conv(tagOfUserName.Get(tag)),
		Email: conv(tagOfUserEmail.Get(tag)),
	}
}

// TagsBson is alias of Tags("bson")
func (*User) TagsBson() UserTags {
	var v *User
	return v.Tags("bson")
}
```

## 3. Use it

Now you can use the generated code to get tags elegantly:

```go
// update mongo document
obj := User{}
tags := obj.TagsBson()

_, err := collection.UpdateOne(
    ctx,
    bson.M{tags.Id: id},
    bson.M{
        "$set", bson.M{
            tags.Name: name,
            tags.Email: email,
        },
    },
)
```

## Project status

Gtag is beta and is considered feature complete.
