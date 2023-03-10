// 自动生成模板Yp_gamer_data
package gameList

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// Yp_gamer_data 结构体
type Yp_gamer_data struct {
      global.GVA_MODEL
      Gamer_username  string `json:"gamer_username" form:"gamer_username" gorm:"column:gamer_username;comment:;size:10;"`
      Gamer_move_count  *int `json:"gamer_move_count" form:"gamer_move_count" gorm:"column:gamer_move_count;comment:;"`
      Gameover_time  *time.Time `json:"gameover_time" form:"gameover_time" gorm:"column:gameover_time;comment:;"`
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName Yp_gamer_data 表名
func (Yp_gamer_data) TableName() string {
  return "yp_gamer_data"
}

