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

type Yp_game_eventApi struct {
}

var yp_game_eventService = service.ServiceGroupApp.GameListServiceGroup.Yp_game_eventService


// CreateYp_game_event 创建Yp_game_event
// @Tags Yp_game_event
// @Summary 创建Yp_game_event
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gameList.Yp_game_event true "创建Yp_game_event"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /yp_game_event/createYp_game_event [post]
func (yp_game_eventApi *Yp_game_eventApi) CreateYp_game_event(c *gin.Context) {
	var yp_game_event gameList.Yp_game_event
	err := c.ShouldBindJSON(&yp_game_event)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    yp_game_event.CreatedBy = utils.GetUserID(c)
    verify := utils.Rules{
        "Event_name":{utils.NotEmpty()},
    }
	if err := utils.Verify(yp_game_event, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := yp_game_eventService.CreateYp_game_event(yp_game_event); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteYp_game_event 删除Yp_game_event
// @Tags Yp_game_event
// @Summary 删除Yp_game_event
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gameList.Yp_game_event true "删除Yp_game_event"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /yp_game_event/deleteYp_game_event [delete]
func (yp_game_eventApi *Yp_game_eventApi) DeleteYp_game_event(c *gin.Context) {
	var yp_game_event gameList.Yp_game_event
	err := c.ShouldBindJSON(&yp_game_event)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    yp_game_event.DeletedBy = utils.GetUserID(c)
	if err := yp_game_eventService.DeleteYp_game_event(yp_game_event); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteYp_game_eventByIds 批量删除Yp_game_event
// @Tags Yp_game_event
// @Summary 批量删除Yp_game_event
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Yp_game_event"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /yp_game_event/deleteYp_game_eventByIds [delete]
func (yp_game_eventApi *Yp_game_eventApi) DeleteYp_game_eventByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := yp_game_eventService.DeleteYp_game_eventByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateYp_game_event 更新Yp_game_event
// @Tags Yp_game_event
// @Summary 更新Yp_game_event
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gameList.Yp_game_event true "更新Yp_game_event"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /yp_game_event/updateYp_game_event [put]
func (yp_game_eventApi *Yp_game_eventApi) UpdateYp_game_event(c *gin.Context) {
	var yp_game_event gameList.Yp_game_event
	err := c.ShouldBindJSON(&yp_game_event)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    yp_game_event.UpdatedBy = utils.GetUserID(c)
      verify := utils.Rules{
          "Event_name":{utils.NotEmpty()},
      }
    if err := utils.Verify(yp_game_event, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := yp_game_eventService.UpdateYp_game_event(yp_game_event); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindYp_game_event 用id查询Yp_game_event
// @Tags Yp_game_event
// @Summary 用id查询Yp_game_event
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query gameList.Yp_game_event true "用id查询Yp_game_event"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /yp_game_event/findYp_game_event [get]
func (yp_game_eventApi *Yp_game_eventApi) FindYp_game_event(c *gin.Context) {
	var yp_game_event gameList.Yp_game_event
	err := c.ShouldBindQuery(&yp_game_event)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reyp_game_event, err := yp_game_eventService.GetYp_game_event(yp_game_event.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reyp_game_event": reyp_game_event}, c)
	}
}

// GetYp_game_eventList 分页获取Yp_game_event列表
// @Tags Yp_game_event
// @Summary 分页获取Yp_game_event列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query gameListReq.Yp_game_eventSearch true "分页获取Yp_game_event列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /yp_game_event/getYp_game_eventList [get]
func (yp_game_eventApi *Yp_game_eventApi) GetYp_game_eventList(c *gin.Context) {
	var pageInfo gameListReq.Yp_game_eventSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := yp_game_eventService.GetYp_game_eventInfoList(pageInfo); err != nil {
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
