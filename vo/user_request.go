package vo

type SignUpParams struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required""`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type LoginParams struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required""`
}
