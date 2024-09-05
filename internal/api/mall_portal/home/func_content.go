package home

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/log"

	"github.com/ChangSZ/mall-go/pkg/validator"
	"github.com/gin-gonic/gin"
)

type contentRequest struct {
	Locale string `form:"locale,default=zh" binding:"omitempty"` //jacky.xie@2024-08-31
}

type contentResponse struct {
	dto.HomeContentResult `json:",inline"`
}

// Content 首页内容信息展示
// @Summary 首页内容信息展示
// @Description 首页内容信息展示
// @Tags HomeController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body contentRequest true "请求信息"
// @Success 200 {object} code.Success{data=contentResponse}
// @Failure 400 {object} code.Failure
// @Router /home/content [get]
func (h *handler) Content(ctx *gin.Context) {
	req := new(contentRequest)
	res := new(contentResponse)

	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	data, err := h.service.Content(ctx, req.Locale) //jacky.xie@2024-08-31
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.HomeContentResult = *data
	api.Success(ctx, res)
}
