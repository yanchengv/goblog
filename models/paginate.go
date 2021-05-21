package models

import (
	"gorm.io/gorm"
)



//获取分页的数据列表
func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 20
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

//func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
//	return func (db *gorm.DB) *gorm.DB {
//		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
//		if page == 0 {
//			page = 1
//		}
//
//		pageSize, _ := strconv.Atoi(r.URL.Query().Get("page_size"))
//		switch {
//		case pageSize > 100:
//			pageSize = 100
//		case pageSize <= 0:
//			pageSize = 2
//		}
//
//		offset := (page - 1) * pageSize
//		return db.Offset(offset).Limit(pageSize)
//	}
//}
