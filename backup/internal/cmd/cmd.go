package cmd

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	managerController "login-demo/internal/controller/manager"
	saasController "login-demo/internal/controller/saas"
	wxController "login-demo/internal/controller/wx"
	"login-demo/internal/logic"
	"login-demo/internal/middleware"

	"os"

	"github.com/gogf/gf/v2/os/gcfg"
)

func UploadShow(r *ghttp.Request) {
	r.Response.Write(`
    <html>
    <head>
        <title>GF Upload File Demo</title>
    </head>
        <body>
            <form enctype="multipart/form-data" action="/saas-api/station/upload" method="post">
                <input type="file" name="file" />
                <input type="submit" value="upload" />
            </form>
        </body>
    </html>
    `)
}

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			s.Use(ghttp.MiddlewareHandlerResponse)
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.ALL("/show", UploadShow)
				group.Group("/saas-api", func(group *ghttp.RouterGroup) {
					// 绑定 Login 结构体中的 Login 方法
					group.GET("", new(managerController.TenantController).SelectList)
					group.POST("", new(managerController.TenantController).Register)
					group.Middleware(middleware.TenantMiddleware)
					group.POST("", new(saasController.Login).Login)
					group.Group("", func(group *ghttp.RouterGroup) {
						group.Middleware(middleware.Auth)
						group.POST("", new(saasController.Login).Refresh)
						group.Bind(
							&saasController.UserController{},
							&saasController.StationController{},
							&saasController.PileController{},
							&saasController.ChargeOrderController{},
							&saasController.ChargePriceController{},
							&saasController.UserPayController{},
							&saasController.RoleController{},
							&saasController.PermissionController{},
							&saasController.AuditLogController{},
						)
					})
				})
				group.Group("/manager-api", func(group *ghttp.RouterGroup) {
					// 绑定 Login 结构体中的 Login 方法
					group.POST("", new(managerController.Login).Login)
					group.Group("", func(group *ghttp.RouterGroup) {
						group.Middleware(middleware.Auth)
						group.POST("", new(managerController.Login).Refresh)
						group.GET("", new(managerController.TenantController).List)
						group.POST("", new(managerController.TenantController).Add)
						group.POST("", new(managerController.TenantController).Update)
						group.POST("", new(managerController.TenantController).Del)
						group.Bind(
							&saasController.UserController{},
							&saasController.StationController{},
							&saasController.PileController{},
							&saasController.ChargeOrderController{},
							&saasController.ChargePriceController{},
							&saasController.UserPayController{},
							&saasController.RoleController{},
							&saasController.PermissionController{},
							&saasController.AuditLogController{},
						)
					})
				})
				// 微信路由
				group.Bind(
					&wxController.WxLoginController{},
				)
				group.GET("", new(wxController.WxChargeController).StationList)
				group.GET("", new(wxController.WxChargeController).PileList)
				group.Group("", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.Auth)
					group.POST("", new(wxController.WxChargeController).StartCharge)
					group.POST("", new(wxController.WxChargeController).MyChargeOrders)
					group.POST("", new(wxController.WxChargeController).AboutMe)
					group.POST("", new(wxController.WxChargeController).StopCharge)
					group.POST("", new(wxController.WxChargeController).PriceList)
					group.POST("", new(wxController.WxChargeController).UploadAvatar)
					group.POST("", new(wxController.WxChargeController).UpdateNickname)
				})
			})

			s.SetIndexFolder(true)
			config, err := gcfg.New()
			localPath := config.MustGet(ctx, "storage.local.path").String()

			// 检查目录是否存在
			if _, err := os.Stat(localPath); os.IsNotExist(err) {
				fmt.Println(localPath, "目录不存在，正在创建")
				// 目录不存在，创建目录
				err := os.MkdirAll(localPath, os.ModePerm)
				if err != nil {
					fmt.Println("创建目录失败:", err)
					return err
				}
				fmt.Println("目录创建成功")
			}

			s.SetServerRoot(localPath)
			logic.User.InitAdmin(ctx)
			s.Run()
			return nil
		},
	}
)
