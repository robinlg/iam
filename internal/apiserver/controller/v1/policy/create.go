package policy

import (
	"github.com/gin-gonic/gin"
	"github.com/robinlg/iam/internal/pkg/code"
	"github.com/robinlg/iam/internal/pkg/middleware"
	"github.com/robinlg/iam/pkg/log"
	v1 "github.com/robinlg/iamapi/apiserver/v1"
	errors "github.com/robinlg/iamerrors"
	"github.com/robinlg/iamlib/pkg/core"
	metav1 "github.com/robinlg/iamlib/pkg/meta/v1"
)

// Create creates a new ladon policy.
// It will convert the policy to string and store it in the storage.
func (p *PolicyController) Create(c *gin.Context) {
	log.L(c).Info("create policy function called.")

	var r v1.Policy
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, "%s", err.Error()), nil)

		return
	}

	if errs := r.Validate(); len(errs) != 0 {
		core.WriteResponse(c, errors.WithCode(code.ErrValidation, "%s", errs.ToAggregate().Error()), nil)

		return
	}

	r.Username = c.GetString(middleware.UsernameKey)

	if err := p.srv.Policies().Create(c, &r, metav1.CreateOptions{}); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, r)
}
