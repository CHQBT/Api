package user

import (
	"milliy/api/auth"
	"milliy/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login godoc
// @Summary login user
// @Description it generates new tokens
// @Tags auth
// @Param userinfo body model.UserLogin true "login and password"
// @Success 200 {object} string "Token"
// @Failure 400 {object} string "Invalid date"
// @Failure 500 {object} string "error while reading from server"
// @Router /v1/auth/login [post]
func (h *newUsers) Login(c *gin.Context) {
	h.Log.Info("Login is working")
	req := model.UserLogin{}

	if err := c.BindJSON(&req); err != nil {
		h.Log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.User.Login(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	token, err := auth.GeneratedRefreshJWTToken(res)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}

	h.Log.Info("login is succesfully ended")
	c.JSON(http.StatusOK, gin.H{
		"Token": token,
	})
}
