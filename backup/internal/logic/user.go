package logic

import (
	"context"
	"errors"
	"fmt"
	"log"
	"login-demo/internal/consts"
	"login-demo/internal/dao"
	"login-demo/internal/model"
	"login-demo/internal/model/do"
	"login-demo/internal/model/entity"
	"login-demo/internal/packed/jwt"

	"crypto/sha256"

	"github.com/fatih/color"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
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
		token, err = jwt.GenerateToken(ctx, username)
		if err != nil {
			return "", "", err
		}
		// todo: 暂时没有 role
		return "", token, nil
	}
}

type AdminInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// InitAdmin 初始化管理员账户，从配置文件中
func (*LogicUser) InitAdmin(ctx context.Context) {
	admin := &AdminInfo{}
	ctx = context.WithValue(ctx, consts.TenantIDKey, consts.DefaultTenantValue)
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
			TenantId: consts.DefaultTenantValue,
		})
	}
}

func (*LogicUser) UserList(ctx context.Context, username string, page model.PageReq) (users []model.UserMore, count int, err error) {
	err = dao.User.Ctx(ctx).LeftJoin("user_wallet", "user.id = user_wallet.user_id").LeftJoin("wx_user", "user.id = wx_user.user_id").LeftJoin("role", "user.role_id = role.id").Fields("user.*, user_wallet.balance, wx_user.avatar_url, role.name as role_name").WhereLike("user.passport", fmt.Sprintf("%%%s%%", username)).OrderDesc("user.update_at").Page(page.PageNo, page.PageSize).ScanAndCount(&users, &count, false)
	return
}

func (*LogicUser) UserCount(ctx context.Context, queryUser do.User) (count int, err error) {
	userModel := dao.User.Ctx(ctx)
	if queryUser.Nickname != nil {
		userModel = userModel.WhereLike("passport", fmt.Sprintf("%%%s%%", queryUser.Passport.(string)))
	}
	if queryUser.RoleId != nil {
		userModel = userModel.Where("role_id", queryUser.RoleId.(int))
	}
	count, err = userModel.Count()
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

func (l *LogicUser) AddAndWallet(ctx context.Context, user *do.User) (err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = l.Add(ctx, user)
		if err != nil {
			return err
		}
		_, err = dao.UserWallet.Ctx(ctx).Insert(do.UserWallet{
			UserId:   user.Id,
			Balance:  0,
			TenantId: user.TenantId,
			CreateAt: gtime.Now(),
			UpdateAt: gtime.Now(),
		})
		return err
	})
	return
}

func (*LogicUser) Del(ctx context.Context, id int) (err error) {
	rs, err := dao.User.Ctx(ctx).Delete("id = ?", id)
	if err != nil {
		return gerror.WrapCode(gcode.New(1, "系统异常，删除失败", ""), err)
	}
	rowsAffected, err := rs.RowsAffected()
	if err != nil {
		return gerror.WrapCode(gcode.New(1, "系统异常，删除失败", ""), err)
	}
	if rowsAffected == 0 {
		err = gerror.NewCode(gcode.New(1, "删除失败，未找到原数据，可能已被其他人删除", ""))
		return err
	}
	return
}

func (*LogicUser) Update(ctx context.Context, user do.User) (err error) {
	if user.Password != nil {
		// 计算密码哈希值
		passwordHash := fmt.Sprintf("%x", sha256.Sum256([]byte(user.Password.(string))))
		user.Password = passwordHash
	}
	user.UpdateAt = gtime.Now()
	rs, err := dao.User.Ctx(ctx).Update(user, "id = ?", user.Id)
	if err != nil {
		return gerror.WrapCode(gcode.New(1, "系统异常，修改失败", ""), err)
	}
	rowsAffected, err := rs.RowsAffected()
	if err != nil {
		return gerror.WrapCode(gcode.New(1, "系统异常，修改失败", ""), err)
	}
	if rowsAffected == 0 {
		err = gerror.NewCode(gcode.New(1, "修改失败，未找到原数据，可能已被其他人删除", ""))
		return err
	}
	return
}

// 设置 username 到 context 中
func (LogicUser) SetUserContext(username string, setCtx func(interface{}, interface{})) {
	setCtx(consts.UserContextKey, model.UserContext{
		Username: username,
	})
}

// 从 context 取出 user 信息“
func (l LogicUser) GetCurrentUser(ctx context.Context) (user model.UserMore, err error) {
	userCtx, ok := ctx.Value(consts.UserContextKey).(model.UserContext)
	if !ok {
		err = gerror.NewCode(gcode.New(404, "未找到登录用户信息", ""))
		return
	}
	users, count, err := User.UserList(ctx, userCtx.Username, model.PageReq{})
	if err != nil {
		return
	}
	if count != 1 {
		err = gerror.NewCode(gcode.New(404, fmt.Sprintf("%s 用户不存在，请重新注册", userCtx.Username), ""))
		return
	}
	user = users[0]
	return
}
