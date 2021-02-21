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
  
	  limit, _ := strconv.Atoi(r.Query("limit"))
	  switch {
	  case limit > 100:
		limit = 100
	  case limit <= 0:
		limit = 10
	  }
  
	  offset := (page - 1) * limit
	  return db.Offset(offset).Limit(limit)
	}
  }
