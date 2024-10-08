package oms_portal_order

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type generateConfirmOrderRequest struct {
	CartIds []int64 `json:"cartIds"`
	Locale  string
}

type generateConfirmOrderResponse struct {
	dto.ConfirmOrderResult `json:",inline"`
}

// GenerateConfirmOrder 根据购物车信息生成确认单
// @Summary 根据购物车信息生成确认单
// @Description 根据购物车信息生成确认单
// @Tags OmsPortalOrderController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body []int64 true "请求信息"
// @Success 200 {object} code.Success{data=generateConfirmOrderResponse}
// @Failure 400 {object} code.Failure
// @Router /order/generateConfirmOrder [post]
func (h *handler) GenerateConfirmOrder(ctx *gin.Context) {
	req1 := new(generateConfirmOrderRequest)
	req := make([]int64, 0)
	res := new(generateConfirmOrderResponse)
	if err := ctx.ShouldBind(&req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}
	// jacky.xie@2024.09.01 未测试
	data, err := h.service.GenerateConfirmOrder(ctx, req, req1.Locale)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.ConfirmOrderResult = *data
	api.Success(ctx, res)
}
