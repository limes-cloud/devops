package model

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

func Today(db *gorm.DB, field string) *gorm.DB {
	cond := fmt.Sprintf("%v>? and %v<?", field, field)
	return db.Where(cond, TodayStartUnix(), TodayEndUnix())
}

// TodayStartUnix 今天开始的秒数
func TodayStartUnix() int64 {
	t := time.Now()
	addTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return addTime.Unix()
}

// TodayEndUnix 今天结束始的秒数
func TodayEndUnix() int64 {
	t := time.Now()
	addTime := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
	return addTime.Unix()
}
