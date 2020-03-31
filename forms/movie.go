package forms

type CreateMovieCommand struct {
	Name   string  `form:"name" binding:"required"`
	Desc   string  `form:"desc" binding:"required"`
	Rating float32 `form:"rating" binding:"required"`
}

type UpdateMovieCommand struct {
	Name   string  `json:"name" binding:"required"`
	Desc   string  `json:"desc" binding:"required"`
	Rating float32 `json:"rating" binding:"required"`
}
