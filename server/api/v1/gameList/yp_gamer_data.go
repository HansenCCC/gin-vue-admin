package gameList

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/gameList"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    gameListReq "github.com/flipped-aurora/gin-vue-admin/server/model/gameList/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type Yp_gamer_dataApi struct {
}

var yp_gamer_dataService = service.ServiceGroupApp.GameListServiceGroup.Yp_gamer_dataService


// CreateYp_gamer_data 创建Yp_gamer_data
// @Tags Yp_gamer_data
// @Summary 创建Yp_gamer_data
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gameList.Yp_gamer_data true "创建Yp_gamer_data"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /yp_gamer_data/createYp_gamer_data [post]
func (yp_gamer_dataApi *Yp_gamer_dataApi) CreateYp_gamer_data(c *gin.Context) {
	var yp_gamer_data gameList.Yp_gamer_data
	err := c.ShouldBindJSON(&yp_gamer_data)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    yp_gamer_data.CreatedBy = utils.GetUserID(c)
    verify := utils.Rules{
        "Gamer_username":{utils.NotEmpty()},
        "Gamer_move_count":{utils.NotEmpty()},
        "Gameover_time":{utils.NotEmpty()},
    }
	if err := utils.Verify(yp_gamer_data, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := yp_gamer_dataService.CreateYp_gamer_data(yp_gamer_data); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteYp_gamer_data 删除Yp_gamer_data
// @Tags Yp_gamer_data
// @Summary 删除Yp_gamer_data
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gameList.Yp_gamer_data true "删除Yp_gamer_data"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /yp_gamer_data/deleteYp_gamer_data [delete]
func (yp_gamer_dataApi *Yp_gamer_dataApi) DeleteYp_gamer_data(c *gin.Context) {
	var yp_gamer_data gameList.Yp_gamer_data
	err := c.ShouldBindJSON(&yp_gamer_data)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    yp_gamer_data.DeletedBy = utils.GetUserID(c)
	if err := yp_gamer_dataService.DeleteYp_gamer_data(yp_gamer_data); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteYp_gamer_dataByIds 批量删除Yp_gamer_data
// @Tags Yp_gamer_data
// @Summary 批量删除Yp_gamer_data
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Yp_gamer_data"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /yp_gamer_data/deleteYp_gamer_dataByIds [delete]
func (yp_gamer_dataApi *Yp_gamer_dataApi) DeleteYp_gamer_dataByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := yp_gamer_dataService.DeleteYp_gamer_dataByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateYp_gamer_data 更新Yp_gamer_data
// @Tags Yp_gamer_data
// @Summary 更新Yp_gamer_data
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gameList.Yp_gamer_data true "更新Yp_gamer_data"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /yp_gamer_data/updateYp_gamer_data [put]
func (yp_gamer_dataApi *Yp_gamer_dataApi) UpdateYp_gamer_data(c *gin.Context) {
	var yp_gamer_data gameList.Yp_gamer_data
	err := c.ShouldBindJSON(&yp_gamer_data)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    yp_gamer_data.UpdatedBy = utils.GetUserID(c)
      verify := utils.Rules{
          "Gamer_username":{utils.NotEmpty()},
          "Gamer_move_count":{utils.NotEmpty()},
          "Gameover_time":{utils.NotEmpty()},
      }
    if err := utils.Verify(yp_gamer_data, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := yp_gamer_dataService.UpdateYp_gamer_data(yp_gamer_data); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindYp_gamer_data 用id查询Yp_gamer_data
// @Tags Yp_gamer_data
// @Summary 用id查询Yp_gamer_data
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query gameList.Yp_gamer_data true "用id查询Yp_gamer_data"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /yp_gamer_data/findYp_gamer_data [get]
func (yp_gamer_dataApi *Yp_gamer_dataApi) FindYp_gamer_data(c *gin.Context) {
	var yp_gamer_data gameList.Yp_gamer_data
	err := c.ShouldBindQuery(&yp_gamer_data)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reyp_gamer_data, err := yp_gamer_dataService.GetYp_gamer_data(yp_gamer_data.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reyp_gamer_data": reyp_gamer_data}, c)
	}
}

// GetYp_gamer_dataList 分页获取Yp_gamer_data列表
// @Tags Yp_gamer_data
// @Summary 分页获取Yp_gamer_data列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query gameListReq.Yp_gamer_dataSearch true "分页获取Yp_gamer_data列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /yp_gamer_data/getYp_gamer_dataList [get]
func (yp_gamer_dataApi *Yp_gamer_dataApi) GetYp_gamer_dataList(c *gin.Context) {
	var pageInfo gameListReq.Yp_gamer_dataSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := yp_gamer_dataService.GetYp_gamer_dataInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
