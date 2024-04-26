///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package pms_brand

import (
	"fmt"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *PmsBrand {
	return new(PmsBrand)
}

func NewQueryBuilder() *pmsBrandQueryBuilder {
	return new(pmsBrandQueryBuilder)
}

func (t *PmsBrand) Create(db *gorm.DB) (id int64, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type pmsBrandQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *pmsBrandQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
	ret := db
	for _, where := range qb.where {
		ret = ret.Where(where.prefix, where.value)
	}
	for _, order := range qb.order {
		ret = ret.Order(order)
	}
	if qb.limit != 0 {
		ret = ret.Limit(qb.limit)
	}
	ret = ret.Offset(qb.offset)
	return ret
}

func (qb *pmsBrandQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (int64, error) {
	db = db.Model(&PmsBrand{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	ret := db.Updates(m)
	err := ret.Error
	if err != nil {
		return 0, errors.Wrap(err, "updates err")
	}
	return ret.RowsAffected, nil
}

func (qb *pmsBrandQueryBuilder) Delete(db *gorm.DB) (int64, error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	ret := db.Delete(&PmsBrand{})
	err := ret.Error
	if err != nil {
		return 0, errors.Wrap(err, "delete err")
	}
	return ret.RowsAffected, nil
}

func (qb *pmsBrandQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&PmsBrand{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *pmsBrandQueryBuilder) First(db *gorm.DB) (*PmsBrand, error) {
	ret := &PmsBrand{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *pmsBrandQueryBuilder) QueryOne(db *gorm.DB) (*PmsBrand, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *pmsBrandQueryBuilder) QueryAll(db *gorm.DB) ([]*PmsBrand, error) {
	var ret []*PmsBrand
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *pmsBrandQueryBuilder) Limit(limit int) *pmsBrandQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *pmsBrandQueryBuilder) Offset(offset int) *pmsBrandQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *pmsBrandQueryBuilder) WhereId(p mysql.Predicate, value int64) *pmsBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *pmsBrandQueryBuilder) WhereIdIn(value []int64) *pmsBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *pmsBrandQueryBuilder) WhereIdNotIn(value []int64) *pmsBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsBrandQueryBuilder) OrderById(asc bool) *pmsBrandQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *pmsBrandQueryBuilder) WhereName(p mysql.Predicate, value string) *pmsBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", p),
		value,
	})
	return qb
}

func (qb *pmsBrandQueryBuilder) WhereNameIn(value []string) *pmsBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", "IN"),
		value,
	})
	return qb
}

func (qb *pmsBrandQueryBuilder) WhereNameNotIn(value []string) *pmsBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsBrandQueryBuilder) OrderByName(asc bool) *pmsBrandQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "name "+order)
	return qb
}

func (qb *pmsBrandQueryBuilder) WhereFirstLetter(p mysql.Predicate, value string) *pmsBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "first_letter", p),
		value,
	})
	return qb
}

func (qb *pmsBrandQueryBuilder) WhereFirstLetterIn(value []string) *pmsBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "first_letter", "IN"),
		value,
	})
	return qb
}

func (qb *pmsBrandQueryBuilder) WhereFirstLetterNotIn(value []string) *pmsBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "first_letter", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsBrandQueryBuilder) OrderByFirstLetter(asc bool) *pmsBrandQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "first_letter "+order)
	return qb
}

func (qb *pmsBrandQueryBuilder) WhereSort(p mysql.Predicate, value int32) *pmsBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sort", p),
		value,
	})
	return qb
}

func (qb *pmsBrandQueryBuilder) WhereSortIn(value []int32) *pmsBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sort", "IN"),
		value,
	})
	return qb
}

func (qb *pmsBrandQueryBuilder) WhereSortNotIn(value []int32) *pmsBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sort", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsBrandQueryBuilder) OrderBySort(asc bool) *pmsBrandQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "sort "+order)
	return qb
}

func (qb *pmsBrandQueryBuilder) WhereFactoryStatus(p mysql.Predicate, value int32) *pmsBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "factory_status", p),
		value,
	})
	return qb
}

func (qb *pmsBrandQueryBuilder) WhereFactoryStatusIn(value []int32) *pmsBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "factory_status", "IN"),
		value,
	})
	return qb
}

func (qb *pmsBrandQueryBuilder) WhereFactoryStatusNotIn(value []int32) *pmsBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "factory_status", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsBrandQueryBuilder) OrderByFactoryStatus(asc bool) *pmsBrandQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "factory_status "+order)
	return qb
}

func (qb *pmsBrandQueryBuilder) WhereShowStatus(p mysql.Predicate, value int32) *pmsBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "show_status", p),
		value,
	})
	return qb
}

func (qb *pmsBrandQueryBuilder) WhereShowStatusIn(value []int32) *pmsBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "show_status", "IN"),
		value,
	})
	return qb
}

func (qb *pmsBrandQueryBuilder) WhereShowStatusNotIn(value []int32) *pmsBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "show_status", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsBrandQueryBuilder) OrderByShowStatus(asc bool) *pmsBrandQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "show_status "+order)
	return qb
}

func (qb *pmsBrandQueryBuilder) WhereProductCount(p mysql.Predicate, value int32) *pmsBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_count", p),
		value,
	})
	return qb
}

func (qb *pmsBrandQueryBuilder) WhereProductCountIn(value []int32) *pmsBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_count", "IN"),
		value,
	})
	return qb
}

func (qb *pmsBrandQueryBuilder) WhereProductCountNotIn(value []int32) *pmsBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_count", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsBrandQueryBuilder) OrderByProductCount(asc bool) *pmsBrandQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "product_count "+order)
	return qb
}

func (qb *pmsBrandQueryBuilder) WhereProductCommentCount(p mysql.Predicate, value int32) *pmsBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_comment_count", p),
		value,
	})
	return qb
}

func (qb *pmsBrandQueryBuilder) WhereProductCommentCountIn(value []int32) *pmsBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_comment_count", "IN"),
		value,
	})
	return qb
}

func (qb *pmsBrandQueryBuilder) WhereProductCommentCountNotIn(value []int32) *pmsBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_comment_count", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsBrandQueryBuilder) OrderByProductCommentCount(asc bool) *pmsBrandQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "product_comment_count "+order)
	return qb
}

func (qb *pmsBrandQueryBuilder) WhereLogo(p mysql.Predicate, value string) *pmsBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "logo", p),
		value,
	})
	return qb
}

func (qb *pmsBrandQueryBuilder) WhereLogoIn(value []string) *pmsBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "logo", "IN"),
		value,
	})
	return qb
}

func (qb *pmsBrandQueryBuilder) WhereLogoNotIn(value []string) *pmsBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "logo", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsBrandQueryBuilder) OrderByLogo(asc bool) *pmsBrandQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "logo "+order)
	return qb
}

func (qb *pmsBrandQueryBuilder) WhereBigPic(p mysql.Predicate, value string) *pmsBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "big_pic", p),
		value,
	})
	return qb
}

func (qb *pmsBrandQueryBuilder) WhereBigPicIn(value []string) *pmsBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "big_pic", "IN"),
		value,
	})
	return qb
}

func (qb *pmsBrandQueryBuilder) WhereBigPicNotIn(value []string) *pmsBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "big_pic", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsBrandQueryBuilder) OrderByBigPic(asc bool) *pmsBrandQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "big_pic "+order)
	return qb
}

func (qb *pmsBrandQueryBuilder) WhereBrandStory(p mysql.Predicate, value string) *pmsBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "brand_story", p),
		value,
	})
	return qb
}

func (qb *pmsBrandQueryBuilder) WhereBrandStoryIn(value []string) *pmsBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "brand_story", "IN"),
		value,
	})
	return qb
}

func (qb *pmsBrandQueryBuilder) WhereBrandStoryNotIn(value []string) *pmsBrandQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "brand_story", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsBrandQueryBuilder) OrderByBrandStory(asc bool) *pmsBrandQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "brand_story "+order)
	return qb
}
