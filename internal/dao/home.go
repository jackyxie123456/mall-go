package dao

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"

	"gorm.io/gorm"
)

type HomeDao struct{}

// GetRecommendBrandList 获取推荐品牌
func (t *HomeDao) GetRecommendBrandList(ctx context.Context, tx *gorm.DB, pageNum, pageSize int, locale string) (
	[]dto.PmsBrand, error) {
	if locale == "EN" || locale == "en" { //jacky.xie@2024-08-31
		return t.getRecommendBrandList_en(ctx, tx, pageNum, pageSize)
	} else {
		return t.getRecommendBrandList_zh(ctx, tx, pageNum, pageSize)
	}
}

func (t *HomeDao) getRecommendBrandList_en(ctx context.Context, tx *gorm.DB, pageNum, pageSize int) (
	[]dto.PmsBrand, error) {
	res := make([]dto.PmsBrand, 0)
	offset := (pageNum - 1) * pageSize
	err := tx.Table("pms_brand b").
		Select("b.*").
		//jacky.xie@2024.09.19
		//Select("b.id,b.first_letter,b.sort,  b.factory_status,  b.show_status,  b.product_count,  b.product_comment_count,  b.logo, b.big_pic, b.brand_story,  b.name_en as name, b.brand_story_en as brand_story").
		Joins("LEFT JOIN sms_home_brand_en hb ON hb.brand_id = b.id").
		Where("hb.recommend_status = 1 AND b.show_status = 1").
		Order("hb.sort DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&res).
		Error
	return res, err
}

func (t *HomeDao) getRecommendBrandList_zh(ctx context.Context, tx *gorm.DB, pageNum, pageSize int) (
	[]dto.PmsBrand, error) {
	res := make([]dto.PmsBrand, 0)
	offset := (pageNum - 1) * pageSize
	err := tx.Table("pms_brand b").
		Select("b.*").
		Joins("LEFT JOIN sms_home_brand hb ON hb.brand_id = b.id").
		Where("hb.recommend_status = 1 AND b.show_status = 1").
		Order("hb.sort DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&res).
		Error
	return res, err
}

// GetHotProductList 获取人气推荐
func (t *HomeDao) GetHotProductList(ctx context.Context, tx *gorm.DB, pageNum, pageSize int, locale string) (
	[]dto.PmsProduct, error) {
	if locale == "EN" || locale == "en" { //jacky.xie@2024-08-31
		return t.getHotProductList_en(ctx, tx, pageNum, pageSize)
	} else {
		return t.getHotProductList_zh(ctx, tx, pageNum, pageSize)
	}
}

func (t *HomeDao) getHotProductList_zh(ctx context.Context, tx *gorm.DB, pageNum, pageSize int) (
	[]dto.PmsProduct, error) {
	res := make([]dto.PmsProduct, 0)
	offset := (pageNum - 1) * pageSize
	err := tx.Table("pms_product p").
		Select("p.*").
		Joins("LEFT JOIN sms_home_recommend_product hp ON hp.product_id = p.id").
		Where("hp.recommend_status = 1 AND p.publish_status = 1").
		Order("hp.sort DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&res).
		Error
	return res, err
}

func (t *HomeDao) getHotProductList_en(ctx context.Context, tx *gorm.DB, pageNum, pageSize int) (
	[]dto.PmsProduct, error) {

	res := make([]dto.PmsProduct, 0)
	offset := (pageNum - 1) * pageSize
	err := tx.Table("pms_product p").
		Select("p.*").
		//jacky.xie@2024.09.19
		//Select("p.*, p.name_en as name, p.sub_title_en as sub_title, p.brand_name_en as brand_name,p.description_en as description, p.product_category_name_en as product_category_name").
		Joins("LEFT JOIN sms_home_recommend_product hp ON hp.product_id = p.id").
		Where("hp.recommend_status = 1 AND p.publish_status = 1").
		Order("hp.sort DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&res).
		Error
	return res, err
}

// GetNewProductList 获取新品推荐
func (t *HomeDao) GetNewProductList(ctx context.Context, tx *gorm.DB, pageNum, pageSize int, locale string) (
	[]dto.PmsProduct, error) {
	if locale == "en" || locale == "EN" {
		return t.getNewProductList_en(ctx, tx, pageNum, pageSize)
	} else {
		return t.getNewProductList_zh(ctx, tx, pageNum, pageSize)
	}
}

func (t *HomeDao) getNewProductList_en(ctx context.Context, tx *gorm.DB, pageNum, pageSize int) (
	[]dto.PmsProduct, error) {
	res := make([]dto.PmsProduct, 0)
	offset := (pageNum - 1) * pageSize
	err := tx.Table("pms_product p").
		Select("p.*").
		//jacky.xie@2024.09.19
		//Select("p.*, p.name_en as name, p.sub_title_en as sub_title, p.brand_name_en as brand_name,p.description_en as description, p.product_category_name_en as product_category_name").
		Joins("LEFT JOIN sms_home_new_product hp ON hp.product_id = p.id").
		Where("hp.recommend_status = 1 AND p.publish_status = 1").
		Order("hp.sort DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&res).
		Error
	return res, err
}

func (t *HomeDao) getNewProductList_zh(ctx context.Context, tx *gorm.DB, pageNum, pageSize int) (
	[]dto.PmsProduct, error) {
	res := make([]dto.PmsProduct, 0)
	offset := (pageNum - 1) * pageSize
	err := tx.Table("pms_product p").
		Select("p.*").
		Joins("LEFT JOIN sms_home_new_product hp ON hp.product_id = p.id").
		Where("hp.recommend_status = 1 AND p.publish_status = 1").
		Order("hp.sort DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&res).
		Error
	return res, err
}

// GetFlashProductList 获取秒杀商品
func (t *HomeDao) GetFlashProductList(ctx context.Context, tx *gorm.DB, flashPromotionId, sessionId int64) (
	[]dto.FlashPromotionProduct, error) {
	res := make([]dto.FlashPromotionProduct, 0)
	err := tx.Table("pms_product p").
		Select("pr.flash_promotion_price, pr.flash_promotion_count, pr.flash_promotion_limit, p.*").
		Joins("LEFT JOIN sms_flash_promotion_product_relation pr ON pr.product_id = p.id").
		Where("pr.flash_promotion_id = ? AND pr.flash_promotion_session_id = ?", flashPromotionId, sessionId).
		Find(&res).
		Error
	return res, err
}

// GetRecommendSubjectList 获取推荐专题
func (t *HomeDao) GetRecommendSubjectList(ctx context.Context, tx *gorm.DB, pageNum, pageSize int) (
	[]dto.CmsSubject, error) {
	res := make([]dto.CmsSubject, 0)
	offset := (pageNum - 1) * pageSize
	err := tx.Table("cms_subject s").
		Select("s.*").
		Joins("LEFT JOIN sms_home_recommend_subject hs ON hs.subject_id = s.id").
		Where("hs.recommend_status = 1 AND s.show_status = 1").
		Order("hs.sort DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&res).
		Error
	return res, err
}
