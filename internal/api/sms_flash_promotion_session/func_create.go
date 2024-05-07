package sms_flash_promotion_session

import (
	"github.com/ChangSZ/mall-go/internal/api"

	"github.com/gin-gonic/gin"
)

type createRequest struct{}

type createResponse struct{}

// Create 添加场次
// @Summary 添加场次
// @Description 添加场次
// @Tags SmsFlashPromotionSessionController
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body createRequest true "请求信息"
// @Success 200 {object} code.Success{data=createResponse}
// @Failure 400 {object} code.Failure
// @Router /flashSession [post]
func (h *handler) Create(ctx *gin.Context) {
	api.Success(ctx, nil)
}
