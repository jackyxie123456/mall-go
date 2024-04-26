///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package ums_role_permission_relation

import (
	"fmt"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *UmsRolePermissionRelation {
	return new(UmsRolePermissionRelation)
}

func NewQueryBuilder() *umsRolePermissionRelationQueryBuilder {
	return new(umsRolePermissionRelationQueryBuilder)
}

func (t *UmsRolePermissionRelation) Create(db *gorm.DB) (id int64, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type umsRolePermissionRelationQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *umsRolePermissionRelationQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *umsRolePermissionRelationQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (int64, error) {
	db = db.Model(&UmsRolePermissionRelation{})

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

func (qb *umsRolePermissionRelationQueryBuilder) Delete(db *gorm.DB) (int64, error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	ret := db.Delete(&UmsRolePermissionRelation{})
	err := ret.Error
	if err != nil {
		return 0, errors.Wrap(err, "delete err")
	}
	return ret.RowsAffected, nil
}

func (qb *umsRolePermissionRelationQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&UmsRolePermissionRelation{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *umsRolePermissionRelationQueryBuilder) First(db *gorm.DB) (*UmsRolePermissionRelation, error) {
	ret := &UmsRolePermissionRelation{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *umsRolePermissionRelationQueryBuilder) QueryOne(db *gorm.DB) (*UmsRolePermissionRelation, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *umsRolePermissionRelationQueryBuilder) QueryAll(db *gorm.DB) ([]*UmsRolePermissionRelation, error) {
	var ret []*UmsRolePermissionRelation
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *umsRolePermissionRelationQueryBuilder) Limit(limit int) *umsRolePermissionRelationQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *umsRolePermissionRelationQueryBuilder) Offset(offset int) *umsRolePermissionRelationQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *umsRolePermissionRelationQueryBuilder) WhereId(p mysql.Predicate, value int64) *umsRolePermissionRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *umsRolePermissionRelationQueryBuilder) WhereIdIn(value []int64) *umsRolePermissionRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *umsRolePermissionRelationQueryBuilder) WhereIdNotIn(value []int64) *umsRolePermissionRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsRolePermissionRelationQueryBuilder) OrderById(asc bool) *umsRolePermissionRelationQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *umsRolePermissionRelationQueryBuilder) WhereRoleId(p mysql.Predicate, value int64) *umsRolePermissionRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "role_id", p),
		value,
	})
	return qb
}

func (qb *umsRolePermissionRelationQueryBuilder) WhereRoleIdIn(value []int64) *umsRolePermissionRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "role_id", "IN"),
		value,
	})
	return qb
}

func (qb *umsRolePermissionRelationQueryBuilder) WhereRoleIdNotIn(value []int64) *umsRolePermissionRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "role_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsRolePermissionRelationQueryBuilder) OrderByRoleId(asc bool) *umsRolePermissionRelationQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "role_id "+order)
	return qb
}

func (qb *umsRolePermissionRelationQueryBuilder) WherePermissionId(p mysql.Predicate, value int64) *umsRolePermissionRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "permission_id", p),
		value,
	})
	return qb
}

func (qb *umsRolePermissionRelationQueryBuilder) WherePermissionIdIn(value []int64) *umsRolePermissionRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "permission_id", "IN"),
		value,
	})
	return qb
}

func (qb *umsRolePermissionRelationQueryBuilder) WherePermissionIdNotIn(value []int64) *umsRolePermissionRelationQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "permission_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsRolePermissionRelationQueryBuilder) OrderByPermissionId(asc bool) *umsRolePermissionRelationQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "permission_id "+order)
	return qb
}
