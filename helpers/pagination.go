package helpers

import (
	"strconv"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Paginate(r *fiber.Ctx) func(db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
	  page, _ := strconv.Atoi(r.Query("page"))
	  if page == 0 {
		page = 1
	  }
  
	  pageSize, _ := strconv.Atoi(r.Query("page_size"))
	  switch {
	  case pageSize > 100:
		pageSize = 100
	  case pageSize <= 0:
		pageSize = 10
	  }
  
	  offset := (page - 1) * pageSize
	  return db.Offset(offset).Limit(pageSize)
	}
  }
