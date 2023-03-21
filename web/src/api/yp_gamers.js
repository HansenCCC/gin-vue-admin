import service from '@/utils/request'

// @Tags Yp_gamers
// @Summary 创建Yp_gamers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Yp_gamers true "创建Yp_gamers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /yp_gamers/createYp_gamers [post]
export const createYp_gamers = (data) => {
  return service({
    url: '/yp_gamers/createYp_gamers',
    method: 'post',
    data
  })
}

// @Tags Yp_gamers
// @Summary 删除Yp_gamers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Yp_gamers true "删除Yp_gamers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /yp_gamers/deleteYp_gamers [delete]
export const deleteYp_gamers = (data) => {
  return service({
    url: '/yp_gamers/deleteYp_gamers',
    method: 'delete',
    data
  })
}

// @Tags Yp_gamers
// @Summary 删除Yp_gamers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Yp_gamers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /yp_gamers/deleteYp_gamers [delete]
export const deleteYp_gamersByIds = (data) => {
  return service({
    url: '/yp_gamers/deleteYp_gamersByIds',
    method: 'delete',
    data
  })
}

// @Tags Yp_gamers
// @Summary 更新Yp_gamers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Yp_gamers true "更新Yp_gamers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /yp_gamers/updateYp_gamers [put]
export const updateYp_gamers = (data) => {
  return service({
    url: '/yp_gamers/updateYp_gamers',
    method: 'put',
    data
  })
}

// @Tags Yp_gamers
// @Summary 用id查询Yp_gamers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Yp_gamers true "用id查询Yp_gamers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /yp_gamers/findYp_gamers [get]
export const findYp_gamers = (params) => {
  return service({
    url: '/yp_gamers/findYp_gamers',
    method: 'get',
    params
  })
}

// @Tags Yp_gamers
// @Summary 分页获取Yp_gamers列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Yp_gamers列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /yp_gamers/getYp_gamersList [get]
export const getYp_gamersList = (params) => {
  return service({
    url: '/yp_gamers/getYp_gamersList',
    method: 'get',
    params
  })
}
