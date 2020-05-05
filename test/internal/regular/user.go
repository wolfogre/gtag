package regular

type User struct {
	Id    int      `json:"id"`
	Name  UserName `json:"name"`
	Email string   `json:"email"`
	age   int
}

type UserName struct {
	First string `json:"first"`
	Last  string `json:"last"`
}
