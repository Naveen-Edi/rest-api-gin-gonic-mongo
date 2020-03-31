package forms

type RegisterUserCommand struct {
	Name   string  `form:"name" binding:"required"`
	Email   string  `idx:"{email},unique" form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}



type LoginUserCommand struct {
	Email   string  `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}
