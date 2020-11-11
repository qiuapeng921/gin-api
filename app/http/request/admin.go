package request

type AdminRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Status   int    `json:"status"`
	Phone    string `json:"phone" form:"phone" binding:"required"`
	RoleId   int    `json:"role_id" form:"role_id" binding:"required"`
}