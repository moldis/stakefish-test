package handler

import (
	"code.stakefish.test/service/ip_validator/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type lookUpRequest struct {
	Domain string `form:"domain" json:"domain" xml:"domain" binding:"required" validate:"required,hostname"`
}

func (h *Handler) lookUp(c *gin.Context) {
	var params lookUpRequest
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, &model.HTTPError{
			Message: err.Error(),
		})
		return
	}

	validate := validator.New()
	err := validate.Struct(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, &model.HTTPError{
			Message: "wrong domain",
		})
		return
	}

	ctx := c.Request.Context()
	clientIP := c.ClientIP()
	query, err := h.queries.LookupDomainIPV4(ctx, clientIP, params.Domain)
	if err != nil {
		c.JSON(http.StatusNotFound, &model.HTTPError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, query)
}

type validateIPRequest struct {
	IP string `json:"ip" binding:"required" validate:"required,ipv4"`
}

type validateIPResponse struct {
	Status bool `json:"status" xml:"status"`
}

func (h *Handler) validate(c *gin.Context) {
	var params validateIPRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, &model.HTTPError{
			Message: "wrong params",
		})
		return
	}

	validate := validator.New()
	err := validate.Struct(params)
	if err != nil {
		c.JSON(http.StatusOK, &validateIPResponse{Status: false})
		return
	}

	c.JSON(http.StatusOK, &validateIPResponse{Status: true})
}

type limitOffset struct {
	Skip  int64 `form:"skip,default=0"`
	Limit int64 `form:"limit,default=20"`
}

func (h *Handler) history(c *gin.Context) {
	var params limitOffset
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, &model.HTTPError{
			Message: err.Error(),
		})
		return
	}

	ctx := c.Request.Context()
	res, err := h.queries.History(ctx, params.Skip, params.Limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, &model.HTTPError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}
