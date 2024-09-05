package pms_portal_brand

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dao"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_brand"
	"github.com/ChangSZ/mall-go/internal/repository/mysql/pms_product"
	"github.com/ChangSZ/mall-go/pkg/copy"
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}

func (s *service) RecommendList(ctx context.Context, pageNum, pageSize int, locale string) ([]dto.PmsBrand, error) {
	//jacky.xie@2024-08-31
	return new(dao.HomeDao).GetRecommendBrandList(ctx, mysql.DB().GetDbR().WithContext(ctx), pageNum, pageSize, locale)

}

func (s *service) Detail(ctx context.Context, brandId int64, locale string) (*dto.PmsBrand, error) {
	//jacky.xie@2024-08-31
	if locale == "en" || locale == "EN" {
		qb := pms_brand.NewQueryBuilder().WhereId(mysql.EqualPredicate, brandId)
		item, err := qb.First(mysql.DB().GetDbR().WithContext(ctx))
		if err != nil {
			return nil, err
		}
		res := &dto.PmsBrand{}
		copy.AssignStruct(item, res)
		//locale 处理  jacky.xie@20240901
		res.Name = item.NameEn
		res.BrandStory = item.BrandStoryEn
		return res, nil
	} else {
		qb := pms_brand.NewQueryBuilder().WhereId(mysql.EqualPredicate, brandId)
		item, err := qb.First(mysql.DB().GetDbR().WithContext(ctx))
		if err != nil {
			return nil, err
		}
		res := &dto.PmsBrand{}
		copy.AssignStruct(item, res)
		return res, nil
	}

}

func (s *service) productList_zh(ctx context.Context,
	brandId int64, pageNum int, pageSize int) ([]dto.PmsProduct, int64, error) {

	qb := pms_product.NewQueryBuilder().
		WhereDeleteStatus(mysql.EqualPredicate, 0).
		WherePublishStatus(mysql.EqualPredicate, 1).
		WhereBrandId(mysql.EqualPredicate, brandId)

	count, err := qb.Count(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, 0, err
	}
	offset := (pageNum - 1) * pageSize
	list, err := qb.
		Limit(pageSize).
		Offset(offset).
		QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, 0, err
	}

	listData := make([]dto.PmsProduct, 0, len(list))
	for _, v := range list {
		tmp := dto.PmsProduct{}
		copy.AssignStruct(v, &tmp)
		listData = append(listData, tmp)
	}
	return listData, count, nil
}

// jacky.xie@2024-08-31
func (s *service) productList_en(ctx context.Context,
	brandId int64, pageNum int, pageSize int) ([]dto.PmsProduct, int64, error) {

	qb := pms_product.NewQueryBuilder().
		WhereDeleteStatus(mysql.EqualPredicate, 0).
		WherePublishStatus(mysql.EqualPredicate, 1).
		WhereBrandId(mysql.EqualPredicate, brandId)

	count, err := qb.Count(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, 0, err
	}
	offset := (pageNum - 1) * pageSize
	list, err := qb.
		Limit(pageSize).
		Offset(offset).
		QueryAll(mysql.DB().GetDbR().WithContext(ctx))
	if err != nil {
		return nil, 0, err
	}

	listData := make([]dto.PmsProduct, 0, len(list))
	for _, v := range list {
		tmp := dto.PmsProduct{}
		copy.AssignStruct(v, &tmp)
		//locale 处理  jacky.xie@20240901
		tmp.BrandName = v.BrandNameEn
		tmp.Name = v.NameEn
		tmp.SubTitle = v.SubTitleEn
		tmp.Description = v.DescriptionEn
		tmp.ProductCategoryName = v.ProductCategoryNameEn
		listData = append(listData, tmp)
	}
	return listData, count, nil
}

func (s *service) ProductList(ctx context.Context,
	brandId int64, pageNum int, pageSize int, locale string) ([]dto.PmsProduct, int64, error) {

	if locale == "en" || locale == "EN" {
		return s.productList_en(ctx, brandId, pageNum, pageSize)
	} else {
		return s.productList_zh(ctx, brandId, pageNum, pageSize)
	}

}
