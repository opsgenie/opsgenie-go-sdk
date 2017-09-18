package teamv2

type Member struct {
	User User `json:"user"`
	Role string `json:"role"`
}

type User struct {
	ID string `json:"id"`
	Username string `json:"username"`
}
