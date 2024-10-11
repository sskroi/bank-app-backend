package handler

import (
	"bank-app-backend/internal/domain"
	"bank-app-backend/internal/service"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type signUpInput struct {
	Email      string  `json:"email" binding:"required,email,max=64"`
	Password   string  `json:"password" binding:"required,min=8,max=64"`
	Passport   string  `json:"passport" binding:"required,min=8,max=64"`
	Name       string  `json:"name" binding:"required,min=1,max=64"`
	Surname    string  `json:"surname" binding:"required,min=1,max=64"`
	Patronymic *string `json:"patronymic" binding:"omitempty,max=64"`
}

// @Summary		Sign up
// @Description	register new user
// @Accept			json
// @Produce		json
// @Param			input	body		signUpInput	true	"sign up info"
// @Success		201		{object}	response "user successfully created"
// @Failure		400		{object}	response
// @Failure		404		{object}	response
// @Failure   409   {object}  response "user with such email already exists"
// @Failure		500		{object}	response
// @Router			/auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input signUpInput

	if err := c.BindJSON(&input); err != nil {
		newErrResponse(c, http.StatusBadRequest, err, "invalid input body")
		return
	}

	err := h.service.Users.SignUp(c.Request.Context(), service.UsersSignUpInput{
		Email:      input.Email,
		Password:   input.Password,
		Passport:   input.Passport,
		Name:       input.Name,
		Surname:    input.Surname,
		Patronymic: input.Patronymic,
	})

	if err != nil {
		if errors.Is(err, domain.ErrUserAlreayExists) {
			newResponse(c, http.StatusConflict, domain.ErrUserAlreayExists.Error())
			return
		}

		newErrResponse(c, http.StatusInternalServerError, err, "internal server error")
		return
	}

	newResponse(c, http.StatusCreated, "success")
}
