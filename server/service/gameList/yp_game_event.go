package gameList

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/gameList"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    gameListReq "github.com/flipped-aurora/gin-vue-admin/server/model/gameList/request"
    "gorm.io/gorm"
)

type Yp_game_eventService struct {
}

// CreateYp_game_event 创建Yp_game_event记录
// Author [piexlmax](https://github.com/piexlmax)
func (yp_game_eventService *Yp_game_eventService) CreateYp_game_event(yp_game_event gameList.Yp_game_event) (err error) {
	err = global.GVA_DB.Create(&yp_game_event).Error
	return err
}

// DeleteYp_game_event 删除Yp_game_event记录
// Author [piexlmax](https://github.com/piexlmax)
func (yp_game_eventService *Yp_game_eventService)DeleteYp_game_event(yp_game_event gameList.Yp_game_event) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&gameList.Yp_game_event{}).Where("id = ?", yp_game_event.ID).Update("deleted_by", yp_game_event.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&yp_game_event).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteYp_game_eventByIds 批量删除Yp_game_event记录
// Author [piexlmax](https://github.com/piexlmax)
func (yp_game_eventService *Yp_game_eventService)DeleteYp_game_eventByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&gameList.Yp_game_event{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&gameList.Yp_game_event{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateYp_game_event 更新Yp_game_event记录
// Author [piexlmax](https://github.com/piexlmax)
func (yp_game_eventService *Yp_game_eventService)UpdateYp_game_event(yp_game_event gameList.Yp_game_event) (err error) {
	err = global.GVA_DB.Save(&yp_game_event).Error
	return err
}

// GetYp_game_event 根据id获取Yp_game_event记录
// Author [piexlmax](https://github.com/piexlmax)
func (yp_game_eventService *Yp_game_eventService)GetYp_game_event(id uint) (yp_game_event gameList.Yp_game_event, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&yp_game_event).Error
	return
}

// GetYp_game_eventInfoList 分页获取Yp_game_event记录
// Author [piexlmax](https://github.com/piexlmax)
func (yp_game_eventService *Yp_game_eventService)GetYp_game_eventInfoList(info gameListReq.Yp_game_eventSearch) (list []gameList.Yp_game_event, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&gameList.Yp_game_event{})
    var yp_game_events []gameList.Yp_game_event
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Event_name != "" {
        db = db.Where("event_name = ?",info.Event_name)
    }
    if info.Uuid != "" {
        db = db.Where("uuid = ?",info.Uuid)
    }
    if info.Properties != "" {
        db = db.Where("properties LIKE ?","%"+ info.Properties+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&yp_game_events).Error
	return  yp_game_events, total, err
}
