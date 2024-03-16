package lib

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Recursive function to construct the query based on the fields provided in the struct
func ConstructQuery[T any](db *gorm.DB, searchStruct T) *gorm.DB {
	query := db
	searchValue := reflect.ValueOf(searchStruct)
	searchType := reflect.TypeOf(searchStruct)
	for i := 0; i < searchValue.NumField(); i++ {
		field := searchType.Field(i)
		value := searchValue.Field(i)

		// Check if the field is a struct
		if value.Kind() == reflect.Struct {
			// Recursively construct query for embedded struct
			query = ConstructQuery(query, value.Interface())
		} else {
			// Construct query for regular field
			query = query.Where(fmt.Sprintf("%s = ?", field.Name), value.Interface())
		}
	}
	return query
}

// Paginate the results
func PaginatedResults(c *gin.Context) func(*gorm.DB) *gorm.DB {
	var data Pagination
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	switch {
	// handle insufficient pageSize
	case pageSize >= 100:
		data.PageSize = 100
	case pageSize <= 0:
		data.PageSize = 10
	}
	// ensure pageNum(offset) is larger than 1
	if pageNum <= 0 {
		data.PageNum = 1
	}
	return func(db *gorm.DB) *gorm.DB {
		offset := (data.PageNum - 1) * data.PageSize
		return db.Offset(offset).Limit(data.PageSize)
	}
}
