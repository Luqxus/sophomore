package types

type User struct {
	UID      string
	Username string
	Email    string
	Password string
}

func (u *User) ResponseUser() *ResponseUser {
	return &ResponseUser{
		Username: u.Username,
		Email:    u.Email,
	}
}

type ResponseUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type RegisterReqData struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginReqData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
