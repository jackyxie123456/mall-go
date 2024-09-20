package pms_product_attribute_category

// PmsProductAttributeCategory 产品属性分类表
//
//go:generate gormgen -structs PmsProductAttributeCategory -input .
type PmsProductAttributeCategory struct {
	Id             int64  //
	Name           string //
	NameEn         string //  jacky.xie@ 2024.09.06
	AttributeCount int32  // 属性数量
	ParamCount     int32  // 参数数量
}
