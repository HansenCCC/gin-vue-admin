package gameList

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type Yp_gamesRouter struct {
}

// InitYp_gamesRouter 初始化 Yp_games 路由信息
func (s *Yp_gamesRouter) InitYp_gamesRouter(Router *gin.RouterGroup) {
	yp_gamesRouter := Router.Group("yp_games").Use(middleware.OperationRecord())
	yp_gamesRouterWithoutRecord := Router.Group("yp_games")
	var yp_gamesApi = v1.ApiGroupApp.GameListApiGroup.Yp_gamesApi
	{
		yp_gamesRouter.POST("createYp_games", yp_gamesApi.CreateYp_games)   // 新建Yp_games
		yp_gamesRouter.DELETE("deleteYp_games", yp_gamesApi.DeleteYp_games) // 删除Yp_games
		yp_gamesRouter.DELETE("deleteYp_gamesByIds", yp_gamesApi.DeleteYp_gamesByIds) // 批量删除Yp_games
		yp_gamesRouter.PUT("updateYp_games", yp_gamesApi.UpdateYp_games)    // 更新Yp_games
	}
	{
		yp_gamesRouterWithoutRecord.GET("findYp_games", yp_gamesApi.FindYp_games)        // 根据ID获取Yp_games
		yp_gamesRouterWithoutRecord.GET("getYp_gamesList", yp_gamesApi.GetYp_gamesList)  // 获取Yp_games列表
	}
}
