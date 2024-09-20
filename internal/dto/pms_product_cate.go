package dto

type PmsProductCategoryParam struct {
	ParentId               int64   `json:"parentId" binding:"omitempty"`             // 父分类的编号
	Name                   string  `json:"name" binding:"required"`                  // 商品分类名称
	NameEn                 string  `json:"nameEn" binding:"required"`                // 商品分类名称 E  //jacky.xie @2024.09.04
	ProductUnit            string  `json:"productUnit"`                              // 分类单位
	ProductUnitEn          string  `json:"productUnitEn"`                            // 分类单位 E  //jacky.xie @2024.09.04
	NavStatus              int     `json:"navStatus" binding:"omitempty,oneof=0 1"`  // 是否在导航栏显示
	ShowStatus             int     `json:"showStatus" binding:"omitempty,oneof=0 1"` // 是否进行显示
	Sort                   int     `json:"sort" binding:"gte=0"`                     // 排序
	Icon                   string  `json:"icon"`                                     // 图标
	Keywords               string  `json:"keywords"`                                 // 关键字
	KeywordsEn             string  `json:"keywordsEn"`                               // 关键字 E  //jacky.xie @2024.09.04
	Description            string  `json:"description"`                              // 描述
	DescriptionEn          string  `json:"descriptionEn"`                            // 描述 E  //jacky.xie @2024.09.04
	ProductAttributeIdList []int64 `json:"productAttributeIdList"`                   // 商品相关筛选属性集合
}

type PmsProductCategory struct {
	Id            int64  `json:"id"`
	ParentId      int64  `json:"parentId"`
	Name          string `json:"name"`
	NameEn        string `json:"nameEn"` //jacky.xie @2024.09.04
	Level         int32  `json:"level"`
	ProductCount  int32  `json:"productCount"`
	ProductUnit   string `json:"productUnit"`
	ProductUnitEn string `json:"productUnitEn"` //jacky.xie @2024.09.04
	NavStatus     int32  `json:"navStatus"`
	ShowStatus    int32  `json:"showStatus"`
	Sort          int32  `json:"sort"`
	Icon          string `json:"icon"`
	Keywords      string `json:"keywords"`
	KeywordsEn    string `json:"keywordsEn"` //jacky.xie @2024.09.04
	Description   string `json:"description"`
	DescriptionEn string `json:"descriptionEn"` //jacky.xie @2024.09.04
}

// PmsProductCategoryWithChildrenItem 包含子级分类的商品分类
type PmsProductCategoryWithChildrenItem struct {
	PmsProductCategory `json:",inline"`
	Children           []PmsProductCategory `json:"children" gorm:"foreignKey:ParentId"` // 子级分类
}

type PmsProductCateUri struct {
	ParentId int64 `uri:"parentId" binding:"omitempty"` // 允许为0
}

// PmsProductCategoryNode 包含子分类的商品分类
type PmsProductCategoryNode struct {
	PmsProductCategory `json:",inline"`
	Children           []PmsProductCategoryNode `json:"children"` // 子分类集合
}
