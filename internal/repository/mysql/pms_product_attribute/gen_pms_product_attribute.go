///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package pms_product_attribute

import (
	"fmt"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *PmsProductAttribute {
	return new(PmsProductAttribute)
}

func NewQueryBuilder() *pmsProductAttributeQueryBuilder {
	return new(pmsProductAttributeQueryBuilder)
}

func (t *PmsProductAttribute) Create(db *gorm.DB) (id int64, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type pmsProductAttributeQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *pmsProductAttributeQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *pmsProductAttributeQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (int64, error) {
	db = db.Model(&PmsProductAttribute{})

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

func (qb *pmsProductAttributeQueryBuilder) Delete(db *gorm.DB) (int64, error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	ret := db.Delete(&PmsProductAttribute{})
	err := ret.Error
	if err != nil {
		return 0, errors.Wrap(err, "delete err")
	}
	return ret.RowsAffected, nil
}

func (qb *pmsProductAttributeQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&PmsProductAttribute{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *pmsProductAttributeQueryBuilder) First(db *gorm.DB) (*PmsProductAttribute, error) {
	ret := &PmsProductAttribute{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *pmsProductAttributeQueryBuilder) QueryOne(db *gorm.DB) (*PmsProductAttribute, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *pmsProductAttributeQueryBuilder) QueryAll(db *gorm.DB) ([]*PmsProductAttribute, error) {
	var ret []*PmsProductAttribute
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *pmsProductAttributeQueryBuilder) Limit(limit int) *pmsProductAttributeQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) Offset(offset int) *pmsProductAttributeQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) WhereId(p mysql.Predicate, value int64) *pmsProductAttributeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) WhereIdIn(value []int64) *pmsProductAttributeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) WhereIdNotIn(value []int64) *pmsProductAttributeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) OrderById(asc bool) *pmsProductAttributeQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) WhereProductAttributeCategoryId(p mysql.Predicate, value int64) *pmsProductAttributeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_attribute_category_id", p),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) WhereProductAttributeCategoryIdIn(value []int64) *pmsProductAttributeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_attribute_category_id", "IN"),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) WhereProductAttributeCategoryIdNotIn(value []int64) *pmsProductAttributeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_attribute_category_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) OrderByProductAttributeCategoryId(asc bool) *pmsProductAttributeQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "product_attribute_category_id "+order)
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) WhereName(p mysql.Predicate, value string) *pmsProductAttributeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", p),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) WhereNameIn(value []string) *pmsProductAttributeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", "IN"),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) WhereNameNotIn(value []string) *pmsProductAttributeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) OrderByName(asc bool) *pmsProductAttributeQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "name "+order)
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) WhereSelectType(p mysql.Predicate, value int32) *pmsProductAttributeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "select_type", p),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) WhereSelectTypeIn(value []int32) *pmsProductAttributeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "select_type", "IN"),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) WhereSelectTypeNotIn(value []int32) *pmsProductAttributeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "select_type", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) OrderBySelectType(asc bool) *pmsProductAttributeQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "select_type "+order)
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) WhereInputType(p mysql.Predicate, value int32) *pmsProductAttributeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "input_type", p),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) WhereInputTypeIn(value []int32) *pmsProductAttributeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "input_type", "IN"),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) WhereInputTypeNotIn(value []int32) *pmsProductAttributeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "input_type", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) OrderByInputType(asc bool) *pmsProductAttributeQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "input_type "+order)
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) WhereInputList(p mysql.Predicate, value string) *pmsProductAttributeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "input_list", p),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) WhereInputListIn(value []string) *pmsProductAttributeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "input_list", "IN"),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) WhereInputListNotIn(value []string) *pmsProductAttributeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "input_list", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) OrderByInputList(asc bool) *pmsProductAttributeQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "input_list "+order)
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) WhereSort(p mysql.Predicate, value int32) *pmsProductAttributeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sort", p),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) WhereSortIn(value []int32) *pmsProductAttributeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sort", "IN"),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) WhereSortNotIn(value []int32) *pmsProductAttributeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sort", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) OrderBySort(asc bool) *pmsProductAttributeQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "sort "+order)
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) WhereFilterType(p mysql.Predicate, value int32) *pmsProductAttributeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "filter_type", p),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) WhereFilterTypeIn(value []int32) *pmsProductAttributeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "filter_type", "IN"),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) WhereFilterTypeNotIn(value []int32) *pmsProductAttributeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "filter_type", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) OrderByFilterType(asc bool) *pmsProductAttributeQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "filter_type "+order)
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) WhereSearchType(p mysql.Predicate, value int32) *pmsProductAttributeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "search_type", p),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) WhereSearchTypeIn(value []int32) *pmsProductAttributeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "search_type", "IN"),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) WhereSearchTypeNotIn(value []int32) *pmsProductAttributeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "search_type", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) OrderBySearchType(asc bool) *pmsProductAttributeQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "search_type "+order)
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) WhereRelatedStatus(p mysql.Predicate, value int32) *pmsProductAttributeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "related_status", p),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) WhereRelatedStatusIn(value []int32) *pmsProductAttributeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "related_status", "IN"),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) WhereRelatedStatusNotIn(value []int32) *pmsProductAttributeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "related_status", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) OrderByRelatedStatus(asc bool) *pmsProductAttributeQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "related_status "+order)
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) WhereHandAddStatus(p mysql.Predicate, value int32) *pmsProductAttributeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "hand_add_status", p),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) WhereHandAddStatusIn(value []int32) *pmsProductAttributeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "hand_add_status", "IN"),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) WhereHandAddStatusNotIn(value []int32) *pmsProductAttributeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "hand_add_status", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) OrderByHandAddStatus(asc bool) *pmsProductAttributeQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "hand_add_status "+order)
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) WhereType(p mysql.Predicate, value int32) *pmsProductAttributeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "type", p),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) WhereTypeIn(value []int32) *pmsProductAttributeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "type", "IN"),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) WhereTypeNotIn(value []int32) *pmsProductAttributeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "type", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeQueryBuilder) OrderByType(asc bool) *pmsProductAttributeQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "type "+order)
	return qb
}
