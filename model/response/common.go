package response

type LoginResponse struct {
	Token string `json:"token" form:"token"`
	ID    string `json:"_id" form:"_id"`
	Email string `json:"email" form:"email"`
	Name  string `json:"name" form:"name"`
}
