// 自动生成模板Yp_game_event
package gameList

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// Yp_game_event 结构体
type Yp_game_event struct {
      global.GVA_MODEL
      Event_name  string `json:"event_name" form:"event_name" gorm:"column:event_name;comment:;"`
      Uuid  string `json:"uuid" form:"uuid" gorm:"column:uuid;comment:;"`
      Properties  string `json:"properties" form:"properties" gorm:"column:properties;comment:;"`
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName Yp_game_event 表名
func (Yp_game_event) TableName() string {
  return "yp_game_event"
}

