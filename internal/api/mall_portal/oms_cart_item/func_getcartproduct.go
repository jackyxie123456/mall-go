package oms_cart_item

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type getCartProductRequest struct {
	ProductId int64  `uri:"productId" binding:"required"` // 商品ID
	Locale    string `form:"locale" default=zh`           // locale
}

type getCartProductResponse struct {
	dto.CartProduct `json:",inline"`
}

// GetCartProduct 获取购物车中指定商品的规格,用于重选规格
// @Summary 获取购物车中指定商品的规格,用于重选规格
// @Description 获取购物车中指定商品的规格,用于重选规格
// @Tags OmsCartItemController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body getCartProductRequest true "请求信息"
// @Success 200 {object} code.Success{data=getCartProductResponse}
// @Failure 400 {object} code.Failure
// @Router /cart/getProduct/{productId} [get]
func (h *handler) GetCartProduct(ctx *gin.Context) {
	req := new(getCartProductRequest)
	res := new(getCartProductResponse)
	if err := ctx.ShouldBindUri(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	data, err := h.service.GetCartProduct(ctx, req.ProductId, req.Locale) //jacky.xie @2024.09.01 未测试
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.CartProduct = *data
	api.Success(ctx, res)
}
