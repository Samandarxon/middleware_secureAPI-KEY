package handler

import (
	"encoding/json"
	"essy_travel/models"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Upload file godo
// @Router /upload [post]
// @ID file_upload
// @Summary Upload file
// @Description Upload file
// @Tags Upload
// @Accept multipart/form-data
// @Produce json
// @Param table_slug query string true "table_slug"
// @Param file formData file true "this is a test file"
// @Success 200 {object} Response{data=string} "ok"
// @Failure 400 {object} Response{data=string} "We need ID!!"
// @Failure 404 {object} Response{data=string} "Can not find file"
func (h *Handler) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		handleResponse(c, 404, "file not")
		return
	}

	dst, _ := os.Getwd()
	dst += "/" + uuid.New().String() + "_" + file.Filename

	c.SaveUploadedFile(file, dst)
	defer os.Remove(dst)

	f, err := os.Open(dst)
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var data []map[string]interface{}

	body, err := io.ReadAll(f)
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.strg.Upload().Upload(&models.UploadRequest{
		TableSlug: c.Query("table_slug"),
		Data:      data,
	})

	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, http.StatusOK, "successfull")
}
