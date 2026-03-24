// Copyright 2025 Robin Liu <robinliu27@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package authorize

import (
	"github.com/gin-gonic/gin"
	"github.com/ory/ladon"
	"github.com/robinlg/iam/internal/authzserver/authorization"
	"github.com/robinlg/iam/internal/authzserver/authorization/authorizer"
	"github.com/robinlg/iam/internal/pkg/code"
	errors "github.com/robinlg/iamerrors"
	"github.com/robinlg/iamlib/pkg/core"
)

// AuthzController create a authorize handler used to handle authorize request.
type AuthzController struct {
	store authorizer.PolicyGetter
}

// NewAuthzController creates a authorize handler.
func NewAuthzController(store authorizer.PolicyGetter) *AuthzController {
	return &AuthzController{
		store: store,
	}
}

// Authorize returns whether a request is allow or deny to access a resource and do some action
// under specified condition.
func (a *AuthzController) Authorize(c *gin.Context) {
	var r ladon.Request
	if err := c.ShouldBind(&r); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, err.Error()), nil)

		return
	}

	auth := authorization.NewAuthorizer(authorizer.NewAuthorization(a.store))
	if r.Context == nil {
		r.Context = ladon.Context{}
	}

	r.Context["username"] = c.GetString("username")
	rsp := auth.Authorize(&r)

	core.WriteResponse(c, nil, rsp)
}
