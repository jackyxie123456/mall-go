package sms_flash_promotion

import "time"

// SmsFlashPromotion 限时购表
//
//go:generate gormgen -structs SmsFlashPromotion -input .
type SmsFlashPromotion struct {
	Id         int64     //
	Title      string    // 秒杀时间段名称
	StartDate  time.Time `gorm:"date"` // 开始日期
	EndDate    time.Time `gorm:"date"` // 结束日期
	Status     int32     // 上下线状态
	CreateTime time.Time `gorm:"autoCreateTime"` // 创建时间
}
