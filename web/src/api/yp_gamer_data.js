import service from '@/utils/request'

// @Tags Yp_gamer_data
// @Summary 创建Yp_gamer_data
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Yp_gamer_data true "创建Yp_gamer_data"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /yp_gamer_data/createYp_gamer_data [post]
export const createYp_gamer_data = (data) => {
  return service({
    url: '/yp_gamer_data/createYp_gamer_data',
    method: 'post',
    data
  })
}

// @Tags Yp_gamer_data
// @Summary 删除Yp_gamer_data
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Yp_gamer_data true "删除Yp_gamer_data"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /yp_gamer_data/deleteYp_gamer_data [delete]
export const deleteYp_gamer_data = (data) => {
  return service({
    url: '/yp_gamer_data/deleteYp_gamer_data',
    method: 'delete',
    data
  })
}

// @Tags Yp_gamer_data
// @Summary 删除Yp_gamer_data
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Yp_gamer_data"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /yp_gamer_data/deleteYp_gamer_data [delete]
export const deleteYp_gamer_dataByIds = (data) => {
  return service({
    url: '/yp_gamer_data/deleteYp_gamer_dataByIds',
    method: 'delete',
    data
  })
}

// @Tags Yp_gamer_data
// @Summary 更新Yp_gamer_data
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Yp_gamer_data true "更新Yp_gamer_data"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /yp_gamer_data/updateYp_gamer_data [put]
export const updateYp_gamer_data = (data) => {
  return service({
    url: '/yp_gamer_data/updateYp_gamer_data',
    method: 'put',
    data
  })
}

// @Tags Yp_gamer_data
// @Summary 用id查询Yp_gamer_data
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Yp_gamer_data true "用id查询Yp_gamer_data"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /yp_gamer_data/findYp_gamer_data [get]
export const findYp_gamer_data = (params) => {
  return service({
    url: '/yp_gamer_data/findYp_gamer_data',
    method: 'get',
    params
  })
}

// @Tags Yp_gamer_data
// @Summary 分页获取Yp_gamer_data列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Yp_gamer_data列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /yp_gamer_data/getYp_gamer_dataList [get]
export const getYp_gamer_dataList = (params) => {
  return service({
    url: '/yp_gamer_data/getYp_gamer_dataList',
    method: 'get',
    params
  })
}
