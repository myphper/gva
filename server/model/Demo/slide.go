// 自动生成模板PkgSlide
package Demo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// PkgSlide 结构体
type PkgSlide struct {
      global.GVA_MODEL
      Url  string `json:"url" form:"url" gorm:"column:url;comment:图片地址;"`  //图片地址 
      IndexSort  *int `json:"indexSort" form:"indexSort" gorm:"column:index_sort;comment:排序;"`  //排序 
}


// TableName PkgSlide 表名
func (PkgSlide) TableName() string {
  return "slide"
}

