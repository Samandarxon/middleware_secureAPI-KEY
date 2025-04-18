package handler

import (
	"essy_travel/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateUser godoc
// @ID create_user
// @Router /user [POST]
// @Summary Create User
// @Description Create User
// @Tags User
// @Accept json
// @Produce json
// @Param object body models.CreateUser true "CreateUserRequestBody"
// @Success 200 {object} Response{data=models.User} "UserBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateUser(c *gin.Context) {
	fmt.Println("CreateUser")
	var user = models.CreateUser{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "Error while json decoding"+err.Error())
		return
	}
	resp, err := h.strg.User().Create(user)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, http.StatusCreated, resp)
}

// GetByIdUser godoc
// @ID get_by_id_user
// @Router /user/{id} [GET]
// @Summary Get By Id User
// @Description Get By Id User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=models.User} "UserBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UserGetById(c *gin.Context) {
	id := c.Param("id")
	resp, err := h.strg.User().GetById(models.UserPrimaryKey{Id: id})
	if err != nil {
		handleResponse(c, 500, "User does not exist: "+err.Error())
		return
	}
	if resp == nil {
		handleResponse(c, http.StatusNoContent, "The data with id is not in the database...")
		return
	}

	handleResponse(c, http.StatusOK, resp)
}

// GetListUser godoc
// @ID get_list_user
// @Router /user [GET]
// @Summary Get List User
// @Description Get List User
// @Tags User
// @Accept json
// @Produce json
// @Param limit query number false "limit"
// @Param offset query number false "offset"
// @Success 200 {object} Response{data=models.GetListUserResponse} "GetListUserResponseBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UserGetList(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))

	resp, err := h.strg.User().GetList(models.GetListUserRequest{Offset: offset, Limit: limit})
	if err != nil {
		handleResponse(c, 500, "User does not exist: "+err.Error())
		return
	}

	handleResponse(c, http.StatusOK, resp)
}

// UpdateUser godoc
// @ID update_user
// @Router /user/{id} [PUT]
// @Summary Update User
// @Description Update User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "UserPrimaryKey_ID"
// @Param object body models.UpdateUser true "UpdateUserBody"
// @Success 200 {object} Response{data=string} "Updated User"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UserUpdate(c *gin.Context) {
	var user = models.UpdateUser{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "Error while json decoding"+err.Error())
		return
	}
	id := c.Param("id")
	user.Id = id
	resp, err := h.strg.User().Update(user)
	if err != nil {
		handleResponse(c, 500, "User does not update: "+err.Error())
		return
	}

	handleResponse(c, http.StatusAccepted, resp)
}

// DeleteUser godoc
// @ID delete_user
// @Router /user/{id} [DELETE]
// @Summary Delete User
// @Description Delete User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "DeleteUserPath"
// @Success 200 {object} Response{data=string} "Deleted User"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UserDelete(c *gin.Context) {
	id := c.Param("id")
	_, err := h.strg.User().Delete(models.UserPrimaryKey{Id: id})
	if err != nil {
		handleResponse(c, 500, "User does not delete: "+err.Error())
		return
	}

	handleResponse(c, http.StatusAccepted, "Deleted:")
}
