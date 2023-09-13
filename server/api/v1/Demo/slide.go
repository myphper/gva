package Demo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/Demo"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    DemoReq "github.com/flipped-aurora/gin-vue-admin/server/model/Demo/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type PkgSlideApi struct {
}

var SlideService = service.ServiceGroupApp.DemoServiceGroup.PkgSlideService


// CreatePkgSlide 创建PkgSlide
// @Tags PkgSlide
// @Summary 创建PkgSlide
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Demo.PkgSlide true "创建PkgSlide"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Slide/createPkgSlide [post]
func (SlideApi *PkgSlideApi) CreatePkgSlide(c *gin.Context) {
	var Slide Demo.PkgSlide
	err := c.ShouldBindJSON(&Slide)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Url":{utils.NotEmpty()},
    }
	if err := utils.Verify(Slide, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := SlideService.CreatePkgSlide(&Slide); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePkgSlide 删除PkgSlide
// @Tags PkgSlide
// @Summary 删除PkgSlide
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Demo.PkgSlide true "删除PkgSlide"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Slide/deletePkgSlide [delete]
func (SlideApi *PkgSlideApi) DeletePkgSlide(c *gin.Context) {
	var Slide Demo.PkgSlide
	err := c.ShouldBindJSON(&Slide)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := SlideService.DeletePkgSlide(Slide); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePkgSlideByIds 批量删除PkgSlide
// @Tags PkgSlide
// @Summary 批量删除PkgSlide
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PkgSlide"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /Slide/deletePkgSlideByIds [delete]
func (SlideApi *PkgSlideApi) DeletePkgSlideByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := SlideService.DeletePkgSlideByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePkgSlide 更新PkgSlide
// @Tags PkgSlide
// @Summary 更新PkgSlide
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Demo.PkgSlide true "更新PkgSlide"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Slide/updatePkgSlide [put]
func (SlideApi *PkgSlideApi) UpdatePkgSlide(c *gin.Context) {
	var Slide Demo.PkgSlide
	err := c.ShouldBindJSON(&Slide)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Url":{utils.NotEmpty()},
      }
    if err := utils.Verify(Slide, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := SlideService.UpdatePkgSlide(Slide); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPkgSlide 用id查询PkgSlide
// @Tags PkgSlide
// @Summary 用id查询PkgSlide
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Demo.PkgSlide true "用id查询PkgSlide"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Slide/findPkgSlide [get]
func (SlideApi *PkgSlideApi) FindPkgSlide(c *gin.Context) {
	var Slide Demo.PkgSlide
	err := c.ShouldBindQuery(&Slide)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reSlide, err := SlideService.GetPkgSlide(Slide.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reSlide": reSlide}, c)
	}
}

// GetPkgSlideList 分页获取PkgSlide列表
// @Tags PkgSlide
// @Summary 分页获取PkgSlide列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query DemoReq.PkgSlideSearch true "分页获取PkgSlide列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Slide/getPkgSlideList [get]
func (SlideApi *PkgSlideApi) GetPkgSlideList(c *gin.Context) {
	var pageInfo DemoReq.PkgSlideSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := SlideService.GetPkgSlideInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
