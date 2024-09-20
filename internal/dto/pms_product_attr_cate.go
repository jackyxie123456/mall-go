package dto

type PmsProductAttributeCategory struct {
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	NameEn         string `json:"nameEn"` //jacky.xie@2024.09.06
	AttributeCount int32  `json:"attributeCount"`
	ParamCount     int32  `json:"paramCount"`
}

type PmsProductAttrCateItem struct {
	PmsProductAttributeCategory `json:",inline"`
	ProductAttributeList        []PmsProductAttribute `json:"productAttributeList"`
}
