package home

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"
)

var _ Service = (*service)(nil)

// 首页内容管理Service
type Service interface {
	i()

	/**
	 * 获取首页内容
	 */
	Content(ctx context.Context, locale string) (*dto.HomeContentResult, error)

	/**
	 * 首页商品推荐
	 */
	RecommendProductList(ctx context.Context, pageNum, pageSize int, locale string) ([]dto.PmsProduct, error) //jacky.xie@2024-08-31

	/**
	 * 获取商品分类
	 * @param parentId 0:获取一级分类；其他：获取指定二级分类
	 */
	GetProductCateList(ctx context.Context, parentId int64, locale string) ([]dto.PmsProductCategory, error) //jacky.xie@2024-08-31

	/**
	 * 根据专题分类分页获取专题
	 * @param cateId 专题分类id
	 */
	GetSubjectList(ctx context.Context, cateId int64, pageNum, pageSize int) ([]dto.CmsSubject, error)

	/**
	 * 分页获取人气推荐商品
	 */
	HotProductList(ctx context.Context, pageNum, pageSize int, locale string) ([]dto.PmsProduct, error) //jacky.xie@2024-08-31

	/**
	 * 分页获取新品推荐商品
	 */
	NewProductList(ctx context.Context, pageNum, pageSize int, locale string) ([]dto.PmsProduct, error) //jacky.xie@2024-08-31
}
