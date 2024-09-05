package pms_portal_product

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type categoryTreeListRequest struct {
	Locale string `form:"locale,default=zh" binding:"omitempty"` //jacky.xie@2024-08-31
}

type categoryTreeListResponse struct {
	List []dto.PmsProductCategoryNode `json:",inline"`
}

// CategoryTreeList 以树形结构获取所有商品分类
// @Summary 以树形结构获取所有商品分类
// @Description 以树形结构获取所有商品分类
// @Tags PmsPortalProductController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body categoryTreeListRequest true "请求信息"
// @Success 200 {object} code.Success{data=dto.PmsProductCategoryNode}
// @Failure 400 {object} code.Failure
// @Router /product/categoryTreeList [get]
func (h *handler) CategoryTreeList(ctx *gin.Context) {
	req := new(categoryTreeListRequest)
	res := new(categoryTreeListResponse)

	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	list, err := h.service.CategoryTreeList(ctx, req.Locale) //jacky.xie@2024-08-31
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.List = list
	api.Success(ctx, res.List)
}
