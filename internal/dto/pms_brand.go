package dto

type PmsBrandParam struct {
	Name                string `json:"name" binding:"required"`
	NameEn              string `json:"nameEn"` //jacky.xie@2024.09.19
	FirstLetter         string `json:"firstLetter"`
	Sort                int32  `json:"sort" binding:"omitempty,gte=0"`
	FactoryStatus       int32  `json:"factoryStatus" binding:"omitempty,oneof=0 1"`
	ShowStatus          int32  `json:"showStatus"  binding:"omitempty,oneof=0 1"`
	ProductCount        int32  `json:"productCount"`
	ProductCommentCount int32  `json:"productCommentCount"`
	Logo                string `json:"logo" binding:"required"`
	BigPic              string `json:"bigPic"`
	BrandStory          string `json:"brandStory"`
	BrandStoryEn        string `json:"brandStoryEn"` //jacky.xie@2024.09.19
}

type PmsBrand struct {
	Id                  int64  `json:"id"`
	Name                string `json:"name"`
	NameEn              string `json:"nameEn"` //jacky.xie@2024.09.19
	FirstLetter         string `json:"firstLetter"`
	Sort                int32  `json:"sort"`
	FactoryStatus       int32  `json:"factoryStatus"`
	ShowStatus          int32  `json:"showStatus"`
	ProductCount        int32  `json:"productCount"`
	ProductCommentCount int32  `json:"productCommentCount"`
	Logo                string `json:"logo"`
	BigPic              string `json:"bigPic"`
	BrandStory          string `json:"brandStory"`
	BrandStoryEn        string `json:"brandStoryEn"` //jacky.xie@2024.09.19
}
