package handler

import (
	"essy_travel/config"
	"essy_travel/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login godoc
// @ID login
// @Router /login [POST]
// @Summary Get By Id Login
// @Description Get By Id Login
// @Tags Login
// @Accept json
// @Produce json
// @Param object body models.Login true "LoginRequestBody"
// @Success 200 {object} Response{data=models.Login} "LoginBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) Login(c *gin.Context) {

	var login = models.Login{}
	err := c.ShouldBindJSON(&login)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "Error while json decoding"+err.Error())
		return
	}
	resp, err := h.strg.Login().GetByLogin(login)
	if err != nil {
		handleResponse(c, 500, "Login or password invalid: "+err.Error())
		return
	}

	if resp == nil {
		handleResponse(c, http.StatusNoContent, "The data with id is not in the database...")
		return
	}
	resp.API_KEY = config.SecretUser

	handleResponse(c, http.StatusOK, resp)
}
