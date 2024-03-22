package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/model/response"
	"server/internal/service"
	"server/internal/util"
	"strconv"
)

var casbinService = service.CasbinServiceApp

func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求资源
		obj := c.Request.URL.Path
		// 获取请求方式
		act := c.Request.Method

		// 获取请求主体（身份 id）
		claims, _ := c.Get("claims")
		typedClaims, _ := claims.(*util.CustomClaims)
		sub := strconv.Itoa(int(typedClaims.AuthorityId))

		// 判断是否存在对应的 ACL
		casbin := casbinService.GetCasbin()

		ok, _ := casbin.Enforce(sub, obj, act)
		if ok {
			c.Next() // 请求成功
		} else {
			response.Response(c, http.StatusForbidden, "用户权限不足", nil)
			c.Abort() // 请求失败
			return
		}
	}
}
