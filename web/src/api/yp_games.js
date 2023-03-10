import service from '@/utils/request'

// @Tags Yp_games
// @Summary 创建Yp_games
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Yp_games true "创建Yp_games"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /yp_games/createYp_games [post]
export const createYp_games = (data) => {
  return service({
    url: '/yp_games/createYp_games',
    method: 'post',
    data
  })
}

// @Tags Yp_games
// @Summary 删除Yp_games
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Yp_games true "删除Yp_games"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /yp_games/deleteYp_games [delete]
export const deleteYp_games = (data) => {
  return service({
    url: '/yp_games/deleteYp_games',
    method: 'delete',
    data
  })
}

// @Tags Yp_games
// @Summary 删除Yp_games
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Yp_games"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /yp_games/deleteYp_games [delete]
export const deleteYp_gamesByIds = (data) => {
  return service({
    url: '/yp_games/deleteYp_gamesByIds',
    method: 'delete',
    data
  })
}

// @Tags Yp_games
// @Summary 更新Yp_games
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Yp_games true "更新Yp_games"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /yp_games/updateYp_games [put]
export const updateYp_games = (data) => {
  return service({
    url: '/yp_games/updateYp_games',
    method: 'put',
    data
  })
}

// @Tags Yp_games
// @Summary 用id查询Yp_games
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Yp_games true "用id查询Yp_games"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /yp_games/findYp_games [get]
export const findYp_games = (params) => {
  return service({
    url: '/yp_games/findYp_games',
    method: 'get',
    params
  })
}

// @Tags Yp_games
// @Summary 分页获取Yp_games列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Yp_games列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /yp_games/getYp_gamesList [get]
export const getYp_gamesList = (params) => {
  return service({
    url: '/yp_games/getYp_gamesList',
    method: 'get',
    params
  })
}
