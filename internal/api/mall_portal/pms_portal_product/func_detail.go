package pms_portal_product

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type detailRequest struct {
	Locale string `form:"locale,default=zh" binding:"omitempty"` //jacky.xie@2024-08-31
}

type detailResponse struct {
	dto.PmsPortalProductDetail `json:",inline"`
}

// Detail 获取前台商品详情
// @Summary 获取前台商品详情
// @Description 获取前台商品详情
// @Tags PmsPortalProductController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body detailRequest true "请求信息"
// @Success 200 {object} code.Success{data=detailResponse}
// @Failure 400 {object} code.Failure
// @Router /product/detail/{id}?locale=en [get]
func (h *handler) Detail(ctx *gin.Context) {
	req := new(detailRequest)
	res := new(detailResponse)
	uri := new(dto.UriID)
	if err := ctx.ShouldBindUri(uri); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}
	//jacky.xie@2024-08-31
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	item, err := h.service.Detail(ctx, uri.Id, req.Locale) //jacky.xie@2024-08-31
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.PmsPortalProductDetail = *item
	api.Success(ctx, res)
}
