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

// Update updates policy by the policy identifier.
func (p *PolicyController) Update(c *gin.Context) {
	log.L(c).Info("update policy function called.")

	var r v1.Policy
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, "%s", err.Error()), nil)

		return
	}

	pol, err := p.srv.Policies().Get(c, c.GetString(middleware.UsernameKey), c.Param("name"), metav1.GetOptions{})
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	// only update policy string
	pol.Policy = r.Policy
	pol.Extend = r.Extend

	if errs := pol.Validate(); len(errs) != 0 {
		core.WriteResponse(c, errors.WithCode(code.ErrValidation, "%s", errs.ToAggregate().Error()), nil)

		return
	}

	if err := p.srv.Policies().Update(c, pol, metav1.UpdateOptions{}); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, pol)
}
