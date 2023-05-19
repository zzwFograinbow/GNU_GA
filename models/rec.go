package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

//type Rec struct {
//	ReID          int    `grom:"column:re_id"`
//	UserID        int    `grom:"column:use_id"`
//	Type          string `grom:"column:type"`
//	Param         string `grom:"column:param"`
//	Status        int    `grom:"column:status"`
//	Create_time   string `grom:"column:create_time"`
//	Done_time     string `grom:"column:done_time"`
//	AlgoServer_ip string `grom:"column:algo_server_ip"`
//}

/*
这很可能是由于 GORM 的默认命名约定导致的。在 GORM 中，默认情况下会将驼峰式命名的 Go 结构体字段名转换为下划线分隔的小写字母形式的数据库列名。

例如，如果您有一个名为 ReviewID 的字段，它将被转换为 review_id 的数据库列名。同样地，如果您的结构体名称为 Review，默认情况下 GORM 会使用 reviews 来命名该表。

如果您想使用不同的表名或列名，可以使用 GORM 提供的 table 和 column 标签来指定。例如，如果您想将 re_id 字段映射到 reviews 表中的 id 列，您可以将其声明为
type Review struct {
    ID int `gorm:"column:id"`
    // other fields
}
注意，这里 column 标签指定了该字段对应的数据库列名。

*/

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Rec) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

	return
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Rec{}).Where(maps).Count(&count)

	return
}

func ExistTagByName(name string) bool {
	var rec Rec
	db.Select("re_id").Where("re_id = ?", name).First(&rec)
	if rec.ReID > 0 {
		return true
	}
	return false
}

func AddTag(status int) bool {
	var count int
	var reID int
	db.Model(&Rec{}).Count(&count)
	reID = count + 1
	fmt.Println(reID)
	db.Create(&Rec{
		ReID:   reID,
		Status: status,
	})
	return true
}
func (rec *Rec) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (rec *Rec) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}
func ExistTagByID(id int) bool {
	var rec Rec
	db.Select("id").Where("id = ?", id).First(&rec)
	if rec.ReID > 0 {
		return true
	}

	return false
}

func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Rec{})

	return true
}

func EditTag(id int, data interface{}) bool {
	db.Model(&Rec{}).Where("id = ?", id).Updates(data)

	return true
}
