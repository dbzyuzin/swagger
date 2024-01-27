package rest

import (
	"net/http"

	"github.com/dbzyuzin/swagger/internal/models"
	"github.com/dbzyuzin/swagger/internal/services"
	"github.com/gin-gonic/gin"
)

func createUser(ctx *gin.Context) {
	var user models.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err = services.CreateNewUser(ctx, user)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}

func userExists(ctx *gin.Context) {
	ok, err := services.UserExists(ctx, ctx.Param("name"))
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
