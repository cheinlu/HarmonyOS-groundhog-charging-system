package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"login-demo/internal/controller"
	"login-demo/internal/middleware"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Use(ghttp.MiddlewareHandlerResponse)
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					middleware.ParseJwtToCtx,
				)
				// 绑定 Login 结构体中的 Login 方法
				group.POST("", new(controller.Login).Login)
				group.Bind()
				group.Group("", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.Auth)
					group.POST("", new(controller.Login).Refresh)
					group.Bind(
						&controller.StationController{},
						&controller.UserController{},
						&controller.PileController{},
						&controller.ChargeOrderController{},
						&controller.ChargePriceController{},
						&controller.UserPayController{},
					)
				})
				// group.POST("", new(controller.WxLoginController).Login)
				group.Bind(
					&controller.WxLoginController{},
				)
				group.GET("", new(controller.WxChargeController).StationList)
				group.GET("", new(controller.WxChargeController).PileList)
				group.Group("", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.Auth)
					// group.Bind(
					// 	&controller.WxChargeController{},
					// )
					group.POST("", new(controller.WxChargeController).StartCharge)
					group.POST("", new(controller.WxChargeController).MyChargeOrders)
					group.POST("", new(controller.WxChargeController).AboutMe)
					group.POST("", new(controller.WxChargeController).StopCharge)
				})
			})
			s.Run()
			return nil
		},
	}
)
