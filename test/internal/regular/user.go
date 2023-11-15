package regular

type User struct {
	Id    int      `json:"id" json_x:"id_x"`
	Name  UserName `json:"name,omitempty"`
	Email string   `json:"email"`
	age   int
}

type UserName struct {
	First string `json:"first"`
	Last  string `json:"last"`
}
