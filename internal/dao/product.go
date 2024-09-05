package dao

import (
	"context"

	"github.com/ChangSZ/mall-go/internal/dto"

	"gorm.io/gorm"
)

type ProductDao struct{}

func (t *ProductDao) getCartProduct_en(ctx context.Context, tx *gorm.DB, id int64) (*dto.CartProduct, error) {
	res := &dto.CartProduct{}
	err := tx.
		Preload("ProductAttributeList", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name")
		}).
		Preload("SkuStockList", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, sku_code, price, stock, pic")
		}).
		Table("pms_product p").
		//Select(`p.id, p.name, p.sub_title, p.price, p.pic, p.product_attribute_category_id, p.stock`).
		Select(`p.id, p.name_en as name, p.sub_title_en as sub_title, p.price, p.pic, p.product_attribute_category_id, p.stock`).
		Joins("LEFT JOIN pms_product_attribute pa ON p.product_attribute_category_id = pa.product_attribute_category_id").
		Order("pa.sort DESC").
		Where("p.id = ? AND pa.type = 0", id).Find(res).Error
	return res, err
}

func (t *ProductDao) getCartProduct_zh(ctx context.Context, tx *gorm.DB, id int64) (*dto.CartProduct, error) {
	res := &dto.CartProduct{}
	err := tx.
		Preload("ProductAttributeList", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name")
		}).
		Preload("SkuStockList", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, sku_code, price, stock, pic")
		}).
		Table("pms_product p").
		Select(`p.id, p.name, p.sub_title, p.price, p.pic, p.product_attribute_category_id, p.stock`).
		Joins("LEFT JOIN pms_product_attribute pa ON p.product_attribute_category_id = pa.product_attribute_category_id").
		Order("pa.sort DESC").
		Where("p.id = ? AND pa.type = 0", id).Find(res).Error
	return res, err
}

// GetCartProduct 获取购物车商品信息
func (t *ProductDao) GetCartProduct(ctx context.Context, tx *gorm.DB, id int64, locale string) (
	*dto.CartProduct, error) {
	//locale  todo @jacky.xie@2024.09.01
	if locale == "en" || locale == "EN" {
		return t.getCartProduct_en(ctx, tx, id)
	} else {
		return t.getCartProduct_zh(ctx, tx, id)
	}
}

func (t *ProductDao) GetPromotionProductList(ctx context.Context, tx *gorm.DB, ids []int64, locale string) (
	[]dto.PromotionProduct, error) {
	//locale  todo @jacky.xie@2024.09.01
	res := []dto.PromotionProduct{}
	err := tx.
		Preload("SkuStockList", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, product_id, sku_code, price, promotion_price, stock, lock_stock")
		}).
		Preload("ProductLadderList", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, product_id, count, discount")
		}).
		Preload("ProductFullReductionList", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, product_id, full_price, reduce_price")
		}).
		Table("pms_product p").
		Select(`p.id, p.name, p.promotion_type, p.gift_growth, p.gift_point`).
		Where("p.id IN ?", ids).Find(&res).Error
	return res, err
}
