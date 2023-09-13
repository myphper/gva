import service from '@/utils/request'

// @Tags PkgSlide
// @Summary 创建PkgSlide
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PkgSlide true "创建PkgSlide"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Slide/createPkgSlide [post]
export const createPkgSlide = (data) => {
  return service({
    url: '/Slide/createPkgSlide',
    method: 'post',
    data
  })
}

// @Tags PkgSlide
// @Summary 删除PkgSlide
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PkgSlide true "删除PkgSlide"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Slide/deletePkgSlide [delete]
export const deletePkgSlide = (data) => {
  return service({
    url: '/Slide/deletePkgSlide',
    method: 'delete',
    data
  })
}

// @Tags PkgSlide
// @Summary 删除PkgSlide
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PkgSlide"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Slide/deletePkgSlide [delete]
export const deletePkgSlideByIds = (data) => {
  return service({
    url: '/Slide/deletePkgSlideByIds',
    method: 'delete',
    data
  })
}

// @Tags PkgSlide
// @Summary 更新PkgSlide
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PkgSlide true "更新PkgSlide"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Slide/updatePkgSlide [put]
export const updatePkgSlide = (data) => {
  return service({
    url: '/Slide/updatePkgSlide',
    method: 'put',
    data
  })
}

// @Tags PkgSlide
// @Summary 用id查询PkgSlide
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PkgSlide true "用id查询PkgSlide"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Slide/findPkgSlide [get]
export const findPkgSlide = (params) => {
  return service({
    url: '/Slide/findPkgSlide',
    method: 'get',
    params
  })
}

// @Tags PkgSlide
// @Summary 分页获取PkgSlide列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取PkgSlide列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Slide/getPkgSlideList [get]
export const getPkgSlideList = (params) => {
  return service({
    url: '/Slide/getPkgSlideList',
    method: 'get',
    params
  })
}
