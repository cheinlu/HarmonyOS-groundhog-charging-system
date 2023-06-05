package logic

import (
	"context"
	"database/sql"
	"fmt"
	"login-demo/internal/dao"
	"login-demo/internal/model"
	"login-demo/internal/model/do"
	"login-demo/internal/model/entity"
	"math/rand"
	"time"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

// 用户充值相关
var UserPay UserPayLogic

type UserPayLogic struct {
}

func (*UserPayLogic) UserPayList(ctx context.Context, page model.PageReq) (userPays []model.PayRecordInfo, count int, err error) {
	err = dao.PayRecord.Ctx(ctx).LeftJoin("user", "user.id=pay_record.user_id").Page(page.PageNo, page.PageSize).Fields("pay_record.*, user.passport as username").ScanAndCount(&userPays, &count, false)
	return
}

func (*UserPayLogic) UserPay(ctx context.Context, userPay do.PayRecord) (err error) {
	tx, err := dao.ChargeOrder.DB().Begin(ctx)
	if err != nil {
		return
	}
	count, err := dao.User.Ctx(ctx).Where("id=?", userPay.UserId).Count()
	if count == 0 {
		err = gerror.NewCode(gcode.New(404, "用户不存在", "用户不存在"))
		return
	}
	if err != nil {
		return
	}
	userWalletExist := new(entity.UserWallet)
	err = dao.UserWallet.Ctx(ctx).Where("user_id=?", userPay.UserId).Scan(userWalletExist)
	if err == sql.ErrNoRows {
		userWallet := do.UserWallet{
			UserId:   userPay.UserId,
			Balance:  userPay.Price,
			CreateAt: gtime.Now(),
			UpdateAt: gtime.Now(),
		}
		_, err = dao.UserWallet.Ctx(ctx).Insert(userWallet)
	} else {
		fmt.Printf("userPay.Price:%T, userWalletExist.Balance:%T\n", userPay.Price, userWalletExist.Balance)
		userWallet := do.UserWallet{
			UserId:   userWalletExist.UserId,
			Balance:  userWalletExist.Balance + userPay.Price.(float64),
			UpdateAt: gtime.Now(),
		}
		_, err = dao.UserWallet.Ctx(ctx).Update(userWallet, "user_id", userWalletExist.UserId)
	}
	if err != nil {
		return
	}
	rand.Seed(time.Now().UnixNano())
	userPay.State, userPay.CreateAt, userPay.UpdateAt = 1, gtime.Now(), gtime.Now()
	userPay.PayAt = gtime.Now()
	userPay.PayCode = fmt.Sprintf("%s-%d", time.Now().Format("200601021504"), rand.Intn(10000))
	if err != nil {
		userPay.State = 2
	}
	_, err = dao.PayRecord.Ctx(ctx).Insert(userPay)
	tx.Commit()
	return
}
