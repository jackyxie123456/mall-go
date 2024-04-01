///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package ums_member_login_log

import (
	"fmt"
	"time"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *UmsMemberLoginLog {
	return new(UmsMemberLoginLog)
}

func NewQueryBuilder() *umsMemberLoginLogQueryBuilder {
	return new(umsMemberLoginLogQueryBuilder)
}

func (t *UmsMemberLoginLog) Create(db *gorm.DB) (id int64, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type umsMemberLoginLogQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *umsMemberLoginLogQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *umsMemberLoginLogQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&UmsMemberLoginLog{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

func (qb *umsMemberLoginLogQueryBuilder) Delete(db *gorm.DB) (err error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if err = db.Delete(&UmsMemberLoginLog{}).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (qb *umsMemberLoginLogQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&UmsMemberLoginLog{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *umsMemberLoginLogQueryBuilder) First(db *gorm.DB) (*UmsMemberLoginLog, error) {
	ret := &UmsMemberLoginLog{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *umsMemberLoginLogQueryBuilder) QueryOne(db *gorm.DB) (*UmsMemberLoginLog, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *umsMemberLoginLogQueryBuilder) QueryAll(db *gorm.DB) ([]*UmsMemberLoginLog, error) {
	var ret []*UmsMemberLoginLog
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *umsMemberLoginLogQueryBuilder) Limit(limit int) *umsMemberLoginLogQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *umsMemberLoginLogQueryBuilder) Offset(offset int) *umsMemberLoginLogQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *umsMemberLoginLogQueryBuilder) WhereId(p mysql.Predicate, value int64) *umsMemberLoginLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *umsMemberLoginLogQueryBuilder) WhereIdIn(value []int64) *umsMemberLoginLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberLoginLogQueryBuilder) WhereIdNotIn(value []int64) *umsMemberLoginLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberLoginLogQueryBuilder) OrderById(asc bool) *umsMemberLoginLogQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *umsMemberLoginLogQueryBuilder) WhereMemberId(p mysql.Predicate, value int64) *umsMemberLoginLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "member_id", p),
		value,
	})
	return qb
}

func (qb *umsMemberLoginLogQueryBuilder) WhereMemberIdIn(value []int64) *umsMemberLoginLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "member_id", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberLoginLogQueryBuilder) WhereMemberIdNotIn(value []int64) *umsMemberLoginLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "member_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberLoginLogQueryBuilder) OrderByMemberId(asc bool) *umsMemberLoginLogQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "member_id "+order)
	return qb
}

func (qb *umsMemberLoginLogQueryBuilder) WhereCreateTime(p mysql.Predicate, value time.Time) *umsMemberLoginLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", p),
		value,
	})
	return qb
}

func (qb *umsMemberLoginLogQueryBuilder) WhereCreateTimeIn(value []time.Time) *umsMemberLoginLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberLoginLogQueryBuilder) WhereCreateTimeNotIn(value []time.Time) *umsMemberLoginLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberLoginLogQueryBuilder) OrderByCreateTime(asc bool) *umsMemberLoginLogQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "create_time "+order)
	return qb
}

func (qb *umsMemberLoginLogQueryBuilder) WhereIp(p mysql.Predicate, value string) *umsMemberLoginLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "ip", p),
		value,
	})
	return qb
}

func (qb *umsMemberLoginLogQueryBuilder) WhereIpIn(value []string) *umsMemberLoginLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "ip", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberLoginLogQueryBuilder) WhereIpNotIn(value []string) *umsMemberLoginLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "ip", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberLoginLogQueryBuilder) OrderByIp(asc bool) *umsMemberLoginLogQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "ip "+order)
	return qb
}

func (qb *umsMemberLoginLogQueryBuilder) WhereCity(p mysql.Predicate, value string) *umsMemberLoginLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "city", p),
		value,
	})
	return qb
}

func (qb *umsMemberLoginLogQueryBuilder) WhereCityIn(value []string) *umsMemberLoginLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "city", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberLoginLogQueryBuilder) WhereCityNotIn(value []string) *umsMemberLoginLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "city", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberLoginLogQueryBuilder) OrderByCity(asc bool) *umsMemberLoginLogQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "city "+order)
	return qb
}

func (qb *umsMemberLoginLogQueryBuilder) WhereLoginType(p mysql.Predicate, value int32) *umsMemberLoginLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "login_type", p),
		value,
	})
	return qb
}

func (qb *umsMemberLoginLogQueryBuilder) WhereLoginTypeIn(value []int32) *umsMemberLoginLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "login_type", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberLoginLogQueryBuilder) WhereLoginTypeNotIn(value []int32) *umsMemberLoginLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "login_type", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberLoginLogQueryBuilder) OrderByLoginType(asc bool) *umsMemberLoginLogQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "login_type "+order)
	return qb
}

func (qb *umsMemberLoginLogQueryBuilder) WhereProvince(p mysql.Predicate, value string) *umsMemberLoginLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "province", p),
		value,
	})
	return qb
}

func (qb *umsMemberLoginLogQueryBuilder) WhereProvinceIn(value []string) *umsMemberLoginLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "province", "IN"),
		value,
	})
	return qb
}

func (qb *umsMemberLoginLogQueryBuilder) WhereProvinceNotIn(value []string) *umsMemberLoginLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "province", "NOT IN"),
		value,
	})
	return qb
}

func (qb *umsMemberLoginLogQueryBuilder) OrderByProvince(asc bool) *umsMemberLoginLogQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "province "+order)
	return qb
}
