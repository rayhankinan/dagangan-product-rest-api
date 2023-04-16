package repositories

type SignInUserRequest struct {
	Username string `json:"username" binding:"required,ascii"`
	Password string `json:"password" binding:"required,ascii"`
}
