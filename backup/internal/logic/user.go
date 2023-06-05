package logic

import (
	"context"
	"errors"
	"fmt"
	"log"
	"login-demo/internal/dao"
	"login-demo/internal/model"
	"login-demo/internal/model/do"
	"login-demo/internal/model/entity"
	"strings"

	"crypto/sha256"

	"github.com/fatih/color"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
)

type LogicUser struct {
}

var (
	User = LogicUser{}
)

// Login 登录
func (lu *LogicUser) Login(ctx context.Context, username string, password string) (role string, token string, err error) {
	// 计算密码哈希值
	passwordHash := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
	// 查询用户信息
	user, err := dao.User.Ctx(ctx).One("passport=? and password=?", username, passwordHash)
	switch {
	case err != nil:
		return "", "", err
	case user == nil:
		return "", "", errors.New("账户或密码错误")
	default:
		// 生成 jwt token
		token, err = MyJwt.GenerateToken(ctx, username)
		if err != nil {
			return "", "", err
		}
		// todo: 暂时没有 role
		return "", token, nil
	}
}

// IsSignedIn 检查是否已经登录
func (lu *LogicUser) IsSignedIn(ctx context.Context, r *ghttp.Request) bool {
	token, exist := lu.getToken(r)
	if !exist {
		return false
	}
	return MyJwt.Valid(r.Context(), token)
}

// Parse 解析 jwt token
func (lu *LogicUser) Parse(ctx context.Context, r *ghttp.Request) (bool, string) {
	token, exist := lu.getToken(r)
	if !exist {
		return false, ""
	}
	claims, ok := MyJwt.Parse(r.Context(), token)
	if !ok {
		return false, ""
	}
	return ok, claims.Username
}

// getToken 从 request 的 header 中获取 token
func (*LogicUser) getToken(r *ghttp.Request) (string, bool) {
	header := r.GetHeader("Authorization")
	headerList := strings.Split(header, " ")
	if len(headerList) != 2 {
		return "", false
	}
	t := headerList[0]
	token := headerList[1]
	if t != "Bearer" {
		return "", false
	}
	if token == "" {
		return "", false
	}
	return token, true
}

type AdminInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// InitAdmin 初始化管理员账户，从配置文件中
func (*LogicUser) InitAdmin(ctx context.Context) {
	admin := &AdminInfo{}
	err := g.Cfg().MustGet(ctx, "admin").Scan(admin)
	if err != nil {
		log.Fatal("读取admin配置失败")
		return
	}
	var user *entity.User
	err = dao.User.Ctx(ctx).Where(do.User{
		Passport: admin.Username,
	}).Scan(&user)
	if err != nil {
		g.Log().Infof(ctx, color.RedString("err=%v, dao.User=%v"), err, user)
		panic("查询用户表失败，是否没有连接正确的数据库")
	}
	if user == nil {
		dao.User.Ctx(ctx).Insert(do.User{
			Passport: admin.Username,
			Password: fmt.Sprintf("%x", sha256.Sum256([]byte(admin.Password))),
			Nickname: admin.Username,
		})
	}
}

func init() {
	ctx := gctx.New()
	User.InitAdmin(ctx)
}

func (*LogicUser) UserList(ctx context.Context, username string, page model.PageReq) (users []model.UserMore, count int, err error) {
	err = dao.User.Ctx(ctx).LeftJoin("user_wallet", "user.id = user_wallet.user_id").Fields("user.*, user_wallet.balance").WhereLike("passport", fmt.Sprintf("%%%s%%", username)).Page(page.PageNo, page.PageSize).ScanAndCount(&users, &count, false)
	return
}

func (*LogicUser) Add(ctx context.Context, user *do.User) (err error) {
	// 计算密码哈希值
	passwordHash := fmt.Sprintf("%x", sha256.Sum256([]byte(user.Password.(string))))
	user.CreateAt, user.UpdateAt, user.Password = gtime.Now(), gtime.Now(), passwordHash
	id, err := dao.User.Ctx(ctx).InsertAndGetId(user)
	user.Id = id
	return
}

func (*LogicUser) Del(ctx context.Context, id int) (err error) {
	_, err = dao.User.Ctx(ctx).Delete("id = ?", id)
	return
}

func (*LogicUser) Update(ctx context.Context, user do.User) (err error) {
	// 计算密码哈希值
	passwordHash := fmt.Sprintf("%x", sha256.Sum256([]byte(user.Password.(string))))
	user.UpdateAt, user.Password = gtime.Now(), passwordHash
	_, err = dao.User.Ctx(ctx).Update(user, "id = ?", user.Id)
	return
}
