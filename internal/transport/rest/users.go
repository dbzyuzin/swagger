package rest

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/dbzyuzin/swagger/internal/models"
	"github.com/dbzyuzin/swagger/internal/pkg/tracing"
	"github.com/gin-gonic/gin"
)

func (s *Rest) createUser(ctx *gin.Context) {
	var user models.User
	err := ctx.BindJSON(&user)
	if err != nil {
		s.lg.Error("Invalid body")
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err = s.service.CreateNewUser(ctx, user)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}

func (s *Rest) userExists(ctx *gin.Context) {
	svcCtx := tracing.Add(ctx, ctx.Param("name")+":"+strconv.Itoa(rand.Intn(100)))

	ok, err := s.service.UserExists(svcCtx, ctx.Param("name"))
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, struct {
		Status bool
	}{
		Status: ok,
	})
}
