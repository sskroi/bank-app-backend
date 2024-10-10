package handler

import "github.com/gin-gonic/gin"


type SignUpInput struct {
	Email    string `json:"email" binding:"required,email,max=64"`
	Password string `json:"password" binding:"required,min=8,max=64"`
}

type SignUpResponse struct {
	Id string `json:"id"`
}

//	@Summary		Sign up
//	@Description	register new user
//	@Accept			json
//	@Produce		json
//	@Param			input	body			SignUpInput	true "sign up info"
//	@Success		200		{object}	SignUpResponse
//	@Failure		400		{object}	errResponse
//	@Failure		404		{object}	errResponse
//	@Failure		500		{object}	errResponse
//	@Router			/auth/sign-up [post]
func (h *Handler) SignUp(c *gin.Context) {

}
