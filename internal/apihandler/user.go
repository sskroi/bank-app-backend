package apihandler

import (
	"bank-app-backend/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"

	"errors"
)

// @Summary		Get user info
// @Security  UserBearerAuth
// @Produce		json
// @Success		200		{object}	  domain.User
// @Failure   403   {object}	response	"User deleted or banned"
// @Failure		500		{object}	response
// @Router			/user [get]
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

	c.JSON(http.StatusOK, user)
}

// type updateUserProfileInput struct {
// 	signInInput
// 	CurrentPassword string `json:"currentPassword"`
// }

// @Summary		Update user profile
// @Security  UserBearerAuth
// @Accept	  json
// @Produce		json
// @Param			input	body		string	true	"New profile data and current password"
// @Success		200		{object}	response "Successfully updated"
// @Failure		401		{object}	response "Incorrect current password"
// @Failure		400		{object}	response
// @Failure		404		{object}	response
// @Failure		409		{object}	response "User with such email already exists"
// @Failure		500		{object}	response
// @Router			/user/update-profile [post]
func (h *Handler) updateUserProfile(c *gin.Context) {
	userPubId, err := getUserPubId(c)
	if err != nil {
		newErrResponse(c, http.StatusInternalServerError, "internal server error", err)
	}

	// NOT IMPLEMENTED, ONLY FOR TESTS

	newResponse(c, http.StatusOK, userPubId.String())
}
