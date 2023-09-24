package logic

import (
	"context"
	"fmt"
	"login-demo/internal/dao"
	"login-demo/internal/model"
	"login-demo/internal/model/do"
	"login-demo/internal/model/entity"
	"math"
	"math/rand"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 充电业务

type ChargeOrderLogic struct {
}

var ChargeOrder ChargeOrderLogic

func (*ChargeOrderLogic) ChargeOrderList(ctx context.Context, query model.ChargeOrderQuery) (chargeOrders []model.ChargeOrderResult, count int, err error) {
	model := dao.ChargeOrder.Ctx(ctx).LeftJoin("station", "station.id = charge_order.station_id").LeftJoin("user", "user.id = charge_order.user_id").LeftJoin("pile", "pile.id = charge_order.pile_id")
	if query.OrderCode != "" {
		model = model.Where("charge_order.order_code like ?", fmt.Sprintf("%%%s%%", query.OrderCode))
	}
	if query.StationId != 0 {
		model = model.Where("charge_order.station_id = ?", query.StationId)
	}
	if query.UserId != 0 {
		model = model.Where("charge_order.user_id = ?", query.UserId)
	}
	if query.StationId != 0 {
		model = model.Where("charge_order.station_id = ?", query.StationId)
	}
	if query.PileId != 0 {
		model = model.Where("charge_order.pile_id = ?", query.PileId)
	}
	if !query.BeforeAt.Equal(time.Time{}) {
		model = model.Where("charge_order.start_at > ?", query.BeforeAt)
	}
	if !query.AfterAt.Equal(time.Time{}) {
		model = model.Where("charge_order.start_at < ?", query.AfterAt)
	}
	if query.State != 0 {
		model = model.Where("charge_order.state = ?", query.State)
	}
	count, err = model.Count()
	if count == 0 {
		return
	}
	if err != nil {
		return
	}
	err = model.Fields("charge_order.*, station.name as station_name, pile.code as pile_code, user.nickname as nickname").OrderDesc("charge_order.update_at").Page(query.PageNo, query.PageSize).Scan(&chargeOrders)
	return
}

func (*ChargeOrderLogic) Update(ctx context.Context, update do.ChargeOrder) error {
	update.UpdateAt = gtime.Now()
	rs, err := dao.ChargeOrder.Ctx(ctx).Update(update, "id=?", update.Id)
	if err != nil {
		return gerror.WrapCode(gcode.New(1, "系统异常，修改失败", ""), err)
	}
	rowsAffected, err := rs.RowsAffected()
	if err != nil {
		return gerror.WrapCode(gcode.New(1, "系统异常，修改失败", ""), err)
	}
	if rowsAffected == 0 {
		err = gerror.NewCode(gcode.New(1, "修改失败，未找到修改原数据，可能已被删除", ""))
		return err
	}
	return nil
}

// 开始充电
// 先不支持按小时充电，需要修改数据库
func (*ChargeOrderLogic) StartCharge(ctx context.Context, pileId int, chargeHours int) (err error) {
	// 获得当前用户
	currentUser, err := User.GetCurrentUser(ctx)
	if err != nil {
		err = gerror.NewCode(gcode.New(1, "获取当前用户失败，请重新登录", err))
		return err
	}
	// 开启事务
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 1. 修改充电桩转态为充电
		pile := &entity.Pile{}
		err = dao.Pile.Ctx(ctx).Where("id=?", pileId).Scan(pile)
		if err != nil {
			err = gerror.NewCode(gcode.New(1, "找不到充电桩信息", err.Error()))
			return err
		}
		if pile.State != 0 {
			err = gerror.NewCode(gcode.New(1, "充电桩状态异常", "充电桩状态异常无法充电，请更换另一个充电桩"))
			return err
		}
		pile.State = 1
		pile.UpdateAt = gtime.Now()
		_, err = dao.Pile.Ctx(ctx).Update(pile, "id=?", pile.Id)
		if err != nil {
			err = gerror.NewCode(gcode.New(1, "系统异常", err.Error()))
			return err
		}

		// 2. 生成充电订单
		chargeOrder := do.ChargeOrder{
			OrderCode: fmt.Sprintf("%s-%d", time.Now().Format("200601021504"), rand.Intn(10000)),
			StationId: pile.StationId,
			PileId:    pile.Id,
			UserId:    currentUser.Id,
			StartAt:   gtime.Now(),
			State:     0,
			TenantId:  pile.TenantId,
			CreateAt:  gtime.Now(),
			UpdateAt:  gtime.Now(),
		}
		_, err = dao.ChargeOrder.Ctx(ctx).Insert(chargeOrder)
		return err
	})
	return
}

// 结束充电
func (*ChargeOrderLogic) StopCharge(ctx context.Context, orderId int) (err error) {
	// 开启事务
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 1. 查询充电订单记录并锁定
		chargeOrderFind := &entity.ChargeOrder{}
		err = dao.ChargeOrder.Ctx(ctx).Where("id=?", orderId).Scan(chargeOrderFind)
		if err != nil {
			return gerror.NewCode(gcode.New(1, "未找到充电订单", err.Error()))
		}
		if chargeOrderFind.State != 0 {
			return gerror.NewCode(gcode.New(1, "该充电订单已经结束", ""))
		}
		chargePrices, _, err := ChargePrice.ChargePriceList(ctx, model.PageReq{})
		if err != nil {
			return gerror.NewCode(gcode.New(1, "未找到价格", err.Error()))
		}

		// 2. 订单结束，计算金额
		hourlyRates := make([]HourlyRate, len(chargePrices))
		for i, price := range chargePrices {
			hourlyRates[i].StartHour = price.StartHour
			hourlyRates[i].EndHour = price.EndHour
			hourlyRates[i].Rate = float64(price.Price)
		}
		amount := CalculateAmount(chargeOrderFind.StartAt.Time, gtime.Now().Time, hourlyRates)

		// 3. 修改订单价格
		chargeOrderFind.State, chargeOrderFind.StopAt, chargeOrderFind.Price = 1, gtime.Now(), amount
		_, err = dao.ChargeOrder.Ctx(ctx).Update(chargeOrderFind, "id=?", chargeOrderFind.Id)
		if err != nil {
			return gerror.WrapCode(gcode.New(1, "修改订单状态失败", err), err)
		}

		// 4. 充电桩可用状态释放
		err = ReleasePile(ctx, chargeOrderFind.PileId)
		if err != nil {
			return gerror.WrapCode(gcode.New(1, "释放充电桩失败", err), err)
		}
		// 5. 用户扣除对应余额
		err = DeductUserBalance(ctx, chargeOrderFind.UserId, amount)
		if err != nil {
			return gerror.WrapCode(gcode.New(1, "扣除用户金额失败", err), err)
		}
		return nil
	})
	return
}

// 充电桩可用状态释放
func ReleasePile(ctx context.Context, pileId int) error {
	pileFind := &entity.Pile{}
	err := dao.Pile.Ctx(ctx).Where("id=?", pileId).Scan(pileFind)
	if err != nil {
		return gerror.NewCode(gcode.New(1, "获取充电桩失败", err.Error()))
	}
	if pileFind.State != 1 {
		return gerror.NewCode(gcode.New(1, "充电桩状态异常", fmt.Sprintf("当前充电桩状态%d", pileFind.State)))
	}
	pileFind.State = 0
	_, err = dao.Pile.Ctx(ctx).Update(pileFind, "id=?", pileFind.Id)
	if err != nil {
		return gerror.NewCode(gcode.New(1, "获取充电桩失败", err.Error()))
	}
	return err
}

// 用户扣除对应余额
func DeductUserBalance(ctx context.Context, userId int, amount float64) error {
	userWalletFind := &entity.UserWallet{}
	err := dao.UserWallet.Ctx(ctx).Where("user_id=?", userId).Scan(userWalletFind)
	if err != nil {
		return gerror.NewCode(gcode.New(1, "未找到用户金额信息", err.Error()))
	}
	userWalletFind.Balance = userWalletFind.Balance - amount
	_, err = dao.UserWallet.Ctx(ctx).Update(userWalletFind, "user_id=?", userWalletFind.UserId)
	if err != nil {
		return gerror.NewCode(gcode.New(1, "修改用户金额失败", err.Error()))
	}
	return err
}

type HourlyRate struct {
	StartHour int     // 起始小时数
	EndHour   int     // 结束小时数
	Rate      float64 // 每小时费率
}

// 计算费率
func CalculateAmount(startTime, endTime time.Time, hourlyRates []HourlyRate) float64 {
	var (
		amount float64 // 总金额
		hours  int     // 总小时数
		mins   int     // 总分钟数
	)

	// 计算总小时数和总分钟数
	duration := endTime.Sub(startTime)
	hours = int(duration.Hours())
	mins = int(duration.Minutes()) % 60

	// 计算总金额
	for i := 0; i < hours; i++ {
		var rate float64
		for _, hourlyRate := range hourlyRates {
			if startTime.Hour() >= hourlyRate.StartHour && startTime.Hour() < hourlyRate.EndHour {
				rate = hourlyRate.Rate
				break
			}
		}
		amount += rate
		startTime = startTime.Add(time.Hour)
	}

	// 如果还有不足一小时的时间，按照半小时计算
	if mins > 0 {
		var rate float64
		for _, hourlyRate := range hourlyRates {
			if startTime.Hour() >= hourlyRate.StartHour && startTime.Hour() < hourlyRate.EndHour {
				rate = hourlyRate.Rate
				break
			}
		}
		amount += rate / 2
	}
	// 四舍五入第一位小数点，解决 float 精度问题
	amount = math.Round(amount*10) / 10
	return amount
}
