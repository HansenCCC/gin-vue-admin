// 自动生成模板Yp_games
package gameList

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// Yp_games 结构体
type Yp_games struct {
      global.GVA_MODEL
      Game_name  string `json:"game_name" form:"game_name" gorm:"column:game_name;comment:;"`
      Game_grade  *float64 `json:"game_grade" form:"game_grade" gorm:"column:game_grade;comment:;"`
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName Yp_games 表名
func (Yp_games) TableName() string {
  return "yp_games"
}

