package core

import (
	"github.com/gin-gonic/gin"
	"reflect"
	authController "server/app/auth/controller"
	baseController "server/app/base/controller"
	casbinController "server/app/casbin/controller"
	initController "server/app/init/controller"
	userController "server/app/user/controller"
	"server/internal/config"
	"server/internal/middleware"
	"strings"
)

func InitializeRout() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery()) // 避免 panic 导致服务器停止

	RegisterRoutes(router, GetControllerList())

	return router
}

func GetControllerList() []interface{} {
	return []interface{}{
		&baseController.BaseController{},
		&userController.UserController{},
		&authController.AuthController{},
		&casbinController.CasbinController{},
		&initController.InitController{},
	}
}

func RegisterRoutes(router *gin.Engine, controllers []interface{}) {
	publicGroup := router.Group(config.GlobalConfig.SystemConfig.RouterPrefix) // 无需鉴权的路由组

	protectedGroup := router.Group(config.GlobalConfig.SystemConfig.RouterPrefix)      // 需要鉴权的路由组
	protectedGroup.Use(middleware.JWTAuthMiddleware()).Use(middleware.CasbinHandler()) // 使用 JWT 和 Casbin 完成身份验证以及访问控制

	for _, ctrl := range controllers {
		ctrlType := reflect.TypeOf(ctrl)
		ctrlValue := reflect.ValueOf(ctrl)
		ctrlName := strings.TrimSuffix(ctrlType.Elem().Name(), "Controller")

		_, jwtRequired := ctrlType.Elem().FieldByName("JWTRequired") // 检查是否需要鉴权

		for i := 0; i < ctrlType.NumMethod(); i++ {
			method := ctrlType.Method(i)
			methodName := method.Name

			if httpMethod, ok := getHTTPMethodFromName(methodName); ok {
				path := "/" + strings.ToLower(ctrlName) + "/" + strings.ToLower(strings.TrimPrefix(methodName, httpMethod)) // 创建路径

				handleFunc := func(c *gin.Context) {
					ctrlValue.MethodByName(methodName).Call([]reflect.Value{reflect.ValueOf(c)})
				} // 通过反射调用对应的方法

				if jwtRequired {
					protectedGroup.Handle(httpMethod, path, handleFunc)
				} else {
					publicGroup.Handle(httpMethod, path, handleFunc)
				}
			}
		} // 遍历 Controller 中实现的方法并添加到路由组
	}
}

// getHTTPMethodFromName
//
//	@Description: 根据方法名前缀获取对应的请求方法
//	@param methodName
//	@return string
//	@return bool
func getHTTPMethodFromName(methodName string) (string, bool) {
	if strings.HasPrefix(methodName, "Get") {
		return "GET", true
	} else if strings.HasPrefix(methodName, "Post") {
		return "POST", true
	}
	return "", false
}
