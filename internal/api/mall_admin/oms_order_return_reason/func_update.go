package oms_order_return_reason

import (
	"github.com/ChangSZ/mall-go/internal/api"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/pkg/log"
	"github.com/ChangSZ/mall-go/pkg/validator"

	"github.com/gin-gonic/gin"
)

type updateRequest struct {
	dto.OmsOrderReturnReason `json:",inline"`
}

type updateResponse struct {
	Count int64 `json:",inline"`
}

// Update 修改退货原因
// @Summary 修改退货原因
// @Description 修改退货原因
// @Tags OmsOrderReturnReasonController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateRequest true "请求信息"
// @Success 200 {object} code.Success{data=int64}
// @Failure 400 {object} code.Failure
// @Router /returnReason/update/{id} [post]
func (h *handler) Update(ctx *gin.Context) {
	req := new(updateRequest)
	res := new(updateResponse)
	if err := ctx.ShouldBind(req); err != nil {
		log.WithTrace(ctx).Error(err)
		api.ValidateFailed(ctx, validator.GetValidationError(err).Error())
		return
	}

	cnt, err := h.service.Update(ctx, req.Id, req.OmsOrderReturnReason)
	if err != nil {
		log.WithTrace(ctx).Error(err)
		api.Failed(ctx, err.Error())
		return
	}
	res.Count = cnt
	api.Success(ctx, res.Count)
}
