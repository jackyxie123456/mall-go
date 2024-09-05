package oms_cart_item

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/pkg/log"

	"github.com/gin-gonic/gin"
)

type listRequest struct {
	Locale string
}

type listResponse struct{}

// List 获取当前会员的购物车列表
// @Summary 获取当前会员的购物车列表
// @Description 获取当前会员的购物车列表
// @Tags OmsCartItemController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} code.Success{data=[]dto.OmsCartItem}
// @Failure 400 {object} code.Failure
// @Router /cart/list [get]
func (h *handler) List(ctx *gin.Context) {
	req := new(listRequest)
	_ = new(listResponse)
	// jacky.xie @2024.09.01 未测试
	list, err := h.service.List(ctx, req.Locale)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	api.Success(ctx, list)
}
