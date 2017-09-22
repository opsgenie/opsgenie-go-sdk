package userv2

// CreateUserResponse is a response of creating user result.
type CreateUserResponse struct {
	Result string `json:"result"`
	Data   Data   `json:"data"`
	*ResponseMeta
}

// Data contains id and name of created user.
type Data struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
