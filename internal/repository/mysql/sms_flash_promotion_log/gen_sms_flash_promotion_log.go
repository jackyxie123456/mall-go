///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package sms_flash_promotion_log

import (
	"fmt"
	"time"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *SmsFlashPromotionLog {
	return new(SmsFlashPromotionLog)
}

func NewQueryBuilder() *smsFlashPromotionLogQueryBuilder {
	return new(smsFlashPromotionLogQueryBuilder)
}

func (t *SmsFlashPromotionLog) Create(db *gorm.DB) (id int64, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type smsFlashPromotionLogQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *smsFlashPromotionLogQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *smsFlashPromotionLogQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (int64, error) {
	db = db.Model(&SmsFlashPromotionLog{})

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

func (qb *smsFlashPromotionLogQueryBuilder) Delete(db *gorm.DB) (int64, error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	ret := db.Delete(&SmsFlashPromotionLog{})
	err := ret.Error
	if err != nil {
		return 0, errors.Wrap(err, "delete err")
	}
	return ret.RowsAffected, nil
}

func (qb *smsFlashPromotionLogQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&SmsFlashPromotionLog{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *smsFlashPromotionLogQueryBuilder) First(db *gorm.DB) (*SmsFlashPromotionLog, error) {
	ret := &SmsFlashPromotionLog{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *smsFlashPromotionLogQueryBuilder) QueryOne(db *gorm.DB) (*SmsFlashPromotionLog, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *smsFlashPromotionLogQueryBuilder) QueryAll(db *gorm.DB) ([]*SmsFlashPromotionLog, error) {
	var ret []*SmsFlashPromotionLog
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *smsFlashPromotionLogQueryBuilder) Limit(limit int) *smsFlashPromotionLogQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *smsFlashPromotionLogQueryBuilder) Offset(offset int) *smsFlashPromotionLogQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *smsFlashPromotionLogQueryBuilder) WhereId(p mysql.Predicate, value int64) *smsFlashPromotionLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *smsFlashPromotionLogQueryBuilder) WhereIdIn(value []int64) *smsFlashPromotionLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *smsFlashPromotionLogQueryBuilder) WhereIdNotIn(value []int64) *smsFlashPromotionLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *smsFlashPromotionLogQueryBuilder) OrderById(asc bool) *smsFlashPromotionLogQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *smsFlashPromotionLogQueryBuilder) WhereMemberId(p mysql.Predicate, value int64) *smsFlashPromotionLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "member_id", p),
		value,
	})
	return qb
}

func (qb *smsFlashPromotionLogQueryBuilder) WhereMemberIdIn(value []int64) *smsFlashPromotionLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "member_id", "IN"),
		value,
	})
	return qb
}

func (qb *smsFlashPromotionLogQueryBuilder) WhereMemberIdNotIn(value []int64) *smsFlashPromotionLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "member_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *smsFlashPromotionLogQueryBuilder) OrderByMemberId(asc bool) *smsFlashPromotionLogQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "member_id "+order)
	return qb
}

func (qb *smsFlashPromotionLogQueryBuilder) WhereProductId(p mysql.Predicate, value int64) *smsFlashPromotionLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_id", p),
		value,
	})
	return qb
}

func (qb *smsFlashPromotionLogQueryBuilder) WhereProductIdIn(value []int64) *smsFlashPromotionLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_id", "IN"),
		value,
	})
	return qb
}

func (qb *smsFlashPromotionLogQueryBuilder) WhereProductIdNotIn(value []int64) *smsFlashPromotionLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *smsFlashPromotionLogQueryBuilder) OrderByProductId(asc bool) *smsFlashPromotionLogQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "product_id "+order)
	return qb
}

func (qb *smsFlashPromotionLogQueryBuilder) WhereMemberPhone(p mysql.Predicate, value string) *smsFlashPromotionLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "member_phone", p),
		value,
	})
	return qb
}

func (qb *smsFlashPromotionLogQueryBuilder) WhereMemberPhoneIn(value []string) *smsFlashPromotionLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "member_phone", "IN"),
		value,
	})
	return qb
}

func (qb *smsFlashPromotionLogQueryBuilder) WhereMemberPhoneNotIn(value []string) *smsFlashPromotionLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "member_phone", "NOT IN"),
		value,
	})
	return qb
}

func (qb *smsFlashPromotionLogQueryBuilder) OrderByMemberPhone(asc bool) *smsFlashPromotionLogQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "member_phone "+order)
	return qb
}

func (qb *smsFlashPromotionLogQueryBuilder) WhereProductName(p mysql.Predicate, value string) *smsFlashPromotionLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_name", p),
		value,
	})
	return qb
}

func (qb *smsFlashPromotionLogQueryBuilder) WhereProductNameIn(value []string) *smsFlashPromotionLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_name", "IN"),
		value,
	})
	return qb
}

func (qb *smsFlashPromotionLogQueryBuilder) WhereProductNameNotIn(value []string) *smsFlashPromotionLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_name", "NOT IN"),
		value,
	})
	return qb
}

func (qb *smsFlashPromotionLogQueryBuilder) OrderByProductName(asc bool) *smsFlashPromotionLogQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "product_name "+order)
	return qb
}

func (qb *smsFlashPromotionLogQueryBuilder) WhereSubscribeTime(p mysql.Predicate, value time.Time) *smsFlashPromotionLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "subscribe_time", p),
		value,
	})
	return qb
}

func (qb *smsFlashPromotionLogQueryBuilder) WhereSubscribeTimeIn(value []time.Time) *smsFlashPromotionLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "subscribe_time", "IN"),
		value,
	})
	return qb
}

func (qb *smsFlashPromotionLogQueryBuilder) WhereSubscribeTimeNotIn(value []time.Time) *smsFlashPromotionLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "subscribe_time", "NOT IN"),
		value,
	})
	return qb
}

func (qb *smsFlashPromotionLogQueryBuilder) OrderBySubscribeTime(asc bool) *smsFlashPromotionLogQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "subscribe_time "+order)
	return qb
}

func (qb *smsFlashPromotionLogQueryBuilder) WhereSendTime(p mysql.Predicate, value time.Time) *smsFlashPromotionLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "send_time", p),
		value,
	})
	return qb
}

func (qb *smsFlashPromotionLogQueryBuilder) WhereSendTimeIn(value []time.Time) *smsFlashPromotionLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "send_time", "IN"),
		value,
	})
	return qb
}

func (qb *smsFlashPromotionLogQueryBuilder) WhereSendTimeNotIn(value []time.Time) *smsFlashPromotionLogQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "send_time", "NOT IN"),
		value,
	})
	return qb
}

func (qb *smsFlashPromotionLogQueryBuilder) OrderBySendTime(asc bool) *smsFlashPromotionLogQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "send_time "+order)
	return qb
}
