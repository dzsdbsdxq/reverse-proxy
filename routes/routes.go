package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
	"proxy/common"
	"proxy/config"
	"proxy/middleware"
)

// InitRoutes 初始化
func InitRoutes() *gin.Engine {
	//设置模式
	gin.SetMode(config.Conf.System.Mode)
	// 创建带有默认中间件的路由:
	// 日志与恢复中间件
	r := gin.Default()

	// 启用全局跨域中间件
	r.Use(middleware.CORSMiddleware())

	for _, route := range common.ReverseConfigHandle.Routes {
		r.Any(config.Conf.System.UrlPathPrefix+`/`+route.Path+`/*name`, func(c *gin.Context) {
			var target = route.Url
			proxyUrl, _ := url.Parse(target)
			c.Request.URL.Path = c.Param("name")
			reverseProxy := httputil.NewSingleHostReverseProxy(proxyUrl)
			//这个用法就相对正规一些，不会破坏原本地请求头
			originalDirector := reverseProxy.Director // 先将原本地处理函数缓存
			reverseProxy.Director = func(req *http.Request) { // 重新赋值新地处理函数
				originalDirector(req) // 执行原本地处理函数
			}
			reverseProxy.ServeHTTP(c.Writer, c.Request)
		})
	}

	common.Log.Info("初始化路由完成！")
	return r
}
