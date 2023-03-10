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

type Yp_gamesApi struct {
}

var yp_gamesService = service.ServiceGroupApp.GameListServiceGroup.Yp_gamesService


// CreateYp_games 创建Yp_games
// @Tags Yp_games
// @Summary 创建Yp_games
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gameList.Yp_games true "创建Yp_games"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /yp_games/createYp_games [post]
func (yp_gamesApi *Yp_gamesApi) CreateYp_games(c *gin.Context) {
	var yp_games gameList.Yp_games
	err := c.ShouldBindJSON(&yp_games)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    yp_games.CreatedBy = utils.GetUserID(c)
    verify := utils.Rules{
        "Game_name":{utils.NotEmpty()},
        "Game_grade":{utils.NotEmpty()},
    }
	if err := utils.Verify(yp_games, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := yp_gamesService.CreateYp_games(yp_games); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteYp_games 删除Yp_games
// @Tags Yp_games
// @Summary 删除Yp_games
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gameList.Yp_games true "删除Yp_games"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /yp_games/deleteYp_games [delete]
func (yp_gamesApi *Yp_gamesApi) DeleteYp_games(c *gin.Context) {
	var yp_games gameList.Yp_games
	err := c.ShouldBindJSON(&yp_games)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    yp_games.DeletedBy = utils.GetUserID(c)
	if err := yp_gamesService.DeleteYp_games(yp_games); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteYp_gamesByIds 批量删除Yp_games
// @Tags Yp_games
// @Summary 批量删除Yp_games
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Yp_games"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /yp_games/deleteYp_gamesByIds [delete]
func (yp_gamesApi *Yp_gamesApi) DeleteYp_gamesByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := yp_gamesService.DeleteYp_gamesByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateYp_games 更新Yp_games
// @Tags Yp_games
// @Summary 更新Yp_games
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body gameList.Yp_games true "更新Yp_games"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /yp_games/updateYp_games [put]
func (yp_gamesApi *Yp_gamesApi) UpdateYp_games(c *gin.Context) {
	var yp_games gameList.Yp_games
	err := c.ShouldBindJSON(&yp_games)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    yp_games.UpdatedBy = utils.GetUserID(c)
      verify := utils.Rules{
          "Game_name":{utils.NotEmpty()},
          "Game_grade":{utils.NotEmpty()},
      }
    if err := utils.Verify(yp_games, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := yp_gamesService.UpdateYp_games(yp_games); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindYp_games 用id查询Yp_games
// @Tags Yp_games
// @Summary 用id查询Yp_games
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query gameList.Yp_games true "用id查询Yp_games"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /yp_games/findYp_games [get]
func (yp_gamesApi *Yp_gamesApi) FindYp_games(c *gin.Context) {
	var yp_games gameList.Yp_games
	err := c.ShouldBindQuery(&yp_games)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reyp_games, err := yp_gamesService.GetYp_games(yp_games.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reyp_games": reyp_games}, c)
	}
}

// GetYp_gamesList 分页获取Yp_games列表
// @Tags Yp_games
// @Summary 分页获取Yp_games列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query gameListReq.Yp_gamesSearch true "分页获取Yp_games列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /yp_games/getYp_gamesList [get]
func (yp_gamesApi *Yp_gamesApi) GetYp_gamesList(c *gin.Context) {
	var pageInfo gameListReq.Yp_gamesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := yp_gamesService.GetYp_gamesInfoList(pageInfo); err != nil {
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
