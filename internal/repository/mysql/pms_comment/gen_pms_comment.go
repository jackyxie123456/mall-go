///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package pms_comment

import (
	"fmt"
	"time"

	"github.com/ChangSZ/mall-go/internal/repository/mysql"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *PmsComment {
	return new(PmsComment)
}

func NewQueryBuilder() *pmsCommentQueryBuilder {
	return new(pmsCommentQueryBuilder)
}

func (t *PmsComment) Create(db *gorm.DB) (id int64, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type pmsCommentQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *pmsCommentQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *pmsCommentQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (int64, error) {
	db = db.Model(&PmsComment{})

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

func (qb *pmsCommentQueryBuilder) Delete(db *gorm.DB) (int64, error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	ret := db.Delete(&PmsComment{})
	err := ret.Error
	if err != nil {
		return 0, errors.Wrap(err, "delete err")
	}
	return ret.RowsAffected, nil
}

func (qb *pmsCommentQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&PmsComment{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *pmsCommentQueryBuilder) First(db *gorm.DB) (*PmsComment, error) {
	ret := &PmsComment{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *pmsCommentQueryBuilder) QueryOne(db *gorm.DB) (*PmsComment, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *pmsCommentQueryBuilder) QueryAll(db *gorm.DB) ([]*PmsComment, error) {
	var ret []*PmsComment
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *pmsCommentQueryBuilder) Limit(limit int) *pmsCommentQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *pmsCommentQueryBuilder) Offset(offset int) *pmsCommentQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *pmsCommentQueryBuilder) WhereId(p mysql.Predicate, value int64) *pmsCommentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *pmsCommentQueryBuilder) WhereIdIn(value []int64) *pmsCommentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *pmsCommentQueryBuilder) WhereIdNotIn(value []int64) *pmsCommentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsCommentQueryBuilder) OrderById(asc bool) *pmsCommentQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *pmsCommentQueryBuilder) WhereProductId(p mysql.Predicate, value int64) *pmsCommentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_id", p),
		value,
	})
	return qb
}

func (qb *pmsCommentQueryBuilder) WhereProductIdIn(value []int64) *pmsCommentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_id", "IN"),
		value,
	})
	return qb
}

func (qb *pmsCommentQueryBuilder) WhereProductIdNotIn(value []int64) *pmsCommentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsCommentQueryBuilder) OrderByProductId(asc bool) *pmsCommentQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "product_id "+order)
	return qb
}

func (qb *pmsCommentQueryBuilder) WhereMemberNickName(p mysql.Predicate, value string) *pmsCommentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "member_nick_name", p),
		value,
	})
	return qb
}

func (qb *pmsCommentQueryBuilder) WhereMemberNickNameIn(value []string) *pmsCommentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "member_nick_name", "IN"),
		value,
	})
	return qb
}

func (qb *pmsCommentQueryBuilder) WhereMemberNickNameNotIn(value []string) *pmsCommentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "member_nick_name", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsCommentQueryBuilder) OrderByMemberNickName(asc bool) *pmsCommentQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "member_nick_name "+order)
	return qb
}

func (qb *pmsCommentQueryBuilder) WhereProductName(p mysql.Predicate, value string) *pmsCommentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_name", p),
		value,
	})
	return qb
}

func (qb *pmsCommentQueryBuilder) WhereProductNameIn(value []string) *pmsCommentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_name", "IN"),
		value,
	})
	return qb
}

func (qb *pmsCommentQueryBuilder) WhereProductNameNotIn(value []string) *pmsCommentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_name", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsCommentQueryBuilder) OrderByProductName(asc bool) *pmsCommentQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "product_name "+order)
	return qb
}

func (qb *pmsCommentQueryBuilder) WhereStar(p mysql.Predicate, value int32) *pmsCommentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "star", p),
		value,
	})
	return qb
}

func (qb *pmsCommentQueryBuilder) WhereStarIn(value []int32) *pmsCommentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "star", "IN"),
		value,
	})
	return qb
}

func (qb *pmsCommentQueryBuilder) WhereStarNotIn(value []int32) *pmsCommentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "star", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsCommentQueryBuilder) OrderByStar(asc bool) *pmsCommentQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "star "+order)
	return qb
}

func (qb *pmsCommentQueryBuilder) WhereMemberIp(p mysql.Predicate, value string) *pmsCommentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "member_ip", p),
		value,
	})
	return qb
}

func (qb *pmsCommentQueryBuilder) WhereMemberIpIn(value []string) *pmsCommentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "member_ip", "IN"),
		value,
	})
	return qb
}

func (qb *pmsCommentQueryBuilder) WhereMemberIpNotIn(value []string) *pmsCommentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "member_ip", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsCommentQueryBuilder) OrderByMemberIp(asc bool) *pmsCommentQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "member_ip "+order)
	return qb
}

func (qb *pmsCommentQueryBuilder) WhereCreateTime(p mysql.Predicate, value time.Time) *pmsCommentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", p),
		value,
	})
	return qb
}

func (qb *pmsCommentQueryBuilder) WhereCreateTimeIn(value []time.Time) *pmsCommentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", "IN"),
		value,
	})
	return qb
}

func (qb *pmsCommentQueryBuilder) WhereCreateTimeNotIn(value []time.Time) *pmsCommentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsCommentQueryBuilder) OrderByCreateTime(asc bool) *pmsCommentQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "create_time "+order)
	return qb
}

func (qb *pmsCommentQueryBuilder) WhereShowStatus(p mysql.Predicate, value int32) *pmsCommentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "show_status", p),
		value,
	})
	return qb
}

func (qb *pmsCommentQueryBuilder) WhereShowStatusIn(value []int32) *pmsCommentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "show_status", "IN"),
		value,
	})
	return qb
}

func (qb *pmsCommentQueryBuilder) WhereShowStatusNotIn(value []int32) *pmsCommentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "show_status", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsCommentQueryBuilder) OrderByShowStatus(asc bool) *pmsCommentQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "show_status "+order)
	return qb
}

func (qb *pmsCommentQueryBuilder) WhereProductAttribute(p mysql.Predicate, value string) *pmsCommentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_attribute", p),
		value,
	})
	return qb
}

func (qb *pmsCommentQueryBuilder) WhereProductAttributeIn(value []string) *pmsCommentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_attribute", "IN"),
		value,
	})
	return qb
}

func (qb *pmsCommentQueryBuilder) WhereProductAttributeNotIn(value []string) *pmsCommentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "product_attribute", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsCommentQueryBuilder) OrderByProductAttribute(asc bool) *pmsCommentQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "product_attribute "+order)
	return qb
}

func (qb *pmsCommentQueryBuilder) WhereCollectCouont(p mysql.Predicate, value int32) *pmsCommentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "collect_couont", p),
		value,
	})
	return qb
}

func (qb *pmsCommentQueryBuilder) WhereCollectCouontIn(value []int32) *pmsCommentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "collect_couont", "IN"),
		value,
	})
	return qb
}

func (qb *pmsCommentQueryBuilder) WhereCollectCouontNotIn(value []int32) *pmsCommentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "collect_couont", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsCommentQueryBuilder) OrderByCollectCouont(asc bool) *pmsCommentQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "collect_couont "+order)
	return qb
}

func (qb *pmsCommentQueryBuilder) WhereReadCount(p mysql.Predicate, value int32) *pmsCommentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "read_count", p),
		value,
	})
	return qb
}

func (qb *pmsCommentQueryBuilder) WhereReadCountIn(value []int32) *pmsCommentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "read_count", "IN"),
		value,
	})
	return qb
}

func (qb *pmsCommentQueryBuilder) WhereReadCountNotIn(value []int32) *pmsCommentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "read_count", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsCommentQueryBuilder) OrderByReadCount(asc bool) *pmsCommentQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "read_count "+order)
	return qb
}

func (qb *pmsCommentQueryBuilder) WhereContent(p mysql.Predicate, value string) *pmsCommentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "content", p),
		value,
	})
	return qb
}

func (qb *pmsCommentQueryBuilder) WhereContentIn(value []string) *pmsCommentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "content", "IN"),
		value,
	})
	return qb
}

func (qb *pmsCommentQueryBuilder) WhereContentNotIn(value []string) *pmsCommentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "content", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsCommentQueryBuilder) OrderByContent(asc bool) *pmsCommentQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "content "+order)
	return qb
}

func (qb *pmsCommentQueryBuilder) WherePics(p mysql.Predicate, value string) *pmsCommentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "pics", p),
		value,
	})
	return qb
}

func (qb *pmsCommentQueryBuilder) WherePicsIn(value []string) *pmsCommentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "pics", "IN"),
		value,
	})
	return qb
}

func (qb *pmsCommentQueryBuilder) WherePicsNotIn(value []string) *pmsCommentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "pics", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsCommentQueryBuilder) OrderByPics(asc bool) *pmsCommentQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "pics "+order)
	return qb
}

func (qb *pmsCommentQueryBuilder) WhereMemberIcon(p mysql.Predicate, value string) *pmsCommentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "member_icon", p),
		value,
	})
	return qb
}

func (qb *pmsCommentQueryBuilder) WhereMemberIconIn(value []string) *pmsCommentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "member_icon", "IN"),
		value,
	})
	return qb
}

func (qb *pmsCommentQueryBuilder) WhereMemberIconNotIn(value []string) *pmsCommentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "member_icon", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsCommentQueryBuilder) OrderByMemberIcon(asc bool) *pmsCommentQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "member_icon "+order)
	return qb
}

func (qb *pmsCommentQueryBuilder) WhereReplayCount(p mysql.Predicate, value int32) *pmsCommentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "replay_count", p),
		value,
	})
	return qb
}

func (qb *pmsCommentQueryBuilder) WhereReplayCountIn(value []int32) *pmsCommentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "replay_count", "IN"),
		value,
	})
	return qb
}

func (qb *pmsCommentQueryBuilder) WhereReplayCountNotIn(value []int32) *pmsCommentQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "replay_count", "NOT IN"),
		value,
	})
	return qb
}

func (qb *pmsCommentQueryBuilder) OrderByReplayCount(asc bool) *pmsCommentQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "replay_count "+order)
	return qb
}
