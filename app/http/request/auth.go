package request

type AuthRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Code     string `json:"code" form:"code" binding:"required"`
}