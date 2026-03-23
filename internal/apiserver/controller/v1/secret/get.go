package secret

import (
	"github.com/gin-gonic/gin"
	"github.com/robinlg/iam/internal/pkg/middleware"
	"github.com/robinlg/iam/pkg/log"
	"github.com/robinlg/iamlib/pkg/core"
	metav1 "github.com/robinlg/iamlib/pkg/meta/v1"
)

// Get get an policy by the secret identifier.
func (s *SecretController) Get(c *gin.Context) {
	log.L(c).Info("get secret function called.")

	secret, err := s.srv.Secrets().Get(c, c.GetString(middleware.UsernameKey), c.Param("name"), metav1.GetOptions{})
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, secret)
}
