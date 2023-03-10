package gameList

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/gameList"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    gameListReq "github.com/flipped-aurora/gin-vue-admin/server/model/gameList/request"
    "gorm.io/gorm"
)

type Yp_gamesService struct {
}

// CreateYp_games 创建Yp_games记录
// Author [piexlmax](https://github.com/piexlmax)
func (yp_gamesService *Yp_gamesService) CreateYp_games(yp_games gameList.Yp_games) (err error) {
	err = global.GVA_DB.Create(&yp_games).Error
	return err
}

// DeleteYp_games 删除Yp_games记录
// Author [piexlmax](https://github.com/piexlmax)
func (yp_gamesService *Yp_gamesService)DeleteYp_games(yp_games gameList.Yp_games) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&gameList.Yp_games{}).Where("id = ?", yp_games.ID).Update("deleted_by", yp_games.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&yp_games).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteYp_gamesByIds 批量删除Yp_games记录
// Author [piexlmax](https://github.com/piexlmax)
func (yp_gamesService *Yp_gamesService)DeleteYp_gamesByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&gameList.Yp_games{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&gameList.Yp_games{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateYp_games 更新Yp_games记录
// Author [piexlmax](https://github.com/piexlmax)
func (yp_gamesService *Yp_gamesService)UpdateYp_games(yp_games gameList.Yp_games) (err error) {
	err = global.GVA_DB.Save(&yp_games).Error
	return err
}

// GetYp_games 根据id获取Yp_games记录
// Author [piexlmax](https://github.com/piexlmax)
func (yp_gamesService *Yp_gamesService)GetYp_games(id uint) (yp_games gameList.Yp_games, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&yp_games).Error
	return
}

// GetYp_gamesInfoList 分页获取Yp_games记录
// Author [piexlmax](https://github.com/piexlmax)
func (yp_gamesService *Yp_gamesService)GetYp_gamesInfoList(info gameListReq.Yp_gamesSearch) (list []gameList.Yp_games, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&gameList.Yp_games{})
    var yp_gamess []gameList.Yp_games
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["game_grade"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	err = db.Limit(limit).Offset(offset).Find(&yp_gamess).Error
	return  yp_gamess, total, err
}
