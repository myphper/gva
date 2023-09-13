package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/Demo"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Demo    Demo.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
