package gameList

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type Yp_gamer_dataRouter struct {
}

// InitYp_gamer_dataRouter 初始化 Yp_gamer_data 路由信息
func (s *Yp_gamer_dataRouter) InitYp_gamer_dataRouter(Router *gin.RouterGroup) {
	yp_gamer_dataRouter := Router.Group("yp_gamer_data").Use(middleware.OperationRecord())
	yp_gamer_dataRouterWithoutRecord := Router.Group("yp_gamer_data")
	var yp_gamer_dataApi = v1.ApiGroupApp.GameListApiGroup.Yp_gamer_dataApi
	{
		yp_gamer_dataRouter.POST("createYp_gamer_data", yp_gamer_dataApi.CreateYp_gamer_data)   // 新建Yp_gamer_data
		yp_gamer_dataRouter.DELETE("deleteYp_gamer_data", yp_gamer_dataApi.DeleteYp_gamer_data) // 删除Yp_gamer_data
		yp_gamer_dataRouter.DELETE("deleteYp_gamer_dataByIds", yp_gamer_dataApi.DeleteYp_gamer_dataByIds) // 批量删除Yp_gamer_data
		yp_gamer_dataRouter.PUT("updateYp_gamer_data", yp_gamer_dataApi.UpdateYp_gamer_data)    // 更新Yp_gamer_data
	}
	{
		yp_gamer_dataRouterWithoutRecord.GET("findYp_gamer_data", yp_gamer_dataApi.FindYp_gamer_data)        // 根据ID获取Yp_gamer_data
		yp_gamer_dataRouterWithoutRecord.GET("getYp_gamer_dataList", yp_gamer_dataApi.GetYp_gamer_dataList)  // 获取Yp_gamer_data列表
	}
}
