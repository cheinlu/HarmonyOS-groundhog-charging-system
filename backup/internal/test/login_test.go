package controller_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	v1 "login-demo/api/v1"
	"login-demo/internal/controller"
	_ "login-demo/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gfile"

	"github.com/gogf/gf/os/gctx"
)

var (
	db         gdb.DB
	dbPrefix   gdb.DB
	dbInvalid  gdb.DB
	configNode gdb.ConfigNode
	dbDir      = gfile.Temp("sqlite")
	ctx        = gctx.New()

	// Error
	ErrorSave = gerror.NewCode(gcode.CodeNotSupported, `Save operation is not supported by sqlite driver`)
)

const (
	TableSize               = 10
	TableName               = "user"
	TableNameWhichIsKeyword = "group"
	TestSchema1             = "test1"
	TestSchema2             = "test2"
	TableNamePrefix         = "gf_"
	CreateTime              = "2018-10-24 10:00:00"
	DBGroupTest             = "test"
	DBGroupPrefix           = "prefix"
	DBGroupInvalid          = "invalid"
)

func TestLogin(t *testing.T) {
	ctx := gctx.New()
	var login controller.Login
	loginReq := v1.LoginReq{
		Username: "admin",
		Password: "admin",
	}

	res, err := login.Login(ctx, &loginReq)
	assert.Nil(t, err)
	assert.NotEmpty(t, res.Token)
}

// func TestRefresh(t *testing.T) {
// 	r := ghttp.NewTestRequest()
// 	r.SetMethod("POST")
// 	r.SetUrl("/refresh")
// 	r.SetHeader("Authorization", "Bearer "+testToken)
// 	resp := r.GetResponse()
// 	defer resp.Close()
// 	assert.Equal(t, resp.StatusCode, 200)
// 	data := resp.GetJson().GetMap("data")
// 	assert.NotEmpty(t, data.GetString("token"))
// }
