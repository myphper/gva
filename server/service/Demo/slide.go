package Demo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Demo"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    DemoReq "github.com/flipped-aurora/gin-vue-admin/server/model/Demo/request"
)

type PkgSlideService struct {
}

// CreatePkgSlide 创建PkgSlide记录
// Author [piexlmax](https://github.com/piexlmax)
func (SlideService *PkgSlideService) CreatePkgSlide(Slide *Demo.PkgSlide) (err error) {
	err = global.GVA_DB.Create(Slide).Error
	return err
}

// DeletePkgSlide 删除PkgSlide记录
// Author [piexlmax](https://github.com/piexlmax)
func (SlideService *PkgSlideService)DeletePkgSlide(Slide Demo.PkgSlide) (err error) {
	err = global.GVA_DB.Delete(&Slide).Error
	return err
}

// DeletePkgSlideByIds 批量删除PkgSlide记录
// Author [piexlmax](https://github.com/piexlmax)
func (SlideService *PkgSlideService)DeletePkgSlideByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]Demo.PkgSlide{},"id in ?",ids.Ids).Error
	return err
}

// UpdatePkgSlide 更新PkgSlide记录
// Author [piexlmax](https://github.com/piexlmax)
func (SlideService *PkgSlideService)UpdatePkgSlide(Slide Demo.PkgSlide) (err error) {
	err = global.GVA_DB.Save(&Slide).Error
	return err
}

// GetPkgSlide 根据id获取PkgSlide记录
// Author [piexlmax](https://github.com/piexlmax)
func (SlideService *PkgSlideService)GetPkgSlide(id uint) (Slide Demo.PkgSlide, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&Slide).Error
	return
}

// GetPkgSlideInfoList 分页获取PkgSlide记录
// Author [piexlmax](https://github.com/piexlmax)
func (SlideService *PkgSlideService)GetPkgSlideInfoList(info DemoReq.PkgSlideSearch) (list []Demo.PkgSlide, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&Demo.PkgSlide{})
    var Slides []Demo.PkgSlide
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&Slides).Error
	return  Slides, total, err
}
