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

type Yp_gamersApi struct {
}

var yp_gamersService = service.ServiceGroupApp.GameListServiceGroup.Yp_gamersService


// CreateYp_gamers 创建Yp_gamers
// @Tags Yp_gamers
// @Summary 创建Yp_gamers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gameList.Yp_gamers true "创建Yp_gamers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /yp_gamers/createYp_gamers [post]
func (yp_gamersApi *Yp_gamersApi) CreateYp_gamers(c *gin.Context) {
	var yp_gamers gameList.Yp_gamers
	err := c.ShouldBindJSON(&yp_gamers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    yp_gamers.CreatedBy = utils.GetUserID(c)
    verify := utils.Rules{
        "Uuid":{utils.NotEmpty()},
    }
	if err := utils.Verify(yp_gamers, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := yp_gamersService.CreateYp_gamers(yp_gamers); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteYp_gamers 删除Yp_gamers
// @Tags Yp_gamers
// @Summary 删除Yp_gamers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gameList.Yp_gamers true "删除Yp_gamers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /yp_gamers/deleteYp_gamers [delete]
func (yp_gamersApi *Yp_gamersApi) DeleteYp_gamers(c *gin.Context) {
	var yp_gamers gameList.Yp_gamers
	err := c.ShouldBindJSON(&yp_gamers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    yp_gamers.DeletedBy = utils.GetUserID(c)
	if err := yp_gamersService.DeleteYp_gamers(yp_gamers); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteYp_gamersByIds 批量删除Yp_gamers
// @Tags Yp_gamers
// @Summary 批量删除Yp_gamers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Yp_gamers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /yp_gamers/deleteYp_gamersByIds [delete]
func (yp_gamersApi *Yp_gamersApi) DeleteYp_gamersByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := yp_gamersService.DeleteYp_gamersByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateYp_gamers 更新Yp_gamers
// @Tags Yp_gamers
// @Summary 更新Yp_gamers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gameList.Yp_gamers true "更新Yp_gamers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /yp_gamers/updateYp_gamers [put]
func (yp_gamersApi *Yp_gamersApi) UpdateYp_gamers(c *gin.Context) {
	var yp_gamers gameList.Yp_gamers
	err := c.ShouldBindJSON(&yp_gamers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    yp_gamers.UpdatedBy = utils.GetUserID(c)
      verify := utils.Rules{
          "Uuid":{utils.NotEmpty()},
      }
    if err := utils.Verify(yp_gamers, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := yp_gamersService.UpdateYp_gamers(yp_gamers); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindYp_gamers 用id查询Yp_gamers
// @Tags Yp_gamers
// @Summary 用id查询Yp_gamers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query gameList.Yp_gamers true "用id查询Yp_gamers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /yp_gamers/findYp_gamers [get]
func (yp_gamersApi *Yp_gamersApi) FindYp_gamers(c *gin.Context) {
	var yp_gamers gameList.Yp_gamers
	err := c.ShouldBindQuery(&yp_gamers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reyp_gamers, err := yp_gamersService.GetYp_gamers(yp_gamers.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reyp_gamers": reyp_gamers}, c)
	}
}

// GetYp_gamersList 分页获取Yp_gamers列表
// @Tags Yp_gamers
// @Summary 分页获取Yp_gamers列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query gameListReq.Yp_gamersSearch true "分页获取Yp_gamers列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /yp_gamers/getYp_gamersList [get]
func (yp_gamersApi *Yp_gamersApi) GetYp_gamersList(c *gin.Context) {
	var pageInfo gameListReq.Yp_gamersSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := yp_gamersService.GetYp_gamersInfoList(pageInfo); err != nil {
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
