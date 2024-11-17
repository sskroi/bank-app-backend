package apihandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// type updateUserProfileInput struct {
// 	signInInput
// 	CurrentPassword string `json:"currentPassword"`
// }

// @Summary		Update user profile
// @Description	Update all user profile info
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
