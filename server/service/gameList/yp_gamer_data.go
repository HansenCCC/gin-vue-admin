package gameList

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/gameList"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    gameListReq "github.com/flipped-aurora/gin-vue-admin/server/model/gameList/request"
    "gorm.io/gorm"
)

type Yp_gamer_dataService struct {
}

// CreateYp_gamer_data 创建Yp_gamer_data记录
// Author [piexlmax](https://github.com/piexlmax)
func (yp_gamer_dataService *Yp_gamer_dataService) CreateYp_gamer_data(yp_gamer_data gameList.Yp_gamer_data) (err error) {
	err = global.GVA_DB.Create(&yp_gamer_data).Error
	return err
}

// DeleteYp_gamer_data 删除Yp_gamer_data记录
// Author [piexlmax](https://github.com/piexlmax)
func (yp_gamer_dataService *Yp_gamer_dataService)DeleteYp_gamer_data(yp_gamer_data gameList.Yp_gamer_data) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&gameList.Yp_gamer_data{}).Where("id = ?", yp_gamer_data.ID).Update("deleted_by", yp_gamer_data.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&yp_gamer_data).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteYp_gamer_dataByIds 批量删除Yp_gamer_data记录
// Author [piexlmax](https://github.com/piexlmax)
func (yp_gamer_dataService *Yp_gamer_dataService)DeleteYp_gamer_dataByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&gameList.Yp_gamer_data{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&gameList.Yp_gamer_data{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateYp_gamer_data 更新Yp_gamer_data记录
// Author [piexlmax](https://github.com/piexlmax)
func (yp_gamer_dataService *Yp_gamer_dataService)UpdateYp_gamer_data(yp_gamer_data gameList.Yp_gamer_data) (err error) {
	err = global.GVA_DB.Save(&yp_gamer_data).Error
	return err
}

// GetYp_gamer_data 根据id获取Yp_gamer_data记录
// Author [piexlmax](https://github.com/piexlmax)
func (yp_gamer_dataService *Yp_gamer_dataService)GetYp_gamer_data(id uint) (yp_gamer_data gameList.Yp_gamer_data, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&yp_gamer_data).Error
	return
}

// GetYp_gamer_dataInfoList 分页获取Yp_gamer_data记录
// Author [piexlmax](https://github.com/piexlmax)
func (yp_gamer_dataService *Yp_gamer_dataService)GetYp_gamer_dataInfoList(info gameListReq.Yp_gamer_dataSearch) (list []gameList.Yp_gamer_data, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&gameList.Yp_gamer_data{})
    var yp_gamer_datas []gameList.Yp_gamer_data
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Gamer_move_count != nil {
        db = db.Where("gamer_move_count <> ?",info.Gamer_move_count)
    }
    if info.Gameover_time != nil {
        db = db.Where("gameover_time <> ?",info.Gameover_time)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["gamer_move_count"] = true
         	orderMap["gameover_time"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	err = db.Limit(limit).Offset(offset).Find(&yp_gamer_datas).Error
	return  yp_gamer_datas, total, err
}
