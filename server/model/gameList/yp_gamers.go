// 自动生成模板Yp_gamers
package gameList

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// Yp_gamers 结构体
type Yp_gamers struct {
      global.GVA_MODEL
      Uuid  string `json:"uuid" form:"uuid" gorm:"column:uuid;comment:;size:32;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;"`
      PlayCount  *int `json:"playCount" form:"playCount" gorm:"column:play_count;comment:;"`
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName Yp_gamers 表名
func (Yp_gamers) TableName() string {
  return "yp_gamers"
}

