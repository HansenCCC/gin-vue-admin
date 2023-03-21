package gameList

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/gameList"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    gameListReq "github.com/flipped-aurora/gin-vue-admin/server/model/gameList/request"
    "gorm.io/gorm"
)

type Yp_gamersService struct {
}

// CreateYp_gamers 创建Yp_gamers记录
// Author [piexlmax](https://github.com/piexlmax)
func (yp_gamersService *Yp_gamersService) CreateYp_gamers(yp_gamers gameList.Yp_gamers) (err error) {
	err = global.GVA_DB.Create(&yp_gamers).Error
	return err
}

// DeleteYp_gamers 删除Yp_gamers记录
// Author [piexlmax](https://github.com/piexlmax)
func (yp_gamersService *Yp_gamersService)DeleteYp_gamers(yp_gamers gameList.Yp_gamers) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&gameList.Yp_gamers{}).Where("id = ?", yp_gamers.ID).Update("deleted_by", yp_gamers.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&yp_gamers).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteYp_gamersByIds 批量删除Yp_gamers记录
// Author [piexlmax](https://github.com/piexlmax)
func (yp_gamersService *Yp_gamersService)DeleteYp_gamersByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&gameList.Yp_gamers{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&gameList.Yp_gamers{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateYp_gamers 更新Yp_gamers记录
// Author [piexlmax](https://github.com/piexlmax)
func (yp_gamersService *Yp_gamersService)UpdateYp_gamers(yp_gamers gameList.Yp_gamers) (err error) {
	err = global.GVA_DB.Save(&yp_gamers).Error
	return err
}

// GetYp_gamers 根据id获取Yp_gamers记录
// Author [piexlmax](https://github.com/piexlmax)
func (yp_gamersService *Yp_gamersService)GetYp_gamers(id uint) (yp_gamers gameList.Yp_gamers, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&yp_gamers).Error
	return
}

// GetYp_gamersInfoList 分页获取Yp_gamers记录
// Author [piexlmax](https://github.com/piexlmax)
func (yp_gamersService *Yp_gamersService)GetYp_gamersInfoList(info gameListReq.Yp_gamersSearch) (list []gameList.Yp_gamers, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&gameList.Yp_gamers{})
    var yp_gamerss []gameList.Yp_gamers
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Uuid != "" {
        db = db.Where("uuid = ?",info.Uuid)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["playCount"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	err = db.Limit(limit).Offset(offset).Find(&yp_gamerss).Error
	return  yp_gamerss, total, err
}
