import service from '@/utils/request'

// @Tags Yp_game_event
// @Summary 创建Yp_game_event
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Yp_game_event true "创建Yp_game_event"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /yp_game_event/createYp_game_event [post]
export const createYp_game_event = (data) => {
  return service({
    url: '/yp_game_event/createYp_game_event',
    method: 'post',
    data
  })
}

// @Tags Yp_game_event
// @Summary 删除Yp_game_event
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Yp_game_event true "删除Yp_game_event"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /yp_game_event/deleteYp_game_event [delete]
export const deleteYp_game_event = (data) => {
  return service({
    url: '/yp_game_event/deleteYp_game_event',
    method: 'delete',
    data
  })
}

// @Tags Yp_game_event
// @Summary 删除Yp_game_event
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Yp_game_event"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /yp_game_event/deleteYp_game_event [delete]
export const deleteYp_game_eventByIds = (data) => {
  return service({
    url: '/yp_game_event/deleteYp_game_eventByIds',
    method: 'delete',
    data
  })
}

// @Tags Yp_game_event
// @Summary 更新Yp_game_event
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Yp_game_event true "更新Yp_game_event"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /yp_game_event/updateYp_game_event [put]
export const updateYp_game_event = (data) => {
  return service({
    url: '/yp_game_event/updateYp_game_event',
    method: 'put',
    data
  })
}

// @Tags Yp_game_event
// @Summary 用id查询Yp_game_event
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Yp_game_event true "用id查询Yp_game_event"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /yp_game_event/findYp_game_event [get]
export const findYp_game_event = (params) => {
  return service({
    url: '/yp_game_event/findYp_game_event',
    method: 'get',
    params
  })
}

// @Tags Yp_game_event
// @Summary 分页获取Yp_game_event列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Yp_game_event列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /yp_game_event/getYp_game_eventList [get]
export const getYp_game_eventList = (params) => {
  return service({
    url: '/yp_game_event/getYp_game_eventList',
    method: 'get',
    params
  })
}
