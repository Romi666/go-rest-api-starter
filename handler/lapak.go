package handler

import (
	"go-rest-api-starter/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//GetListLapakByPasar is a function to get list lapak by pasar
func (h *Handler) GetListLapakByPasar(c *gin.Context) {
	var (
		response model.Response
		request  model.RequestListLapak
	)

	currentTime := time.Now()
	timeFormat := currentTime.Format("2006-01-02 15:04:05")
	err := c.BindJSON(&request)
	if err != nil {
		response.ResponseCode = "400"
		response.Deskripsi = "Bad Request"
		c.JSON(http.StatusBadRequest, response)
		c.Abort()
		return
	}

	res, err := h.Controller.GetListLapakDipasar(request.PasarID, request.Limit)
	if err != nil {
		response.ResponseCode = "500"
		response.Deskripsi = "Internal Server Error"
		response.Data = map[string]interface{}{
			"Message": err.Error(),
			"Date":    timeFormat,
		}
		c.JSON(http.StatusBadRequest, response)
		c.Abort()
		return
	}

	response.ResponseCode = "00"
	response.Deskripsi = "Success"
	response.Data = res
	c.JSON(http.StatusOK, response)
	return
}

