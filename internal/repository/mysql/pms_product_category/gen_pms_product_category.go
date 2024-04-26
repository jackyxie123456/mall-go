///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package pms_product_category

import (
	"fmt"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *PmsProductCategory {
	return new(PmsProductCategory)
}

func NewQueryBuilder() *pmsProductCategoryQueryBuilder {
	return new(pmsProductCategoryQueryBuilder)
}

func (t *PmsProductCategory) Create(db *gorm.DB) (id int64, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type pmsProductCategoryQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *pmsProductCategoryQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *pmsProductCategoryQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&PmsProductCategory{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

func (qb *pmsProductCategoryQueryBuilder) Update(db *gorm.DB, data *PmsProductCategory) (cnt int64, err error) {
	db = db.Model(&PmsProductCategory{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	ret := db.Updates(data)
	err = ret.Error
	if err != nil {
		return 0, errors.Wrap(err, "update err")
	}
	return ret.RowsAffected, nil
}

func (qb *pmsProductCategoryQueryBuilder) Delete(db *gorm.DB) (err error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if err = db.Delete(&PmsProductCategory{}).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (qb *pmsProductCategoryQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&PmsProductCategory{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *pmsProductCategoryQueryBuilder) First(db *gorm.DB) (*PmsProductCategory, error) {
	ret := &PmsProductCategory{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *pmsProductCategoryQueryBuilder) QueryOne(db *gorm.DB) (*PmsProductCategory, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *pmsProductCategoryQueryBuilder) QueryAll(db *gorm.DB) ([]*PmsProductCategory, error) {
	var ret []*PmsProductCategory
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *pmsProductCategoryQueryBuilder) Limit(limit int) *pmsProductCategoryQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) Offset(offset int) *pmsProductCategoryQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) WhereId(p mysql.Predicate, value int64) *pmsProductCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) WhereIdIn(value []int64) *pmsProductCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) WhereIdNotIn(value []int64) *pmsProductCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) OrderById(asc bool) *pmsProductCategoryQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) WhereParentId(p mysql.Predicate, value int64) *pmsProductCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "parent_id", p),
		value,
	})
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) WhereParentIdIn(value []int64) *pmsProductCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "parent_id", "IN"),
		value,
	})
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) WhereParentIdNotIn(value []int64) *pmsProductCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "parent_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) OrderByParentId(asc bool) *pmsProductCategoryQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "parent_id "+order)
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) WhereName(p mysql.Predicate, value string) *pmsProductCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", p),
		value,
	})
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) WhereNameIn(value []string) *pmsProductCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", "IN"),
		value,
	})
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) WhereNameNotIn(value []string) *pmsProductCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) OrderByName(asc bool) *pmsProductCategoryQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "name "+order)
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) WhereLevel(p mysql.Predicate, value int32) *pmsProductCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "level", p),
		value,
	})
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) WhereLevelIn(value []int32) *pmsProductCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "level", "IN"),
		value,
	})
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) WhereLevelNotIn(value []int32) *pmsProductCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "level", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) OrderByLevel(asc bool) *pmsProductCategoryQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "level "+order)
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) WhereProductCount(p mysql.Predicate, value int32) *pmsProductCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_count", p),
		value,
	})
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) WhereProductCountIn(value []int32) *pmsProductCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_count", "IN"),
		value,
	})
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) WhereProductCountNotIn(value []int32) *pmsProductCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_count", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) OrderByProductCount(asc bool) *pmsProductCategoryQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "product_count "+order)
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) WhereProductUnit(p mysql.Predicate, value string) *pmsProductCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_unit", p),
		value,
	})
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) WhereProductUnitIn(value []string) *pmsProductCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_unit", "IN"),
		value,
	})
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) WhereProductUnitNotIn(value []string) *pmsProductCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_unit", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) OrderByProductUnit(asc bool) *pmsProductCategoryQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "product_unit "+order)
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) WhereNavStatus(p mysql.Predicate, value int32) *pmsProductCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "nav_status", p),
		value,
	})
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) WhereNavStatusIn(value []int32) *pmsProductCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "nav_status", "IN"),
		value,
	})
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) WhereNavStatusNotIn(value []int32) *pmsProductCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "nav_status", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) OrderByNavStatus(asc bool) *pmsProductCategoryQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "nav_status "+order)
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) WhereShowStatus(p mysql.Predicate, value int32) *pmsProductCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "show_status", p),
		value,
	})
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) WhereShowStatusIn(value []int32) *pmsProductCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "show_status", "IN"),
		value,
	})
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) WhereShowStatusNotIn(value []int32) *pmsProductCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "show_status", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) OrderByShowStatus(asc bool) *pmsProductCategoryQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "show_status "+order)
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) WhereSort(p mysql.Predicate, value int32) *pmsProductCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sort", p),
		value,
	})
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) WhereSortIn(value []int32) *pmsProductCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sort", "IN"),
		value,
	})
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) WhereSortNotIn(value []int32) *pmsProductCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sort", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) OrderBySort(asc bool) *pmsProductCategoryQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "sort "+order)
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) WhereIcon(p mysql.Predicate, value string) *pmsProductCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "icon", p),
		value,
	})
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) WhereIconIn(value []string) *pmsProductCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "icon", "IN"),
		value,
	})
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) WhereIconNotIn(value []string) *pmsProductCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "icon", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) OrderByIcon(asc bool) *pmsProductCategoryQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "icon "+order)
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) WhereKeywords(p mysql.Predicate, value string) *pmsProductCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "keywords", p),
		value,
	})
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) WhereKeywordsIn(value []string) *pmsProductCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "keywords", "IN"),
		value,
	})
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) WhereKeywordsNotIn(value []string) *pmsProductCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "keywords", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) OrderByKeywords(asc bool) *pmsProductCategoryQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "keywords "+order)
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) WhereDescription(p mysql.Predicate, value string) *pmsProductCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "description", p),
		value,
	})
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) WhereDescriptionIn(value []string) *pmsProductCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "description", "IN"),
		value,
	})
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) WhereDescriptionNotIn(value []string) *pmsProductCategoryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "description", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsProductCategoryQueryBuilder) OrderByDescription(asc bool) *pmsProductCategoryQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "description "+order)
	return qb
}
