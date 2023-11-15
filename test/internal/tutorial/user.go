package tutorial

//go:generate go run github.com/wolfogre/gtag/cmd/gtag -types User -tags bson .
type User struct {
	Id    int    `bson:"_id"`
	Name  string `bson:"name"`
	Email string `bson:"email"`
}
