package gameList

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type Yp_game_eventRouter struct {
}

// InitYp_game_eventRouter 初始化 Yp_game_event 路由信息
func (s *Yp_game_eventRouter) InitYp_game_eventRouter(Router *gin.RouterGroup) {
	yp_game_eventRouter := Router.Group("yp_game_event").Use(middleware.OperationRecord())
	yp_game_eventRouterWithoutRecord := Router.Group("yp_game_event")
	var yp_game_eventApi = v1.ApiGroupApp.GameListApiGroup.Yp_game_eventApi
	{
		yp_game_eventRouter.POST("createYp_game_event", yp_game_eventApi.CreateYp_game_event)   // 新建Yp_game_event
		yp_game_eventRouter.DELETE("deleteYp_game_event", yp_game_eventApi.DeleteYp_game_event) // 删除Yp_game_event
		yp_game_eventRouter.DELETE("deleteYp_game_eventByIds", yp_game_eventApi.DeleteYp_game_eventByIds) // 批量删除Yp_game_event
		yp_game_eventRouter.PUT("updateYp_game_event", yp_game_eventApi.UpdateYp_game_event)    // 更新Yp_game_event
	}
	{
		yp_game_eventRouterWithoutRecord.GET("findYp_game_event", yp_game_eventApi.FindYp_game_event)        // 根据ID获取Yp_game_event
		yp_game_eventRouterWithoutRecord.GET("getYp_game_eventList", yp_game_eventApi.GetYp_game_eventList)  // 获取Yp_game_event列表
	}
}
