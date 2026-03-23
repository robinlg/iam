package user

import (
	"github.com/gin-gonic/gin"
	"github.com/robinlg/iam/internal/pkg/code"
	"github.com/robinlg/iam/pkg/log"
	errors "github.com/robinlg/iamerrors"
	"github.com/robinlg/iamlib/pkg/auth"
	"github.com/robinlg/iamlib/pkg/core"
	metav1 "github.com/robinlg/iamlib/pkg/meta/v1"
)

// ChangePasswordRequest defines the ChangePasswordRequest data format.
type ChangePasswordRequest struct {
	// Old password.
	// Required: true
	OldPassword string `json:"oldPassword" binding:"omitempty"`

	// New password.
	// Required: true
	NewPassword string `json:"newPassword" binding:"password"`
}

// ChangePassword change the user's password by the user identifier.
func (u *UserController) ChangePassword(c *gin.Context) {
	log.L(c).Info("change password function called.")

	var r ChangePasswordRequest

	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, "%s", err.Error()), nil)

		return
	}

	user, err := u.srv.Users().Get(c, c.Param("name"), metav1.GetOptions{})
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	if err := user.Compare(r.OldPassword); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrPasswordIncorrect, "%s", err.Error()), nil)

		return
	}

	user.Password, _ = auth.Encrypt(r.NewPassword)
	if err := u.srv.Users().ChangePassword(c, user); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, nil)
}
