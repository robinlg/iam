package secret

import (
	"github.com/gin-gonic/gin"
	"github.com/robinlg/iam/internal/pkg/code"
	"github.com/robinlg/iam/internal/pkg/middleware"
	"github.com/robinlg/iam/pkg/log"
	errors "github.com/robinlg/iamerrors"
	"github.com/robinlg/iamlib/pkg/core"
	metav1 "github.com/robinlg/iamlib/pkg/meta/v1"
)

// List list all the secrets.
func (s *SecretController) List(c *gin.Context) {
	log.L(c).Info("list secret function called.")
	var r metav1.ListOptions
	if err := c.ShouldBindQuery(&r); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, "%s", err.Error()), nil)

		return
	}

	secrets, err := s.srv.Secrets().List(c, c.GetString(middleware.UsernameKey), r)
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, secrets)
}
