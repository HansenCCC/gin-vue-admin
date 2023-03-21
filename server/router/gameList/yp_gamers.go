package gameList

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type Yp_gamersRouter struct {
}

// InitYp_gamersRouter 初始化 Yp_gamers 路由信息
func (s *Yp_gamersRouter) InitYp_gamersRouter(Router *gin.RouterGroup) {
	yp_gamersRouter := Router.Group("yp_gamers").Use(middleware.OperationRecord())
	yp_gamersRouterWithoutRecord := Router.Group("yp_gamers")
	var yp_gamersApi = v1.ApiGroupApp.GameListApiGroup.Yp_gamersApi
	{
		yp_gamersRouter.POST("createYp_gamers", yp_gamersApi.CreateYp_gamers)   // 新建Yp_gamers
		yp_gamersRouter.DELETE("deleteYp_gamers", yp_gamersApi.DeleteYp_gamers) // 删除Yp_gamers
		yp_gamersRouter.DELETE("deleteYp_gamersByIds", yp_gamersApi.DeleteYp_gamersByIds) // 批量删除Yp_gamers
		yp_gamersRouter.PUT("updateYp_gamers", yp_gamersApi.UpdateYp_gamers)    // 更新Yp_gamers
	}
	{
		yp_gamersRouterWithoutRecord.GET("findYp_gamers", yp_gamersApi.FindYp_gamers)        // 根据ID获取Yp_gamers
		yp_gamersRouterWithoutRecord.GET("getYp_gamersList", yp_gamersApi.GetYp_gamersList)  // 获取Yp_gamers列表
	}
}
