package controller

import (
	"errors"
	"fmt"
	"github.com/allegro/bigcache/v3"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"server/internal/e"
	"server/internal/global"
	"server/internal/model"
	"server/internal/model/response"
	"server/internal/service"
	"server/internal/util"
	"strconv"
)

type AuthController struct{}

// PostLogin
//
//	@Description: 完成用户登录
//	@receiver ac
//	@Router    /auth/login
func (ac *AuthController) PostLogin(c *gin.Context) {
	var logonRequest model.LoginRequest

	if err := c.ShouldBindJSON(&logonRequest); err != nil {
		response.Response(c, http.StatusInternalServerError, "参数解析错误！", nil)
		return
	}

	// 验证码
	openCaptcha := global.Config.Captcha.OpenCaptcha               // 是否开启验证码
	openCaptchaTimeout := global.Config.Captcha.OpenCaptchaTimeout // 验证码超时时间

	key := c.ClientIP() // 客户端 IP

	item, err := util.GetCacheItem(key)
	if err != nil {
		// 当条目不存在时或者超时时，初始化条目
		if errors.Is(err, bigcache.ErrEntryNotFound) || errors.Is(err, e.ErrCacheEntryTimeout) {
			util.SetCacheItem(key, []byte("1"), openCaptchaTimeout)
		} else {
			global.Logger.Error("获取缓存条目错误！", zap.Error(err))
			return
		}
	}

	var oc = openCaptcha == 0 || openCaptcha < util.ItemToInt(item)

	if !oc || (logonRequest.CaptchaId != "" && logonRequest.Answer != "" && global.CaptchaStore.Verify(logonRequest.CaptchaId, logonRequest.Answer, true)) {
		u := model.User{Username: logonRequest.Username, Password: logonRequest.Password}
		var user = &model.User{}
		if user, err = service.AuthServiceApp.Login(u); err != nil {
			global.Logger.Error(fmt.Sprintf("用户 %v 登陆失败：用户名不存在或者密码错误！", logonRequest.Username), zap.Error(err))
			util.SetCacheItem(key, []byte(strconv.Itoa(util.ItemToInt(item)+1)), openCaptchaTimeout)
			response.Response(c, http.StatusInternalServerError, fmt.Sprintf("用户 %v 登陆失败：用户名不存在或者密码错误！", logonRequest.Username), nil)
			return
		} // 用户身份校验失败
		if user.Enable != 1 {
			global.Logger.Error(fmt.Sprintf("用户 %v 登陆失败：用户被冻结，禁止登录!", logonRequest.Username))
			util.SetCacheItem(key, []byte(strconv.Itoa(util.ItemToInt(item)+1)), openCaptchaTimeout)
			response.Response(c, http.StatusInternalServerError, fmt.Sprintf("用户 %v 登陆失败：用户被冻结，禁止登录！", logonRequest.Username), nil)
			return
		} // 用户账户被冻结
		ac.TokenNext(c, user) // 用户登录成功，返回 Token
		return
	} else {
		util.SetCacheItem(key, []byte(strconv.Itoa(util.ItemToInt(item)+1)), openCaptchaTimeout)
		response.Response(c, http.StatusUnauthorized, fmt.Sprintf("用户 %v 登陆失败：验证码错误！", logonRequest.Username), nil)
		return
	}
}

func (ac *AuthController) TokenNext(c *gin.Context, user *model.User) {
	jwt, err := util.GenerateJWT(util.CustomClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		Username:    user.Username,
		AuthorityId: user.AuthorityId,
	})
	if err != nil {
		global.Logger.Error(fmt.Sprintf("用户 %v 登陆失败：获取 token 失败！", user.Username), zap.Error(err))
		response.Response(c, http.StatusUnauthorized, fmt.Sprintf("用户 %v 登陆失败：获取 token 失败！", user.Username), nil)
		return
	} else {
		global.Logger.Info(fmt.Sprintf("用户 %v 登陆成功！", user.Username))
		response.Response(c, http.StatusOK, "登陆成功！", gin.H{"token": jwt})
		return
	}

}
