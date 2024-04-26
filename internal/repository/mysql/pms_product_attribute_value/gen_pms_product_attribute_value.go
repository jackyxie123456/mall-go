///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package pms_product_attribute_value

import (
	"fmt"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *PmsProductAttributeValue {
	return new(PmsProductAttributeValue)
}

func NewQueryBuilder() *pmsProductAttributeValueQueryBuilder {
	return new(pmsProductAttributeValueQueryBuilder)
}

func (t *PmsProductAttributeValue) Create(db *gorm.DB) (id int64, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type pmsProductAttributeValueQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *pmsProductAttributeValueQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *pmsProductAttributeValueQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&PmsProductAttributeValue{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

func (qb *pmsProductAttributeValueQueryBuilder) Update(db *gorm.DB, data *PmsProductAttributeValue) (cnt int64, err error) {
	db = db.Model(&PmsProductAttributeValue{})

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

func (qb *pmsProductAttributeValueQueryBuilder) Delete(db *gorm.DB) (err error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if err = db.Delete(&PmsProductAttributeValue{}).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (qb *pmsProductAttributeValueQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&PmsProductAttributeValue{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *pmsProductAttributeValueQueryBuilder) First(db *gorm.DB) (*PmsProductAttributeValue, error) {
	ret := &PmsProductAttributeValue{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *pmsProductAttributeValueQueryBuilder) QueryOne(db *gorm.DB) (*PmsProductAttributeValue, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *pmsProductAttributeValueQueryBuilder) QueryAll(db *gorm.DB) ([]*PmsProductAttributeValue, error) {
	var ret []*PmsProductAttributeValue
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *pmsProductAttributeValueQueryBuilder) Limit(limit int) *pmsProductAttributeValueQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *pmsProductAttributeValueQueryBuilder) Offset(offset int) *pmsProductAttributeValueQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *pmsProductAttributeValueQueryBuilder) WhereId(p mysql.Predicate, value int64) *pmsProductAttributeValueQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeValueQueryBuilder) WhereIdIn(value []int64) *pmsProductAttributeValueQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeValueQueryBuilder) WhereIdNotIn(value []int64) *pmsProductAttributeValueQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeValueQueryBuilder) OrderById(asc bool) *pmsProductAttributeValueQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *pmsProductAttributeValueQueryBuilder) WhereProductId(p mysql.Predicate, value int64) *pmsProductAttributeValueQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_id", p),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeValueQueryBuilder) WhereProductIdIn(value []int64) *pmsProductAttributeValueQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_id", "IN"),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeValueQueryBuilder) WhereProductIdNotIn(value []int64) *pmsProductAttributeValueQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeValueQueryBuilder) OrderByProductId(asc bool) *pmsProductAttributeValueQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "product_id "+order)
	return qb
}

func (qb *pmsProductAttributeValueQueryBuilder) WhereProductAttributeId(p mysql.Predicate, value int64) *pmsProductAttributeValueQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_attribute_id", p),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeValueQueryBuilder) WhereProductAttributeIdIn(value []int64) *pmsProductAttributeValueQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_attribute_id", "IN"),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeValueQueryBuilder) WhereProductAttributeIdNotIn(value []int64) *pmsProductAttributeValueQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_attribute_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeValueQueryBuilder) OrderByProductAttributeId(asc bool) *pmsProductAttributeValueQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "product_attribute_id "+order)
	return qb
}

func (qb *pmsProductAttributeValueQueryBuilder) WhereValue(p mysql.Predicate, value string) *pmsProductAttributeValueQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "value", p),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeValueQueryBuilder) WhereValueIn(value []string) *pmsProductAttributeValueQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "value", "IN"),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeValueQueryBuilder) WhereValueNotIn(value []string) *pmsProductAttributeValueQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "value", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsProductAttributeValueQueryBuilder) OrderByValue(asc bool) *pmsProductAttributeValueQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "value "+order)
	return qb
}
