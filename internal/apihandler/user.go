package apihandler

import (
	"bank-app-backend/internal/domain"
	"bank-app-backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"

	"errors"
)

// @Summary		Get user info
// @Security  UserBearerAuth
// @Produce		json
// @Success		200		{object}  domain.User
// @Failure   403   {object}	response	"User deleted or banned"
// @Failure		500		{object}	response
// @Router	  /user [get]
func (h *Handler) getUser(c *gin.Context) {
	userPubId, err := getUserPubId(c)
	if err != nil {
		newErrResponse(c, http.StatusInternalServerError, "internal server error", err)
	}

	user, err := h.service.Users.Get(c.Request.Context(), userPubId)
	if err != nil {
		if errors.Is(err, domain.ErrUserDeleted) {
			newResponse(c, http.StatusForbidden, err.Error())
		} else {
			newErrResponse(c, http.StatusInternalServerError, "internal server error", err)
		}
		return
	}

	user.ID = 0
	c.JSON(http.StatusOK, user)
}

type updateUserInput struct {
	Email           string  `json:"email" binding:"omitempty,email,max=64"`
	Password        string  `json:"password" binding:"omitempty,min=8,max=64"`
	Passport        string  `json:"passport" binding:"omitempty,min=8,max=64"`
	Name            string  `json:"name" binding:"omitempty,min=1,max=64"`
	Surname         string  `json:"surname" binding:"omitempty,min=1,max=64"`
	Patronymic      *string `json:"patronymic" binding:"omitempty,max=64"`
	CurrentPassword string  `json:"currentPassword" binding:"required,min=8,max=64"`
}

// @Summary		Update user profile
// @Security  UserBearerAuth
// @Accept	  json
// @Produce		json
// @Param			input body		  updateUserInput	true	"New profile data and current password"
// @Success		200		{object}	response "Successfully updated"
// @Failure		401		{object}	response "Incorrect current password"
// @Failure		400		{object}	response
// @Failure		404		{object}	response
// @Failure		409		{object}	response "User with such email already exists"
// @Failure		500		{object}	response
// @Router		/user [patch]
func (h *Handler) updateUserProfile(c *gin.Context) {
	var input updateUserInput

	if err := c.BindJSON(&input); err != nil {
		newErrResponse(c, http.StatusBadRequest, "invalid input body", err)
		return
	}

	userPubId, err := getUserPubId(c)
	if err != nil {
		newErrResponse(c, http.StatusInternalServerError, "internal server error", err)
	}

	err = h.service.Users.Update(c.Request.Context(), userPubId, input.CurrentPassword, service.UsersSignUpInput{
		Email:      input.Email,
		Password:   input.Password,
		Passport:   input.Passport,
		Name:       input.Name,
		Surname:    input.Surname,
		Patronymic: input.Patronymic,
	})

	if err != nil {
		if errors.Is(err, domain.ErrUserAlreadyExists) {
			newResponse(c, http.StatusConflict, domain.ErrUserAlreadyExists.Error())
			return
		} else if errors.Is(err, domain.ErrInvalidLoginCredentials) {
			newResponse(c, http.StatusUnauthorized, domain.ErrInvalidLoginCredentials.Error())
			return
		}

		newErrResponse(c, http.StatusInternalServerError, "internal server error", err)
		return
	}

	newResponse(c, http.StatusOK, "updated")
}
