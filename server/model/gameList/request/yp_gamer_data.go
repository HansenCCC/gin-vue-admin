package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/gameList"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type Yp_gamer_dataSearch struct{
    gameList.Yp_gamer_data
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}
