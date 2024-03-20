package middleware

import (
	"github.com/gin-gonic/gin"
	"server/app/casbin/service"
	"server/internal/global"
	"strings"
)

var casbinService = service.CasbinServiceApp

func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求资源
		path := c.Request.URL.Path
		obj := strings.TrimPrefix(path, global.Config.System.RouterPrefix)
		// 获取请求方式
		act := c.Request.Method
		// 获取请求主体
		sub := "admin"

		// 判断是否存在对应的 ACL
		casbin := casbinService.GetCasbin()
		ok, _ := casbin.Enforce(sub, obj, act)
		if ok {
			c.Next() // 请求成功
		} else {
			c.Abort() // 请求失败
			return
		}
	}
}
