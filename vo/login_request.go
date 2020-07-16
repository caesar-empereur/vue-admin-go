package vo

type LoginRequest struct {
	username string `json:"username"`
	password string `json:"password"`
}
