package Demo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PkgSlideRouter struct {
}

// InitPkgSlideRouter 初始化 PkgSlide 路由信息
func (s *PkgSlideRouter) InitPkgSlideRouter(Router *gin.RouterGroup) {
	SlideRouter := Router.Group("Slide").Use(middleware.OperationRecord())
	SlideRouterWithoutRecord := Router.Group("Slide")
	var SlideApi = v1.ApiGroupApp.DemoApiGroup.PkgSlideApi
	{
		SlideRouter.POST("createPkgSlide", SlideApi.CreatePkgSlide)   // 新建PkgSlide
		SlideRouter.DELETE("deletePkgSlide", SlideApi.DeletePkgSlide) // 删除PkgSlide
		SlideRouter.DELETE("deletePkgSlideByIds", SlideApi.DeletePkgSlideByIds) // 批量删除PkgSlide
		SlideRouter.PUT("updatePkgSlide", SlideApi.UpdatePkgSlide)    // 更新PkgSlide
	}
	{
		SlideRouterWithoutRecord.GET("findPkgSlide", SlideApi.FindPkgSlide)        // 根据ID获取PkgSlide
		SlideRouterWithoutRecord.GET("getPkgSlideList", SlideApi.GetPkgSlideList)  // 获取PkgSlide列表
	}
}
