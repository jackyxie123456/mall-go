package pms_portal_brand

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
)

var _ Service = (*service)(nil)

// 前台品牌管理Service
type Service interface {
	i()

	/**
	 * 分页获取推荐品牌
	 */
	//jacky.xie@2024-08-31
	RecommendList(ctx context.Context, pageNum, pageSize int, locale string) ([]dto.PmsBrand, error)

	/**
	 * 获取品牌详情
	 */
	//jacky.xie@2024-08-31
	Detail(ctx context.Context, brandId int64, locale string) (*dto.PmsBrand, error)

	/**
	 * 分页获取品牌关联商品
	 */
	//jacky.xie@2024-08-31
	ProductList(ctx context.Context, brandId int64, pageNum int, pageSize int, locale string) ([]dto.PmsProduct, int64, error)
}
