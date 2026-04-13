package auth

type LoginInput struct {
	Phone    string
	Password string
}

type LoginResponse struct {
	Token string
}
