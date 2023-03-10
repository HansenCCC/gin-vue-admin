package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/gameList"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type Yp_game_eventSearch struct{
    gameList.Yp_game_event
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
