package vo

type LoginResponse struct {
	Success bool `json:"success"`

	Token string `json:"token"`

	Message string `json:"message"`
}
