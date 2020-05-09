// user.go
package tutorial

type User struct {
	Id    int    `bson:"_id"`
	Name  string `bson:"name"`
	Email string `bson:"email"`
}
